package interfaces

import (
	"github.com/MohdAhzan/Uniclub_ecommerce_Microservice_project/pkg/utils/domain"
	"github.com/MohdAhzan/Uniclub_ecommerce_Microservice_project/pkg/utils/models"
)

type AdminUseCase interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error)
	GetUsers() ([]models.UserDetailsAtAdmin, error)
	BlockUser(id int) error
	UnBlockUser(id int) error
	ChangePassword(passChange models.AdminPasswordChange, id int) error
}
