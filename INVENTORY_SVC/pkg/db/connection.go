package db

import (
	"database/sql"
	"fmt"

	"github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/config"
	"github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/utils/domain"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {

  connString:=fmt.Sprintf("host=%s user=%s password=%s ",cfg.DBHost,cfg.DBUser,cfg.DBPassword)

  db,err:=sql.Open("postgres",connString) 
  if err!=nil{
    fmt.Println("creating new Database")
    return &gorm.DB{},err
  }

  rows,err:=db.Query("SELECT 1 FROM pg_database WHERE datname = $1",cfg.DBName)
  if err!=nil{
    fmt.Println("Error checking if database exists")
    return &gorm.DB{},err
  } 

  if rows.Next() {
        rows.Close()
  }else{
    _,err:=db.Exec("CREATE DATABASE "+cfg.DBName)
    if err!=nil{
    fmt.Println("Error creating"+cfg.DBName)
      return &gorm.DB{},err
    }
        fmt.Println(cfg.DBName+" created")
  }

   fmt.Println("config in connection DB",cfg) 
    
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)

	DB, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
  
  if dbErr!=nil{
    return DB,dbErr
  }
  if err := DB.AutoMigrate(&domain.Category{}); err != nil {
		return DB, err
	}
  if err := DB.AutoMigrate(&domain.Inventories{}); err != nil {
		return DB, err
	}

  if err:= DB.AutoMigrate(&domain.CategoryOffers{}); err !=nil{
    return DB,err
  }

  if err:= DB.AutoMigrate(&domain.InventoryOffers{}); err !=nil{
    return DB,err
  }
  
  if err:= DB.AutoMigrate(&domain.Coupons{}); err !=nil{
    return DB,err
  }

  return   DB,nil
}


