package domain

import "github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/utils/models"

type Admin struct {
	ID       uint   `json:"id" gorm:"unique;not null"`
	Name     string `json:"name" gorm:"validate:required"`
	Email    string `json:"email" gorm:"validate:required"`
	Password string `json:"password" gorm:"validate:required"`
}

type TokenAdmin struct {
	Admin       models.AdminDetailsResponse
	AccessToken string
}
