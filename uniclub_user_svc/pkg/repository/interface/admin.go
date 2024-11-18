package interfaces

import (
	"github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/utils/domain"
	"github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/utils/models"
)

type AdminRepository interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error)
	GetUsers() ([]models.UserDetailsAtAdmin, error)
	GetUserByID(id int) (domain.Users, error)
	UpdateBlockUserByID(user domain.Users) error
	GetAdminHashPassword(id int) (string, error)
	UpdateAdminPass(id int, NewPass string) error
}
