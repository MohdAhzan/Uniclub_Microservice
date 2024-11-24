package services

import (
	"context"
	"fmt"

	"github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/helper"
	"github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/pb"
	interfaces "github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/usecase/interface"
	"github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/utils/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type InvServiceServer struct {
  categoryUsecase interfaces.CategoryUseCase
  invUsecase  interfaces.InventoryUseCase
  offerUsecase  interfaces.OfferUsecase
  couponUsecase  interfaces.CouponUseCase
  h helper.InventoryServiceHelper
  // pb.UnimplementedUserServiceServie
  pb.UnimplementedInventoryServiceServer
}

func NewInventoryServer(inv interfaces.InventoryUseCase,cat  interfaces.CategoryUseCase ,off interfaces.OfferUsecase, coupon interfaces.CouponUseCase, helper helper.InventoryServiceHelper) pb.InventoryServiceServer {

  return &InvServiceServer{
    categoryUsecase: cat,
    invUsecase:  inv,
    offerUsecase: off,
    couponUsecase: coupon,
    h: helper ,
  }
}


func (u *InvServiceServer) AddCategory(ctx context.Context, req *pb.AddCategoryReq) (*pb.AddCategoryRes, error) {

  resData, err := u.categoryUsecase.AddCategory(req.CategoryName)

  if err != nil {
    return &pb.AddCategoryRes{
      Message: "failed to add category..", 
    },status.Error(codes.Internal,fmt.Sprintf("%v",err))

  }

  return &pb.AddCategoryRes{
    Category: &pb.Category{
      Id: int32(resData.ID),
      CategoryName: resData.Category,
    }, 
    Message : "successsfully added "+req.CategoryName+" category",
  }, nil
}

func (u *InvServiceServer)GetCategories(ctx context.Context ,req *pb.GetCategoriesReq)(*pb.GetCategoriesRes,error){

  resData,err:=u.categoryUsecase.GetCategories()
  if err!=nil{
    return &pb.GetCategoriesRes{
    },status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }
  var categories  []*pb.Category

  for _,cat:=range resData{

    categories = append(categories, &pb.Category{
      Id: int32(cat.ID),
      CategoryName: cat.Category,
    })
  }

  return &pb.GetCategoriesRes{
    Categories: categories ,

  },nil

}





func (u *InvServiceServer)UpdateCategory(ctx context.Context ,req *pb.UpdateCategoryReq)(*pb.UpdateCategoryRes,error){

  resData,err:=u.categoryUsecase.UpdateCategory(req.CurrentCategoryName,req.NewCategoryName)
  if err!=nil{
    return &pb.UpdateCategoryRes{
      Message: "failed updating category",
    },status.Error(codes.Internal,fmt.Sprintf("%v",err))

  }

  return &pb.UpdateCategoryRes{
    Category: &pb.Category{
      Id: int32(resData.ID),
      CategoryName: resData.Category,
    },
    Message: fmt.Sprintf("successfully updated Category"),
  },nil
}


func (u *InvServiceServer)DeleteCategory(ctx context.Context ,req *pb.DeleteCategoryReq)(*pb.DeleteCategoryRes,error){

  err:=u.categoryUsecase.DeleteCategory(req.CategoryId)

  if err!=nil{
    return &pb.DeleteCategoryRes{
      Message: "failed category deletiion" ,
    },
    status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }

  return &pb.DeleteCategoryRes{
    Message: "successfully delelted category" ,
  },nil
}

func (u *InvServiceServer)AddInventory(ctx context.Context, req *pb.AddInventoryReq)( *pb.AddInventoryRes, error  ) {
  resData,err:=u.invUsecase.AddInventory(models.AddInventory{
    CategoryID: int(req.CategoryId), 
    ProductName:req.ProductName,
    Size: req.Size,
    Stock: int(req.Stock),
    Price: req.Price, 
  })
   
  if err!=nil{
    return nil ,status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }
  return &pb.AddInventoryRes{
    ProductId: uint32(resData.Product_ID),
    Stock: resData.Stock,
  },nil

}




func (u *InvServiceServer)GetProductsForAdmin(ctx context.Context, req *pb.GetProductsForAdminReq)( *pb.GetProductsForAdminRes, error  ) {

  resData,err:=u.invUsecase.GetProductsForAdmin()
  if err!=nil{
    return nil ,status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }

  var invDatas []*pb.Inventory

  for  _,data:=range resData{

    invDatas = append(invDatas, &pb.Inventory{
      ProductId: uint32(data.Product_ID),
      CategoryId: int32(data.CategoryID),
      ProductName: data.ProductName,
      Size: data.Size,
      Stock: int32(data.Stock),
      Price: data.Price,
      IfPresentAtWishlist: data.IfPresentAtWishlist,
      IfPresentAtCart: data.IfPresentAtCart,
      CategoryOffer: data.Categoryoffer,
      ProductOffer: data.Productoffer,
      DiscountedPrice:data.DiscountedPrice ,

    })

  }


  return &pb.GetProductsForAdminRes{
    Inventories:invDatas,
  }  ,nil

}



func (u *InvServiceServer)GetProductsForUsers(ctx context.Context, req *pb.GetProductsForUsersReq)( *pb.GetProductsForUsersRes, error  ) {

  resData,err:=u.invUsecase.GetProductsForUsers()
  if err!=nil{

    return nil ,status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }

  var invDatas []*pb.Inventory

  for  _,data:=range resData{

    invDatas = append(invDatas, &pb.Inventory{
      ProductId: uint32(data.Product_ID),
      CategoryId: int32(data.CategoryID),
      ProductName: data.ProductName,
      Size: data.Size,
      Stock: int32(data.Stock),
      Price: data.Price,
      IfPresentAtWishlist: data.IfPresentAtWishlist,
      IfPresentAtCart: data.IfPresentAtCart,
      CategoryOffer: data.Categoryoffer,
      ProductOffer: data.Productoffer,
      DiscountedPrice:data.DiscountedPrice ,

    })
  }
  return &pb.GetProductsForUsersRes{
    Inventories: invDatas,
  },nil

}


func (u *InvServiceServer)DeleteInventory(ctx context.Context, req *pb.DeleteInventoryReq)( *pb.DeleteInventoryRes, error  ) {

  err:=u.invUsecase.DeleteInventory(int(req.ProductId))
  if err!=nil{
    return nil ,status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }

  return nil,nil

}
func (u *InvServiceServer)EditInventory(ctx context.Context, req *pb.EditInventoryReq)( *pb.EditInventoryRes, error  ) {

  err:=u.invUsecase.EditInventory(int(req.ProductId),models.EditInventory{
    CategoryID: int(req.CategoryId),
    ProductName: req.ProductName,
    Size: req.Size,
    Stock: int(req.Stock),
    Price: req.Price,

  })

  if err!=nil{
    return nil ,status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }
  return nil,nil

}


func (u *InvServiceServer)SearchProducts(ctx context.Context, req *pb.SearchProductsReq)( *pb.SearchProductsRes, error  ) {

  resData,err:=u.invUsecase.SearchProducts(req.ProductName)
  if err!=nil{
    return nil ,status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }

  var invDatas []*pb.Inventory

  for  _,data:=range resData{

    invDatas = append(invDatas, &pb.Inventory{
      ProductId: uint32(data.Product_ID),
      CategoryId: int32(data.CategoryID),
      ProductName: data.ProductName,
      Size: data.Size,
      Stock: int32(data.Stock),
      Price: data.Price,
      IfPresentAtWishlist: data.IfPresentAtWishlist,
      IfPresentAtCart: data.IfPresentAtCart,
      CategoryOffer: data.Categoryoffer,
      ProductOffer: data.Productoffer,
      DiscountedPrice:data.DiscountedPrice ,

    })
  }
  return &pb.SearchProductsRes{
    Inventories: invDatas,
  },nil
}



func (u *InvServiceServer)AddCategoryOffer(ctx context.Context, req *pb.AddCategoryOfferReq)( *pb.AddCategoryOfferRes, error  ) {

  validTime,err:=u.h.StringToTime(req.ValidTill)

  if err!=nil{
    return &pb.AddCategoryOfferRes{},status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }


  err=u.offerUsecase.AddCategoryOffer(models.AddCategoryOffer{
    CategoryID: int(req.CategoryId),
    OfferName: req.OfferName ,
    DiscountRate: req.DiscountRate,

    ValidTill: validTime,
  })

  if err!=nil{
    return nil ,status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }
  return nil,nil

}

func (u *InvServiceServer)GetAllCategoryOffers(ctx context.Context, req *pb.GetAllCategoryOffersReq)( *pb.GetAllCategoryOffersRes, error  ) {
 


  resData,err:=u.offerUsecase.GetAllCategoryOffers()
  if err!=nil{
    return nil ,status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }



    var offDatas []*pb.CategoryOffer

  for  _,data:=range resData{
    validTime:=u.h.TimeToString(data.ValidTill)
    createdat:=u.h.TimeToString(data.CreatedAt)

    offDatas= append(offDatas, &pb.CategoryOffer{
        Id: uint32(data.ID),
      CategoryId: uint32(data.CategoryID),
      OfferName: data.OfferName,
      DiscountRate: data.DiscountRate,
      CreatedAt: createdat,
      ValidTill:validTime,

    })
  }

  return nil,nil


}


func (u *InvServiceServer)EditCategoryOffer(ctx context.Context, req *pb.EditCategoryOfferReq)( *pb.EditCategoryOfferRes, error  ) {

err:=u.offerUsecase.EditCategoryOffer(float64(req.NewDiscount),int(req.CategoryId))
  if err!=nil{
    return nil ,status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }

  return &pb.EditCategoryOfferRes{
      Message: "successfully edited categoryoffer",
  }  ,nil
}

func (u *InvServiceServer)ValidorInvalidCategoryOffers(ctx context.Context ,req *pb.ValidorInvalidCategoryOffersReq)(*pb.ValidorInvalidCategoryOffersRes,error){

  err:=u.offerUsecase.ValidorInvalidCategoryOffers(req.Status,int(req.CategoryId))
  if err!=nil{
    return &pb.ValidorInvalidCategoryOffersRes{
    },status.Error(codes.Internal,fmt.Sprintf("error checking categoryoffers %v",err))

  }

  return &pb.ValidorInvalidCategoryOffersRes{
   Message: "validation succesfull",
  },nil
}


func (u *InvServiceServer)AddInventoryOffer(ctx context.Context ,req *pb.AddInventoryOfferReq)(*pb.AddInventoryOfferRes,error){

  validtill,err:=u.h.StringToTime(req.ValidTill)
  if err!=nil{
      
    return &pb.AddInventoryOfferRes{},status.Error(codes.Aborted,fmt.Sprintf("%v",err))
    
  }
      
  err=u.offerUsecase.AddInventoryOffer(models.AddInventoryOffer{
    InventoryID: int(req.InventoryId),
    OfferName:req.OfferName ,
    DiscountRate: req.DiscountRate,
     ValidTill: validtill,
  })

  if err!=nil{
    return &pb.AddInventoryOfferRes{
      Message: "error adding inventory offer",
    },status.Error(codes.Internal,fmt.Sprintf("%v",err))

  }

  return &pb.AddInventoryOfferRes{
   Message: "added inventory offer",
  },nil
}
func (u *InvServiceServer)GetInventoryOffers(ctx context.Context ,req *pb.GetInventoryOffersReq)(*pb.GetInventoryOffersRes,error){

      
  resData,err:=u.offerUsecase.GetInventoryOffers()
    if err!=nil{
    return nil ,status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }



    var offDatas []*pb.InventoryOffer

  for  _,data:=range resData{
    validTime:=u.h.TimeToString(data.ValidTill)
    createdat:=u.h.TimeToString(data.CreatedAt)

    offDatas= append(offDatas, &pb.InventoryOffer{
        Id: uint32(data.ID),
      InventoryId: uint32(data.InventoryID),
      ProductName: data.ProductName,
      OfferName: data.OfferName,
      DiscountRate: data.DiscountRate,
      CreatedAt: createdat,
      ValidTill:validTime,
      IsActive: data.IsActive,

    })
  }
  return nil,nil

}

func (u *InvServiceServer)EditInventoryOffer(ctx context.Context, req *pb.EditInventoryOfferReq)( *pb.EditInventoryOfferRes, error  ) {

err:=u.offerUsecase.EditInventoryOffer(float64(req.NewDiscount),int(req.InventoryId))
  if err!=nil{
    return nil ,status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }

  return &pb.EditInventoryOfferRes{
      Message: "successfully edited inventory offer",
  }  ,nil
}

func (u *InvServiceServer)ValidorInvalidInventoryOffers(ctx context.Context ,req *pb.ValidorInvalidInventoryOffersReq)(*pb.ValidorInvalidInventoryOffersRes,error){

  err:=u.offerUsecase.ValidorInvalidInventoryOffers(req.Status,int(req.InventoryId))
  if err!=nil{
    return &pb.ValidorInvalidInventoryOffersRes{
    },status.Error(codes.Internal,fmt.Sprintf("error checking inventory offers %v",err))

  }

  return &pb.ValidorInvalidInventoryOffersRes{
   Message: "validation succesfull",
  },nil
}
