package handler

import (
	"net/http"

	response "github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/utils/Response"
	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/utils/models"
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/domain"
	"github.com/gin-gonic/gin"
)

func (cat *InventoryServiceHandler) AddCategory(c *gin.Context) {
	var category domain.Category
	if err := c.BindJSON(&category); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	CategoryResponse, err := cat.GrpcInvClient.AddCategory(category.Category)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the Category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added Category", CategoryResponse, nil)
	c.JSON(http.StatusOK, successRes)

}

func (Cat *InventoryServiceHandler) GetCategory(c *gin.Context) {

	categories, err := Cat.GrpcInvClient.GetCategories()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "couldn't get categories", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully got all categories", categories, nil)
	c.JSON(http.StatusOK, successRes)

}
func (cat *InventoryServiceHandler) UpdateCategory(c *gin.Context) {

	var update models.Rename

	if err := c.BindJSON(&update); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	categories, err := cat.GrpcInvClient.UpdateCategory(update.Current, update.New)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't update the  category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfullly updated the categories", categories, nil)
	c.JSON(http.StatusOK, successRes)

}

func (cat *InventoryServiceHandler) DeleteCategory(c *gin.Context) {
	CategoryID := c.Query("id")
	err := cat.GrpcInvClient.DeleteCategory(CategoryID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't delete the category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	succesRes := response.ClientResponse(http.StatusOK, "successfully deleted category", nil, nil)
	c.JSON(http.StatusOK, succesRes)
}

