package server

import (
	"fmt"
	"log"

	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/api/handler"
	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
  engine *gin.Engine
}

func NewServerHTTP(usersvcHandler *handler.UserServiceHandler) *ServerHTTP {


  router := gin.New()

  router.Use(gin.Logger())

  // ADMIN ROUTES


  router.POST("/admin/login", usersvcHandler.AdminLoginHandler)

  router.Use(middleware.AdminAuthMiddleware)
  {

    router.PUT("/change_password", usersvcHandler.ChangeAdminPassword)

    userManagement := router.Group("/users")
    {
      userManagement.GET("", usersvcHandler.GetUsers)
      userManagement.PUT("/block", usersvcHandler.BlockUser)
      userManagement.PUT("/unblock", usersvcHandler.UnBlockUser)
    }
  }

  // CLIENT ROUTES 

  router.POST("/user/signup", usersvcHandler.UserSignUp)
  router.POST("/user/login", usersvcHandler.UserLoginHandler)

  router.Use(middleware.UserAuthMiddleware)
  {

    profile := router.Group("/profile")
    {
      profile.GET("/details", usersvcHandler.GetUserDetails)
      profile.GET("/address", usersvcHandler.GetAddressess)
      profile.POST("/address", usersvcHandler.AddAddressess)
      profile.DELETE("/address", usersvcHandler.DeleteAddress)
      edit := profile.Group("/edit")
      {
        edit.PUT("/account", usersvcHandler.EditUserDetails)
        edit.PUT("/address", usersvcHandler.EditAddress)
      }
    }
  }
  return &ServerHTTP{engine: router}
}


func (s *ServerHTTP) Start() {
  log.Printf("starting server on 7000")
  err := s.engine.Run(":7000")
  if err != nil {
    fmt.Println("asldkjf;lksdjflsldjfljsad;lkfjlsdjflj",err)
    log.Printf("error while starting the server")
  }
}
