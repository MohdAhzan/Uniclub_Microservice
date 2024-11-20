package di

import (
	server "github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/api"
	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/api/handler"
	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/api/middleware"
	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/client"
	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/config"
)


  

func InitializeAPI(cfg config.Config)(*server.ServerHTTP,error){
  
  middleware.CfgHelper(cfg)
  usersvcClient:=client.NewUserServiceClient(cfg)
  usersvcHandler:=handler.NewUserServiceHandler(usersvcClient)
  serverHTTP:=server.NewServerHTTP(usersvcHandler)

    return serverHTTP,nil
}