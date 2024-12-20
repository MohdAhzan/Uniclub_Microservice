package interfaces

import (
	"github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/utils/domain"
	"github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/utils/models"
)

type UserUseCase interface {
  UserSignup(user models.UserDetails, refCode string) (models.TokenUsers, error)
  UserLoginHandler(user models.UserLogin) (models.TokenUsers, error)
  GetUserDetails(id int) (models.UserDetailsResponse, error)
  EditUserDetails(id int, details models.EditUserDetails) error
  AddAddress(id int, address models.AddAddress) error
  GetAddressess(id int) ([]domain.Address, error)
  EditAddress(id int, userid uint, address models.EditAddress) error
  ChangePassword(id int, changePass models.ChangePassword) error
  DeleteAddress(addressID int, userID int) error
  GetWallet(userID int) (models.GetWallet, error)
}
