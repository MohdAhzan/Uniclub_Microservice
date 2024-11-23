package helper

//
import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"mime/multipart"
	"net/smtp"
	"time"

	"github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/config"
	aws_config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws"

	"golang.org/x/crypto/bcrypt"
)

type  InventoryServiceHelper struct {
	cfg config.Config
}


func NewHelper(config config.Config) *InventoryServiceHelper {
	return &InventoryServiceHelper{
		cfg: config,
	}
}
func (h *InventoryServiceHelper)StringToTime(timeStr string) (time.Time, error) {
	layout := time.RFC3339
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func (h *InventoryServiceHelper) TimeToString(t time.Time) string {
	return t.Format(time.RFC3339)
}

func (h *InventoryServiceHelper) AddImageToAwsS3(file *multipart.FileHeader) (string, error) {

	cfg, err := aws_config.LoadDefaultConfig(context.TODO(), aws_config.WithRegion("ap-southeast-2"))
	if err != nil {
		return "", err
	}

	client := s3.NewFromConfig(cfg)

	uploader := manager.NewUploader(client)
	f, openErr := file.Open()
	if openErr != nil {
		return "", openErr
	}
	defer f.Close()

	result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String("jpeg123"),
		Key:         aws.String(file.Filename),
		Body:        f,
		ACL:         types.ObjectCannedACLPublicRead,
		ContentType: aws.String("image/png"),
	})

	if uploadErr != nil {
		fmt.Println("uploadERR", uploadErr)
		return "", uploadErr
	}

	return result.Location, nil
}


func (h *InventoryServiceHelper) PasswordHashing(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}

	hash := string(hashedPassword)
	return hash, nil
}



func (h *InventoryServiceHelper) GenerateReferralCode() (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 12

	// Initialize the result string.
	result := make([]byte, length)

	// Generate a random index for each character in the result string.
	for i := range result {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[idx.Int64()]
	}

	return string(result), nil
}



// var client *twilio.RestClient
//
// func (h *UserServiceHelper) TwilioSetup(accountSID string, authToken string) {
//
// 	client = twilio.NewRestClientWithParams(twilio.ClientParams{
// 		Username: accountSID,
// 		Password: authToken,
// 	})
// }
//
// func (h *Helper) TwilioSendOTP(phoneNo string, serviceSID string) (string, error) {
// 	// fmt.Println("phone no is =", phoneNo, "     and servicesid is =", serviceSID)
//
// 	to := "+91" + phoneNo
// 	params := &openApi.CreateVerificationParams{}
// 	params.SetTo(to)
// 	params.SetChannel("sms")
//
// 	resp, err := client.VerifyV2.CreateVerification(serviceSID, params)
// 	// fmt.Println("VErificatoino Params", params)
// 	if err != nil {
//
// 		return " ", err
// 	}
// 	fmt.Println("verificatoin SID", *resp.Sid)
// 	return *resp.Sid, nil
//
// }
//
// func (h *Helper) TwilioVerifyOTP(serviceSID string, code string, phoneNo string) error {
//
// 	params := &openApi.CreateVerificationCheckParams{}
// 	params.SetTo("+91" + phoneNo)
// 	params.SetCode(code)
// 	resp, err := client.VerifyV2.CreateVerificationCheck(serviceSID, params)
//
// 	if err != nil {
// 		fmt.Println("ERRORR is", err)
// 		return err
// 	}
//
// 	if *resp.Status == "approved" {
// 		return nil
// 	}
//
// 	return errors.New("failed to validate otp")
//
// }
//
// func (h *Helper) AddImageToAwsS3(file *multipart.FileHeader) (string, error) {
//
//   aws.Config
// 	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-2"))
// 	if err != nil {
// 		return "", err
// 	}
//
// 	client := s3.NewFromConfig(cfg)
//
// 	uploader := manager.NewUploader(client)
// 	f, openErr := file.Open()
// 	if openErr != nil {
// 		return "", openErr
// 	}
// 	defer f.Close()
//
// 	result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
// 		Bucket:      aws.String("jpeg123"),
// 		Key:         aws.String(file.Filename),
// 		Body:        f,
// 		ACL:         types.ObjectCannedACLPublicRead,
// 		ContentType: aws.String("image/png"),
// 	})
//
// 	if uploadErr != nil {
// 		fmt.Println("uploadERR", uploadErr)
// 		return "", uploadErr
// 	}
//
// 	return result.Location, nil
// }
//
func (h *InventoryServiceHelper) SendMailToPhone(To, Subject, Msg string) error {

	TO := []string{To}

	//setup authentication
	auth := smtp.PlainAuth("", h.cfg.SMTP_USERNAME, h.cfg.SMTP_PASSWORD, h.cfg.SMTP_HOST)

	//message body
	msg := []byte("To: " + TO[0] + "\r\n" +
		"Subject: " + Subject + "\r\n" +
		"\r\n" +
		Msg + "\r\n")
	//send mail to recipient
	err := smtp.SendMail(h.cfg.SMTP_HOST+":"+h.cfg.SMTP_PORT, auth, h.cfg.SMTP_USERNAME, TO, msg)
	if err != nil {
		return err
	}
	return nil

}

func (h *InventoryServiceHelper) GetTimeFromPeriod(timePeriod string) (time.Time, time.Time) {

	endDate := time.Now()

	if timePeriod == "week" {
		startDate := endDate.AddDate(0, 0, -6)
		return startDate, endDate
	}

	if timePeriod == "month" {
		startDate := endDate.AddDate(0, -1, 0)
		return startDate, endDate
	}

	if timePeriod == "year" {
		startDate := endDate.AddDate(-1, 0, 0)
		return startDate, endDate
	}

	return endDate.AddDate(0, 0, -6), endDate

}

//
// func (h *Helper) ConvertToExel(sales []models.OrderDetailsAdmin) (*excelize.File, error) {
//
// 	filename := "../salesReport/sales_report.xlsx"
// 	file := excelize.NewFile()
//
// 	file.SetCellValue("Sheet1", "A1", "Product")
// 	file.SetCellValue("Sheet1", "B1", "Amount Sold")
//
// 	// Bold style for headings
// 	boldStyle, err := file.NewStyle(`{"font":{"bold":true}}`)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	file.SetCellStyle("Sheet1", "A1", "B1", boldStyle)
//
// 	var Total float64
// 	var Limit int
// 	for i, sale := range sales {
// 		col1 := fmt.Sprintf("A%d", i+2)
// 		col2 := fmt.Sprintf("B%d", i+2)
//
// 		file.SetCellValue("Sheet1", col1, sale.ProductName)
// 		file.SetCellValue("Sheet1", col2, sale.TotalAmount)
// 		Limit = i + 3
// 		Total += sale.TotalAmount
//
// 	}
// 	col1 := fmt.Sprintf("A%d", Limit)
// 	file.SetCellValue("Sheet1", col1, "Final Total")
// 	col2 := fmt.Sprintf("B%d", Limit)
// 	file.SetCellValue("Sheet1", col2, Total)
//
// 	// Larger font size for 'Final Total'
// 	largerFontStyle, err := file.NewStyle(`{"font":{"size":10}}`)
// 	if err != nil {
// 		return nil, err
// 	}
// 	file.SetCellStyle("Sheet1", col1, col2, largerFontStyle)
//
// 	if err := file.SaveAs(filename); err != nil {
// 		return nil, err
// 	}
//
// 	return file, nil
//
// 	// var Total float64
// 	// for i, sale := range sales {
// 	// 	col1 := fmt.Sprintf("A%d", i+2)
// 	// 	col2 := fmt.Sprintf("B%d", i+2)
//
// 	// 	file.SetCellValue("Sheet1", col1, sale.ProductName)
// 	// 	file.SetCellValue("Sheet1", col2, sale.TotalAmount)
// 	// 	Total += sale.TotalAmount
// 	// }
//
// 	// Limit := len(sales) + 2
// 	// col1 := fmt.Sprintf("A%d", Limit)
// 	// file.SetCellValue("Sheet1", col1, "Final Total")
// 	// col2 := fmt.Sprintf("B%d", Limit)
// 	// file.SetCellValue("Sheet1", col2, Total)
//
// 	// // Larger font size for 'Final Total'
// 	// largerFontStyle, err := file.NewStyle(`{"font":{"size":10}}`)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// file.SetCellStyle("Sheet1", col1, col2, largerFontStyle)
//
// 	// if err := file.SaveAs(filename); err != nil {
// 	// 	return nil, err
// 	// }
//
// 	// return file, nil
//
// }

