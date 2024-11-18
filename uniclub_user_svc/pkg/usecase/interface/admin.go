package interfaces

import (
	"github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/utils/domain"
	"github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/utils/models"
)

type AdminUseCase interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error)
	GetUsers() ([]models.UserDetailsAtAdmin, error)
	BlockUser(id int) error
	UnBlockUser(id int) error
	ChangePassword(passChange models.AdminPasswordChange, id int) error
}
