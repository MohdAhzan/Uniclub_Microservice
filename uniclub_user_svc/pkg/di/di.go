package di

import (
	"fmt"

	server "github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/api"

	"github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/api/services"
	"github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/config"
	"github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/db"
	"github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/helper"
	"github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/repository"
	"github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.ServerGRPC, error) {
	gormDB, err :=db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
  
  fmt.Println("gorm ",gormDB) 
  
  helper:=helper.NewHelper(cfg)
  userRepository:=repository.NewUserRepository(gormDB)
	adminRepository := repository.NewAdminRepository(gormDB)
  adminUsecase:=usecase.NewAdminUsecase(adminRepository,helper,userRepository)
  userUsecase:=usecase.NewUserUseCase(userRepository,cfg,helper)
	userServer := services.NewUserServer(adminUsecase,userUsecase)

  serverGRPC,err := server.NewGrpcServer(cfg,userServer)
  if err!=nil{
    return nil,err
  }
  

	return serverGRPC, nil

}
