package interfaces

import (
	"mime/multipart"
	"time"

	"github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/utils/models"
)

type InventoryServiceHelper interface {
  PasswordHashing(string) (string, error)
  GenerateTokenClients(user models.UserDetailsResponse) (string, string, error)
  GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error)
  CompareHashAndPassword(hashedPassword string, password string) error
  GenerateReferralCode() (string, error)
  StringToTime(timeStr string) (time.Time, error) 
  TimeToString(t time.Time) string 

  AddImageToAwsS3(file *multipart.FileHeader) (string, error)

  SendMailToPhone(To, Subject, Msg string) error


}
