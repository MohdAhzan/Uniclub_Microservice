package usecase

import (
	"errors"
	"fmt"

	helper_interface "github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/helper/interface"
	interfaces "github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/repository/interface"
	"github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/utils/domain"
	"github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/utils/models"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepository interfaces.AdminRepository
	userRepository  interfaces.UserRepository
	helper          helper_interface.UserServiceHelper 
}

func NewAdminUsecase(repo interfaces.AdminRepository, h helper_interface.UserServiceHelper, u interfaces.UserRepository) *adminUseCase{
	return &adminUseCase{
		adminRepository: repo,
    userRepository:  u,
		helper:          h,
	}
}

func (ad *adminUseCase) LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin,error)  {

	adminCompareDetails, err := ad.adminRepository.LoginHandler(adminDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}
  


	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetails.Password))
	if err != nil {
		return domain.TokenAdmin{}, err
	}
	var adminDetailsResponse models.AdminDetailsResponse

	err = copier.Copy(&adminDetailsResponse, &adminCompareDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	access, err := ad.helper.GenerateTokenAdmin(adminDetailsResponse)

	if err != nil {
		return domain.TokenAdmin{}, err
	}

	return domain.TokenAdmin{
		Admin:       adminDetailsResponse,
		AccessToken: access,
	}, nil
}

func (ad *adminUseCase) GetUsers() ([]models.UserDetailsAtAdmin, error) {
	users, err := ad.adminRepository.GetUsers()
	if err != nil {
		return []models.UserDetailsAtAdmin{}, errors.New("Error fetching UserDetails")
	}
	return users, nil
}

func (ad *adminUseCase) BlockUser(id int) error {
	user, err := ad.adminRepository.GetUserByID(id)
	if err != nil {
		return err
	}

	if user.Blocked {
		return errors.New("user already blocked")
	} else {
		user.Blocked = true
	}

	err = ad.adminRepository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}

	return nil
}

func (ad *adminUseCase) UnBlockUser(id int) error {
	user, err := ad.adminRepository.GetUserByID(id)
	if err != nil {
		return err
	}
	if user.Blocked {
		user.Blocked = false
	} else {
		return errors.New("user already unblocked")
	}

	err = ad.adminRepository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}
	return nil
}



func (ad *adminUseCase) ChangePassword(changePasswordDetails models.AdminPasswordChange, id int) error {

	if changePasswordDetails.NewPassword == changePasswordDetails.CurrentPassword {

		return fmt.Errorf("New Password is same as old one.Try again!!!")
	}

	hashedPass, err := ad.adminRepository.GetAdminHashPassword(id)
	if err != nil {
		return err
	}

	fmt.Println("hashed PAss", hashedPass)

	err = ad.helper.CompareHashAndPassword(hashedPass, changePasswordDetails.CurrentPassword)
	fmt.Println(changePasswordDetails.CurrentPassword, "currentPass")

	if err != nil {
		return fmt.Errorf("Incorrect Password Try again!!!")

	}

	if len(changePasswordDetails.NewPassword) < 3 {

		return fmt.Errorf("Password is too short")
	}

	if changePasswordDetails.NewPassword != changePasswordDetails.ConfirmPassword {

		return fmt.Errorf("Password mismatch Try again!!!")

	}

	newHashedPass, err := ad.helper.PasswordHashing(changePasswordDetails.NewPassword)
	if err != nil {
		return err
	}

	err = ad.adminRepository.UpdateAdminPass(id, newHashedPass)
	if err != nil {
		return err
	}

	return nil
}

