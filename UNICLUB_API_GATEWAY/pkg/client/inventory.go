package client

import (
	"context"
	"fmt"

	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/client/interfaces"
	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/config"
	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/pb/inventorysvc"
	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/utils"
	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/utils/models"
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)



type inventorySvcClient struct {
  Client  inventorysvc.InventoryServiceClient
  cfg config.Config
}


func NewInventoryServiceClient(cfg config.Config) interfaces.InventoryServiceClient {

  fmt.Println("GRPC CLIENT SERVICE URL",cfg.InventorySvcUrl)
  grpcConnection, err := grpc.NewClient(cfg.InventorySvcUrl,grpc.WithTransportCredentials(insecure.NewCredentials()))

  if err != nil {
    fmt.Println("Could not connect", err)
  }

  grpcClient := inventorysvc.NewInventoryServiceClient(grpcConnection)

  return &inventorySvcClient{
    Client: grpcClient,
    cfg: cfg,
  }

}

func(c *inventorySvcClient)AddCategory(category string) (domain.Category, error){
  res,err:=c.Client.AddCategory(context.Background(),&inventorysvc.AddCategoryReq{
    CategoryName: category,
  })
  if err!=nil{
    return domain.Category{},err
  }

 return domain.Category{
    ID: uint(res.Category.Id),
    Category: res.Category.CategoryName,
  },nil
}


func(c *inventorySvcClient)GetCategories() ([]domain.Category, error){

  res,err:=c.Client.GetCategories(context.Background(),&inventorysvc.GetCategoriesReq{})
  if err!=nil{
    return nil,err
  }

  var categoryRes []domain.Category

  for _,cat:= range res.Categories{

    categoryRes=append(categoryRes, domain.Category{
      ID: uint(cat.Id),
      Category: cat.CategoryName,
    })
  }

  return categoryRes,nil
}

func(c *inventorySvcClient)UpdateCategory(current string, new string) (domain.Category, error){

  res,err:=c.Client.UpdateCategory(context.Background(),&inventorysvc.UpdateCategoryReq{
    CurrentCategoryName: current,
    NewCategoryName: new,
  })
  if err!=nil{
    return domain.Category{},err
  }


  return domain.Category{
    ID: uint(res.Category.Id),
    Category: res.Category.CategoryName,
  },nil
}

func(c *inventorySvcClient)DeleteCategory(CategoryID string) error{

  _,err:=c.Client.DeleteCategory(context.Background(),&inventorysvc.DeleteCategoryReq{
    CategoryId: CategoryID,
  })
  if err!=nil{
    return err
  }

  return nil
}




func (c *inventorySvcClient) AddInventory(inventory models.AddInventory ) (models.InventoryResponse, error){


  res,err:=c.Client.AddInventory(context.Background(),&inventorysvc.AddInventoryReq{
    ProductId: uint32(inventory.Product_ID),
    CategoryId: int32(inventory.CategoryID),
    ProductName: inventory.ProductName,
    Size: inventory.Size,
    Stock: int32(inventory.Stock),
  })
  if err!=nil{
    return  models.InventoryResponse{},err
  }

  return models.InventoryResponse{
    Product_ID: uint(res.ProductId),
    Stock: res.Stock,
  },nil
}



func (c *inventorySvcClient)	GetProductsForAdmin()([]models.Inventories, error){

  resData,err:=c.Client.GetProductsForAdmin(context.Background(),&inventorysvc.GetProductsForAdminReq{})
  if err!=nil{
    return nil,err
  }

  var invRes []models.Inventories

  for _,inv:= range resData.Inventories{

        
    invRes=append(invRes, models.Inventories{
      Product_ID: uint(inv.CategoryId),
      CategoryID: int(inv.CategoryId),
      ProductName: inv.ProductName,
      Size: inv.Size,
      Stock: int(inv.Stock),
      Price:inv.Price,
      IfPresentAtWishlist: inv.IfPresentAtCart,
      IfPresentAtCart: inv.IfPresentAtCart, 
      Categoryoffer: inv.CategoryOffer,
      Productoffer: inv.ProductOffer,
      DiscountedPrice: inv.DiscountedPrice,
    })
  }

  return invRes,nil
    
    
}

func (c *inventorySvcClient)	GetProductsForUsers()([]models.Inventories, error){

  resData,err:=c.Client.GetProductsForUsers(context.Background(),&inventorysvc.GetProductsForUsersReq{})
  if err!=nil{
    return nil,err
  }

  var invRes []models.Inventories

  for _,inv:= range resData.Inventories{
    invRes=append(invRes, models.Inventories{
      Product_ID: uint(inv.CategoryId),
      CategoryID: int(inv.CategoryId),
      ProductName: inv.ProductName,
      Size: inv.Size,
      Stock: int(inv.Stock),
      Price:inv.Price,
      IfPresentAtWishlist: inv.IfPresentAtCart,
      IfPresentAtCart: inv.IfPresentAtCart, 
      Categoryoffer: inv.CategoryOffer,
      Productoffer: inv.ProductOffer,
      DiscountedPrice: inv.DiscountedPrice,
    })
  }
  return invRes,nil
    
    
}

func (c *inventorySvcClient)DeleteInventory(pid int) error {


  _,err:=c.Client.DeleteInventory(context.Background(),&inventorysvc.DeleteInventoryReq{
    ProductId: int32(pid),
  })
  if err!=nil{
    return  err
  }
  
  return nil
}

func (c *inventorySvcClient)EditInventory(pid int, model models.EditInventory) error{

  _,err:=c.Client.EditInventory(context.Background(),&inventorysvc.EditInventoryReq{
    ProductId: int32(pid),
    CategoryId:int32(model.CategoryID),
    ProductName: model.ProductName,
    Size: model.Size,
    Stock: int32(model.Stock),
    Price: model.Price,

  })
  if err!=nil{
    return  err
  }
  
  return nil
}

func (c *inventorySvcClient) SearchProducts(pdtName string) ([]models.Inventories, error){

  resData,err:=c.Client.SearchProducts(context.Background(),&inventorysvc.SearchProductsReq{
    ProductName: pdtName,
  })
  if err!=nil{
    return nil,err
  }

  var pdts []models.Inventories
  
  for _,pdt:= range resData.Inventories{
  
      pdts = append(pdts, models.Inventories{
        Product_ID: uint(pdt.ProductId),
      CategoryID: int(pdt.CategoryId),
      ProductName: pdt.ProductName,
      Size: pdt.Size,
      Stock: int(pdt.Stock),
      Price: pdt.Price,
      IfPresentAtWishlist: pdt.IfPresentAtCart,
      IfPresentAtCart: pdt.IfPresentAtCart,
      Categoryoffer: pdt.CategoryOffer,
      Productoffer: pdt.ProductOffer,
      DiscountedPrice: pdt.DiscountedPrice,

    })
       
  }
    
  return pdts, nil
}



func(c *inventorySvcClient)AddCategoryOffer(model models.AddCategoryOffer) error{

  validtillStr:=utils.TimeToString(model.ValidTill)

  _,err:=c.Client.AddCategoryOffer(context.Background(),&inventorysvc.AddCategoryOfferReq{
    CategoryId: int32(model.CategoryID),
    OfferName: model.OfferName,
    DiscountRate: model.DiscountRate ,
    ValidTill:validtillStr,
  })
  if err!=nil{
    return err
  }

  return nil
    
}
func (c *inventorySvcClient)EditCategoryOffer(newDiscount float64, cID int) error{

  _,err:=c.Client.EditCategoryOffer(context.Background(),&inventorysvc.EditCategoryOfferReq{
      NewDiscount: float32(newDiscount),
    CategoryId: int32(cID),
  })
  if err!=nil{
    return  err
  }
  
  return nil
}

func (c *inventorySvcClient)GetAllCategoryOffers() ([]domain.CategoryOffers, error){

  resData,err:=c.Client.GetAllCategoryOffers(context.Background(),&inventorysvc.GetAllCategoryOffersReq{})
  if err!=nil{
    return nil,err
  }

  var catoffs []domain.CategoryOffers
  
  for _,c:= range resData.CategoryOffers{
  
    validtillStr,err:=utils.StringToTime(c.ValidTill)
    if err!=nil{
      return nil,err
    }
    createdatStr,err:=utils.StringToTime(c.CreatedAt)
    if err!=nil{
      return nil,err
    }
      catoffs = append(catoffs, domain.CategoryOffers{
        ID: uint(c.Id),
      CategoryID: uint(c.CategoryId),
      ValidTill: validtillStr,
      DiscountRate: c.DiscountRate,
      OfferName:c.OfferName,
      IsActive: c.IsActive,
      CreatedAt: createdatStr,
    })
       
  }
    
  return catoffs, nil
}

func (c *inventorySvcClient)ValidorInvalidCategoryOffers(status bool, cID int) error{

  _,err:=c.Client.ValidorInvalidCategoryOffers(context.Background(),&inventorysvc.ValidorInvalidCategoryOffersReq{
    Status: status,
    CategoryId: int32(cID),

  })
  if err!=nil{
    return  err
  }
  
  return nil
}


func(c *inventorySvcClient)AddInventoryOffer(model models.AddInventoryOffer) error{

  validtillStr:=utils.TimeToString(model.ValidTill)

  _,err:=c.Client.AddInventoryOffer(context.Background(),&inventorysvc.AddInventoryOfferReq{
    InventoryId: int32(model.InventoryID),
    OfferName: model.OfferName,
    DiscountRate: model.DiscountRate ,
    ValidTill:validtillStr,
  })
  if err!=nil{
    return err
  }

  return nil
    
}
 
// 
// 
//
func (c *inventorySvcClient)EditInventoryOffer(newDiscount float64, InventoryID int) error{

  _,err:=c.Client.EditInventoryOffer(context.Background(),&inventorysvc.EditInventoryOfferReq{
      NewDiscount: float32(newDiscount),
    InventoryId: int32(InventoryID) ,
  })

  if err!=nil{
    return  err
  }
  
  return nil
}

func (c *inventorySvcClient)GetInventoryOffers() ([]models.GetInventoryOffers, error){

  resData,err:=c.Client.GetInventoryOffers(context.Background(),&inventorysvc.GetInventoryOffersReq{})
  if err!=nil{
    return nil,err
  }

  var catoffs []models.GetInventoryOffers
  
  for _,c:= range resData.InventoryOffers{
  
    validtillStr,err:=utils.StringToTime(c.ValidTill)
    if err!=nil{
      return nil,err
    }
    createdatStr,err:=utils.StringToTime(c.CreatedAt)
    if err!=nil{
      return nil,err
    }
      catoffs = append(catoffs, models.GetInventoryOffers{
        ID: uint(c.Id),
      InventoryID: uint(c.InventoryId),
      ProductName: c.ProductName ,
      ValidTill: validtillStr,
      DiscountRate: c.DiscountRate,
      OfferName:c.OfferName,
      IsActive: c.IsActive,
      CreatedAt: createdatStr,
    })
       
  }
    
  return catoffs, nil
}

func (c *inventorySvcClient) ValidorInvalidInventoryOffers(status bool, inventoryID int) error{

  _,err:=c.Client.ValidorInvalidInventoryOffers(context.Background(),&inventorysvc.ValidorInvalidInventoryOffersReq{
    Status: status,
    InventoryId: int32(inventoryID),
  })
  if err!=nil{
    return  err
  }
  
  return nil
}
