package interfaces

import (
	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/utils/models"
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/domain"
)

type InventoryServiceClient interface{

 	AddCategory(category string) (domain.Category, error)
	GetCategories() ([]domain.Category, error)
	UpdateCategory(current string, new string) (domain.Category, error)
	DeleteCategory(CategoryID string) error

	AddInventory(inventory models.AddInventory ) (models.InventoryResponse, error)
	GetProductsForAdmin() ([]models.Inventories, error)
	GetProductsForUsers() ([]models.Inventories, error)
	DeleteInventory(pid int) error
	EditInventory(pid int, model models.EditInventory) error
	SearchProducts(pdtName string) ([]models.Inventories, error)

	AddCategoryOffer(model models.AddCategoryOffer) error
	GetAllCategoryOffers() ([]domain.CategoryOffers, error)
	EditCategoryOffer(newDiscount float64, cID int) error
	ValidorInvalidCategoryOffers(status bool, cID int) error

	AddInventoryOffer(model models.AddInventoryOffer) error
	GetInventoryOffers() ([]models.GetInventoryOffers, error)
	EditInventoryOffer(newDiscount float64, InventoryID int) error
	ValidorInvalidInventoryOffers(status bool, inventoryID int) error




} 
