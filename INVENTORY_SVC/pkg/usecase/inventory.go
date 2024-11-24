package usecase

import (
	"errors"
	"fmt"

	helper_interfaces "github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/helper"
	interfaces "github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/repository/interface"

	"github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/utils/models"
)

type InventoryUseCase struct{
  repository interfaces.InventoryRepository
   helper_interfaces.InventoryServiceHelper
  offerRepo  interfaces.OfferRepository
}

func NewInventoryUseCase(repo interfaces.InventoryRepository, h helper_interfaces.InventoryServiceHelper, off interfaces.OfferRepository) *InventoryUseCase {
  return &InventoryUseCase{
    repository: repo,
    InventoryServiceHelper:h,
    offerRepo:  off,
  }

}

func (Inv *InventoryUseCase) AddInventory(inventory models.AddInventory ) (models.InventoryResponse, error) {

  exists, err := Inv.repository.CheckCategoryID(inventory.CategoryID)
  if err != nil {
    return models.InventoryResponse{}, err
  }
  if !exists {
    return models.InventoryResponse{}, errors.New("category of this ID doesn't exist ")
  }

  exists, err = Inv.repository.CheckProduct(inventory.ProductName, inventory.Size)

  if err != nil {
    return models.InventoryResponse{}, err
  }
  if exists {
    errMsg := fmt.Sprintf("Product %s of Size: %s already exists", inventory.ProductName, inventory.Size)
    return models.InventoryResponse{}, errors.New(errMsg)
  }


  inventoryResponse, Err := Inv.repository.AddInventory(inventory)

  if Err != nil {
    return models.InventoryResponse{}, Err
  }
  return inventoryResponse, nil
}

func (Inv *InventoryUseCase) GetProductsForAdmin() ([]models.Inventories, error) {
  productDetails, err := Inv.repository.ListProducts()
  if err != nil {
    return []models.Inventories{}, err
  }

  //check if any offers are there

  for i, Product := range productDetails {

    // if the category id of these products are in offer table discount the price to new one

    CategoryDiscountRate, CategoryOffer, err := Inv.offerRepo.GetCategoryOfferDiscountPercentage(Product.CategoryID)
    if err != nil {
      return []models.Inventories{}, err
    }
    productDetails[i].Categoryoffer = CategoryOffer
    ProductDiscountRate, ProductOffer, err := Inv.offerRepo.GetInventoryOfferDiscountPercentage(int(Product.Product_ID))
    if err != nil {
      return []models.Inventories{}, err
    }
    productDetails[i].Productoffer = ProductOffer

    var discount float64
    DiscountRate := CategoryDiscountRate + ProductDiscountRate
    if DiscountRate > 0 {
      discount = (Product.Price * float64(DiscountRate)) / 100
    }

    //Discounted Price = Original Price - (Original Price * (Discount Percentage / 100))

    Product.DiscountedPrice = Product.Price - discount

    fmt.Println("discounted Price", Product.DiscountedPrice)
    fmt.Println("ORginal Price", Product.Price)

    productDetails[i].DiscountedPrice = Product.DiscountedPrice
  }
  return productDetails, nil
}

func (Inv *InventoryUseCase) GetProductsForUsers() ([]models.Inventories, error) {

  productDetails, err := Inv.repository.ListProducts()
  if err != nil {
    return []models.Inventories{}, err
  }

  //check if any offers are there

  for i, Product := range productDetails {

    // if the category id of these products are in offer table discount the price to new one
    CategoryDiscountRate, categoryOffer, err := Inv.offerRepo.GetCategoryOfferDiscountPercentage(Product.CategoryID)
    if err != nil {
      return []models.Inventories{}, err
    }
    productDetails[i].Categoryoffer = categoryOffer
    ProductDiscountRate, productOffer, err := Inv.offerRepo.GetInventoryOfferDiscountPercentage(int(Product.Product_ID))
    if err != nil {
      return []models.Inventories{}, err
    }

    productDetails[i].Productoffer = productOffer

    var discount float64
    DiscountRate := CategoryDiscountRate + ProductDiscountRate
    if DiscountRate > 0 {
      discount = (Product.Price * float64(DiscountRate)) / 100
    }

    //Discounted Price = Original Price - (Original Price * (Discount Percentage / 100))

    Product.DiscountedPrice = Product.Price - discount

    // fmt.Println("discounted Price", Product.DiscountedPrice)
    // fmt.Println("ORginal Price", Product.Price)

    productDetails[i].DiscountedPrice = Product.DiscountedPrice
  }

  return productDetails, nil
}

func (Inv *InventoryUseCase) DeleteInventory(pid int) error {

  err := Inv.repository.DeleteInventory(pid)
  if err != nil {
    return err
  }

  return nil
}

func (Inv *InventoryUseCase) EditInventory(pid int, model models.EditInventory) error {

  err := Inv.repository.EditInventory(pid, model)

  if err != nil {
    return err
  }

  return nil

}

func (Inv *InventoryUseCase) SearchProducts(pdtName string) ([]models.Inventories, error) {

  productDetails, err := Inv.repository.SearchProducts(pdtName)
  if err != nil {
    return []models.Inventories{}, err
  }

  //check if any offers are there

  for i, Product := range productDetails {

    // if the category id of these products are in offer table discount the price to new one

    CategoryDiscountRate, CategoryOffer, err := Inv.offerRepo.GetCategoryOfferDiscountPercentage(Product.CategoryID)
    if err != nil {
      return []models.Inventories{}, err
    }

    productDetails[i].Categoryoffer = CategoryOffer
    ProductDiscountRate, ProductOffer, err := Inv.offerRepo.GetInventoryOfferDiscountPercentage(int(Product.Product_ID))
    if err != nil {
      return []models.Inventories{}, err
    }

    productDetails[i].Productoffer = ProductOffer
    var discount float64
    DiscountRate := CategoryDiscountRate + ProductDiscountRate
    if DiscountRate > 0 {
      discount = (Product.Price * float64(DiscountRate)) / 100
    }

    //Discounted Price = Original Price - (Original Price * (Discount Percentage / 100))

    Product.DiscountedPrice = Product.Price - discount

    fmt.Println("discounted Price", Product.DiscountedPrice)
    fmt.Println("ORginal Price", Product.Price)

    productDetails[i].DiscountedPrice = Product.DiscountedPrice
  }

  return productDetails, nil

}
