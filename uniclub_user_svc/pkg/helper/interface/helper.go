package interfaces

import "github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/utils/models"

type UserServiceHelper interface {
	PasswordHashing(string) (string, error)
	GenerateTokenClients(user models.UserDetailsResponse) (string, string, error)
	GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error)
	CompareHashAndPassword(hashedPassword string, password string) error
	// TwilioSetup(accountSID string, authToken string)
	// TwilioSendOTP(phoneNo string, serviceSID string) (string, error)
	// TwilioVerifyOTP(serviceSID string, code string, phoneNo string) error
	// SendMailToPhone(To, Subject, Msg string) error
	GenerateReferralCode() (string, error)
}
