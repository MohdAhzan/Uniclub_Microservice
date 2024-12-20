package models

type Inventories struct {
	Product_ID          uint    `json:"product_id"`
	CategoryID          int     `json:"category_id"`
	ProductName         string  `json:"product_name"`
	Size                string  `json:"size" `
	Stock               int     `json:"stock"`
	Price               float64 `json:"price"`
	IfPresentAtWishlist bool    `json:"if_present_at_wishlist"`
	IfPresentAtCart     bool    `json:"if_present_at_cart"`
	Categoryoffer       string  `json:"category_offer"`
	Productoffer        string  `json:"product_offer"`
	DiscountedPrice     float64 `json:"discounted_price"`
}

type AddInventory struct {
	Product_ID  uint    `json:"product_id"`
	CategoryID  int     `json:"category_id"`
	ProductName string  `json:"product_name"`
	Size        string  `json:"size"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

type EditInventory struct {
	CategoryID  int     `json:"category_id"`
	ProductName string  `json:"product_name"`
	Size        string  `json:"size"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

type InventoryResponse struct {
	Product_ID uint   `json:"product_id"`
	Stock      string `json:"stock"`
}

type Rename struct {
	Current string `json:"current"`
	New     string `json:"new"`
}

