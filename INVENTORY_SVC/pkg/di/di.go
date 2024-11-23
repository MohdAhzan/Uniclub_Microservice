package di

import (
	"fmt"

	server "github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/api"
	"github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/repository"

	"github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/api/services"
	"github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/config"
	"github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/db"
	"github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/helper"
	"github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.ServerGRPC, error) {
	gormDB, err :=db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
  
  fmt.Println("gorm ",gormDB) 
  
  helper:=helper.NewHelper(cfg)
  categoryRepository:=repository.NewCategoryRepository(gormDB)
  inventoryRepository := repository.NewInventoryRepository(gormDB)
  offerRepository:=repository.NewOfferRepository(gormDB)
  couponRepository:=repository.NewCouponRepository(gormDB)
  categoryUsecase:=usecase.NewCategoryUseCase(categoryRepository)
  inventoryUsecase:=usecase.NewInventoryUseCase(inventoryRepository,*helper,offerRepository)
  offerUsecase:=usecase.NewOfferUseCase(offerRepository,categoryRepository,inventoryRepository)
  couponUsecase:=usecase.NewCouponUseCase(couponRepository)
	inventoryuServer := services.NewInventoryServer(inventoryUsecase,categoryUsecase,offerUsecase,couponUsecase,*helper)

  serverGRPC,err := server.NewGrpcServer(cfg,inventoryuServer)
  if err!=nil{
    return nil,err
  }
  

	return serverGRPC, nil

}
