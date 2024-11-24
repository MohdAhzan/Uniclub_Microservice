package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/client/interfaces"
	response "github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/utils/Response"
	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/utils/models"
	"github.com/gin-gonic/gin"
)

type UserServiceHandler struct{
  
  GrpcClient interfaces.UserServiceClient

}

func NewUserServiceHandler (usersvcClient interfaces.UserServiceClient)*UserServiceHandler{

    
  return &UserServiceHandler{
    GrpcClient: usersvcClient,
  }

}


func (s *UserServiceHandler)AdminLoginHandler(c *gin.Context){

   	var adminDetails models.AdminLogin
	if err := c.BindJSON(&adminDetails); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in the correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	admin, err := s.GrpcClient.AdminLoginHandler(adminDetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	c.Set("Access", admin.AccessToken)
	// c.Set("Refresh", admin.RefreshToken)

	successRes := response.ClientResponse(http.StatusOK, "Admin authenticated succesfully", admin, nil)
	c.JSON(http.StatusOK, successRes)
}

func (s *UserServiceHandler) GetUsers(c *gin.Context) {

	users, err := s.GrpcClient.GetUsers()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't retrieve details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully retrived the users", users, nil)
	c.JSON(http.StatusOK, successRes)
}





func (ad *UserServiceHandler) BlockUser(c *gin.Context) {
	id := c.Query("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		response.ClientResponse(http.StatusBadRequest, "error string conversion", nil, err.Error())
	}
	err = ad.GrpcClient.BlockUser(userID)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't block user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}

	successRes := response.ClientResponse(http.StatusOK, "successfully blocked the user ", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ad *UserServiceHandler) UnBlockUser(c *gin.Context) {
	id := c.Query("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		response.ClientResponse(http.StatusBadRequest, "error string conversion", nil, err.Error())
	}
	err = ad.GrpcClient.UnBlockUser(userID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't block user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRess := response.ClientResponse(http.StatusOK, "successfully unblocked the user", nil, nil)
	c.JSON(http.StatusOK, successRess)

}


func (ad *UserServiceHandler) ChangeAdminPassword(c *gin.Context) {

	var adminPassChange models.AdminPasswordChange

	if err := c.BindJSON(&adminPassChange); err != nil {

		errRes := response.ClientResponse(http.StatusBadRequest, "error BindingJson Invalid Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)

		return
	}
	id, exist := c.Get("id")
	fmt.Println("LOGG ADMIN JWT ID", id)

	if !exist {

		errRes := response.ClientResponse(http.StatusBadRequest, "error getting admin Id", nil, fmt.Errorf("No admin id exist"))
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	err := ad.GrpcClient.ChangeAdminPassword(adminPassChange, id.(int))
	if err != nil {

		errRes := response.ClientResponse(http.StatusBadRequest, "Error changing your password", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}

	successRes := response.ClientResponse(http.StatusOK, "successfully changed your password", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

