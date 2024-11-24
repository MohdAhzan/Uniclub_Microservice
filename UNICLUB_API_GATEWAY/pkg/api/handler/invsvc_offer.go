package handler

import (
	"net/http"
	"strconv"

	response "github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/utils/Response"
	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/utils/models"
	"github.com/gin-gonic/gin"
)

func (o *InventoryServiceHandler) AddCategoryOffer(c *gin.Context) {

	var offerModel models.AddCategoryOffer

	err := c.BindJSON(&offerModel)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error parsing json", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err = o.GrpcInvClient.AddCategoryOffer(offerModel)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error adding category offer", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}

	successRes := response.ClientResponse(http.StatusOK, "successfully added category Offer", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (o *InventoryServiceHandler) GetAllCategoryOffers(c *gin.Context) {

	catOffers, err := o.GrpcInvClient.GetAllCategoryOffers()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error fetching category offers", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}

	successRes := response.ClientResponse(http.StatusOK, "successfully fetched category Offers", catOffers, nil)
	c.JSON(http.StatusOK, successRes)

}

func (o *InventoryServiceHandler) EditCategoryOffer(c *gin.Context) {

	newDis := c.Query("new_discount")
	newDiscount, err := strconv.ParseFloat(newDis, 64)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in string conversion into float64", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	cid := c.Query("category_id")
	categoryID, err := strconv.Atoi(cid)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in string conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err = o.GrpcInvClient.EditCategoryOffer(newDiscount, categoryID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error editing category offer", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully edited category Offer", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (o *InventoryServiceHandler) ValidorInvalidCategoryOffers(c *gin.Context) {

	statusString := c.Query("status")
	status, err := strconv.ParseBool(statusString)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in string conversion into bool enter valid query parameter", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	cid := c.Query("category_id")
	categoryID, err := strconv.Atoi(cid)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in string conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err = o.GrpcInvClient.ValidorInvalidCategoryOffers(status, categoryID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error editing category offer", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully edited category Offer", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (o *InventoryServiceHandler) AddInventoryOffer(c *gin.Context) {

	var offerModel models.AddInventoryOffer
	err := c.BindJSON(&offerModel)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error parsing json", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err = o.GrpcInvClient.AddInventoryOffer(offerModel)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error adding product offer", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}

	successRes := response.ClientResponse(http.StatusOK, "successfully added product Offer", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (o *InventoryServiceHandler) GetInventoryOffers(c *gin.Context) {
	offerData, err := o.GrpcInvClient.GetInventoryOffers()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error fetching product offers", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully fetched all product offers", offerData, nil)
	c.JSON(http.StatusOK, successRes)
}

func (o *InventoryServiceHandler) EditInventoryOffer(c *gin.Context) {

	discount_rate := c.Query("new_discount")

	newDiscount, err := strconv.ParseFloat(discount_rate, 64)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in string conversion into float64", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	pid := c.Query("product_id")
	InventoryID, err := strconv.Atoi(pid)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in string conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err = o.GrpcInvClient.EditInventoryOffer(newDiscount, InventoryID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error editing product offer", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully edited product Offer", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (o *InventoryServiceHandler) ValidorInvalidInventoryOffers(c *gin.Context) {

	statusString := c.Query("status")
	status, err := strconv.ParseBool(statusString)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in string conversion into bool enter valid query parameter", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	pid := c.Query("product_id")
	inventoryID, err := strconv.Atoi(pid)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in string conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err = o.GrpcInvClient.ValidorInvalidInventoryOffers(status, inventoryID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error editing product offer", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully edited product Offer", nil, nil)
	c.JSON(http.StatusOK, successRes)
}


