package server

import (
	"log"

	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/api/handler"
	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
  engine *gin.Engine
}

func NewServerHTTP(usersvcHandler *handler.UserServiceHandler, inventorysvcHandler *handler.InventoryServiceHandler) *ServerHTTP {


  router := gin.New()

  router.Use(gin.Logger())

  // ADMIN ROUTES


  router.POST("/admin/login", usersvcHandler.AdminLoginHandler)

  router.Use(middleware.AdminAuthMiddleware)
  {


    admin:=router.Group("/admin")
    {
      router.PUT("/change_password", usersvcHandler.ChangeAdminPassword)

      userManagement := admin.Group("/users")
      {
        userManagement.GET("", usersvcHandler.GetUsers)
        userManagement.PUT("/block", usersvcHandler.BlockUser)
        userManagement.PUT("/unblock", usersvcHandler.UnBlockUser)
      }
      categorymanagement := admin.Group("/category")
      {
        categorymanagement.GET("", inventorysvcHandler.GetCategory)
        categorymanagement.POST("", inventorysvcHandler.AddCategory)
        categorymanagement.PUT("", inventorysvcHandler.UpdateCategory)
        categorymanagement.DELETE("", inventorysvcHandler.DeleteCategory)

      }

      productmanagement := admin.Group("/products")
      {
        productmanagement.POST("", inventorysvcHandler.AddInventory)
        productmanagement.GET("", inventorysvcHandler.GetProductsForAdmin)
        productmanagement.DELETE("", inventorysvcHandler.DeleteInventory)
        productmanagement.PUT("/:id/edit_details", inventorysvcHandler.EditInventoryDetails)
      }

      offerManagment := admin.Group("/offers")
      {
        offerManagment.POST("/category", inventorysvcHandler.AddCategoryOffer)
        offerManagment.GET("/category", inventorysvcHandler.GetAllCategoryOffers)
        offerManagment.PUT("/category", inventorysvcHandler.EditCategoryOffer)
        offerManagment.DELETE("/category", inventorysvcHandler.ValidorInvalidCategoryOffers)

        offerManagment.POST("/product", inventorysvcHandler.AddInventoryOffer)
        offerManagment.GET("/product", inventorysvcHandler.GetInventoryOffers)
        offerManagment.PUT("/product", inventorysvcHandler.EditInventoryOffer)
        offerManagment.DELETE("/product", inventorysvcHandler.ValidorInvalidInventoryOffers)
      }
    }

  }

  // CLIENT ROUTES 

  router.POST("/user/signup", usersvcHandler.UserSignUp)
  router.POST("/user/login", usersvcHandler.UserLoginHandler)

  router.Use(middleware.UserAuthMiddleware)
  {

    user:=router.Group("/user")
    {

      profile := user.Group("/profile")
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

  }
  return &ServerHTTP{engine: router}
}


func (s *ServerHTTP) Start() {
  log.Printf("starting server on 7000")
  err := s.engine.Run(":7000")
  if err != nil {
    log.Printf("error while starting the server \n %v \n",err)
  }
}
