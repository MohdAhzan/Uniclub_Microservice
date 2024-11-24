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

type InventoryServiceHandler struct{
  
  GrpcInvClient interfaces.InventoryServiceClient

}


func NewInventoryServiceHandler (invsvcClient interfaces.InventoryServiceClient)*InventoryServiceHandler{

    
  return &InventoryServiceHandler{
    GrpcInvClient: invsvcClient,
  }

}

//Category Handlers.....

//Inventory handlers().....

func (Inv *InventoryServiceHandler) AddInventory(c *gin.Context) {

	var inventory models.AddInventory


	CategoryID, err := strconv.Atoi(c.Request.FormValue("category_id"))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "form value error", nil, err.Error())
		c.JSON(400, errRes)
		return
	}

	ProductName := c.Request.FormValue("product_name")
	Size := c.Request.FormValue("size")
	Stock, err := strconv.Atoi(c.Request.FormValue("stock"))
	if err != nil {
		errRes := response.ClientResponse(400, "form value errror", nil, err.Error())
		c.JSON(400, errRes)
		return
	}
	Price, err := strconv.Atoi(c.Request.FormValue("price"))

	fmt.Println("price", Price)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "form value error", nil, err.Error())
		c.JSON(400, errRes)
		return
	}
	if Price < 0 {
		errRes := response.ClientResponse(400, "form value error", nil, "Invalid Price")
		c.JSON(400, errRes)
		return
	}

	inventory.CategoryID = CategoryID
	inventory.ProductName = ProductName
	inventory.Size = Size
	inventory.Stock = Stock
	inventory.Price = float64(Price)


    

	inventoryResponse, Err := Inv.GrpcInvClient.AddInventory(inventory)
	if Err != nil {
		errRes := response.ClientResponse(400, "error adding products to inventory", nil, Err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(200, "successfully added Inventory", inventoryResponse, nil)
	c.JSON(200, successRes)

}

func (Inv *InventoryServiceHandler) GetProductsForAdmin(c *gin.Context) {

	productDetails, err := Inv.GrpcInvClient.GetProductsForAdmin()

	if err != nil {
		errRes := response.ClientResponse(400, "couldnt get product details for admin", nil, err.Error())
		c.JSON(400, errRes)
		return
	}
	successRes := response.ClientResponse(200, "successfully retrieved product details", productDetails, nil)
	c.JSON(200, successRes)
}

func (Inv *InventoryServiceHandler) GetProductsForUsers(c *gin.Context) {

	///

	///
	productDetails, err := Inv.GrpcInvClient.GetProductsForUsers()

	if err != nil {
		errRes := response.ClientResponse(400, "couldnt get product details for admin", nil, err.Error())
		c.JSON(400, errRes)
		return
	}
	successRes := response.ClientResponse(200, "successfully retrieved product details", productDetails, nil)
	c.JSON(200, successRes)
}

func (Inv *InventoryServiceHandler) DeleteInventory(c *gin.Context) {

	product_id := c.Query("id")

	pid, err := strconv.Atoi(product_id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error converting to int", nil, err.Error())
		c.JSON(400, errRes)
		return
	}
	Err := Inv.GrpcInvClient.DeleteInventory(pid)
	if Err != nil {
		errRes := response.ClientResponse(400, "fields provided are in wrong format", nil, Err.Error())
		c.JSON(400, errRes)
		return
	}
	successRes := response.ClientResponse(200, "successfully deleted the inventory", nil, nil)
	c.JSON(200, successRes)
}

func (inv *InventoryServiceHandler) EditInventoryDetails(c *gin.Context) {

	productID := c.Query("id")

	pid, err := strconv.Atoi(productID)
	if err != nil {
		errRes := response.ClientResponse(400, "error converting the id", nil, err.Error())
		c.JSON(400, errRes)
		return
	}
	var model models.EditInventory

	Err := c.BindJSON(&model)
	if Err != nil {
		errRes := response.ClientResponse(400, "error binding model", nil, Err.Error())
		c.JSON(400, errRes)
		return
	}

	eRR := inv.GrpcInvClient.EditInventory(pid, model)
	if eRR != nil {
		errRes := response.ClientResponse(400, "error editing product check category id", nil, eRR.Error())
		c.JSON(400, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully edited product", nil, nil)
	c.JSON(200, successRes)
}

func (inv *InventoryServiceHandler) SearchProducts(c *gin.Context) {

	pdtName := c.Query("search")

	searchedPdts, err := inv.GrpcInvClient.SearchProducts(pdtName)

	if err != nil {
		errRes := response.ClientResponse(400, "failed to Search Products", nil, err.Error())
		c.JSON(400, errRes)
		return
	}

	successRes := response.ClientResponse(200, "successfully searched the product", searchedPdts, nil)
	c.JSON(200, successRes)
}

//Offer Handlers.....


