package interfaces

import (
	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/utils/models"
)

type UserServiceClient interface{

  AdminLoginHandler(adminDetails models.AdminLogin) (models.TokenAdmin, error)
	
  GetUsers() ([]models.UserDetailsAtAdmin, error)
	BlockUser(id int) error
	UnBlockUser(id int) error
	ChangeAdminPassword(passChange models.AdminPasswordChange, id int) error

  UserSignup(user models.UserDetails, refCode string) (models.TokenUsers, error)
  UserLoginHandler(user models.UserLogin) (models.TokenUsers, error)
  GetUserDetails(id int) (models.UserDetailsResponse, error)
  EditUserDetails(id int, details models.EditUserDetails) error
  AddAddress(id int, address models.AddAddress) error
  GetAddressess(id int) ([]models.Address, error)
  EditAddress(id int, userid uint, address models.EditAddress) error
  DeleteAddress(addressID int, userID int) error
  GetWallet(userID int) (models.GetWallet, error)

}
