package client

import (
	"context"
	"fmt"

	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/client/interfaces"
	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/config"
	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/pb/usersvc"
	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/utils/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


type usersvcClient struct {
  Client usersvc.UserServiceClient
}


func NewUserServiceClient(cfg config.Config) interfaces.UserServiceClient {

  fmt.Println("GRPC CLIENT SERVICE URL",cfg.UserSvcUrl)
  grpcConnection, err := grpc.NewClient(cfg.UserSvcUrl,grpc.WithTransportCredentials(insecure.NewCredentials()))
 
  if err != nil {
    fmt.Println("Could not connect", err)
  }

  grpcClient := usersvc.NewUserServiceClient(grpcConnection)

  return &usersvcClient{
    Client: grpcClient,
  }

}

func (c *usersvcClient) UserSignup(user models.UserDetails, refCode string) (models.TokenUsers, error) {


  resData, err := c.Client.UserSignup(context.Background(), &usersvc.UserSignupRequest{
    UserModel: &usersvc.SignupModel{

      Name: user.Name,
      Email:     user.Email,
      Phone:     user.Phone,
      Password:  user.Password,
      ConfirmPassword: user.ConfirmPassword ,

    },
    ReferallCode: refCode,
  })
  if err != nil {
    return models.TokenUsers{}, err
  }

  return models.TokenUsers{
    Users: models.UserDetailsResponse{

      Id: int(resData.ResponseData.UserID),
      Name: resData.ResponseData.Name,
      Email: resData.ResponseData.Email,
      Phone: resData.ResponseData.Phone,
      ReferralID: resData.ResponseData.ReferralID,
    },
    AccessToken: resData.AccessToken,
    RefreshToken: resData.RefreshToken,
  }, nil
}

func (u *usersvcClient)UserLoginHandler(user models.UserLogin) (models.TokenUsers, error){

  resData, err := u.Client.UserLoginHandler(context.Background(), &usersvc.UserLoginRequest{
    Email:    user.Email,
    Password: user.Password,
  })
  if err != nil {
    return models.TokenUsers{}, err
  }

  return models.TokenUsers{
    Users: models.UserDetailsResponse{
      Id: int(resData.Id),
      Name: resData.Name,
      Email: resData.Email,
      Phone: resData.Phone,
      ReferralID: resData.ReferralID,

    },
    AccessToken: resData.AccessToken,
    RefreshToken: resData.RefreshToken,
  }, nil
}


func (u *usersvcClient)	AdminLoginHandler(adminDetails models.AdminLogin) (models.TokenAdmin, error){

  resData, err := u.Client.AdminLogin(context.Background(), &usersvc.AdminLoginRequest{
    Email:    adminDetails.Email,
    Password: adminDetails.Password,
  })
  if err != nil {
    return models.TokenAdmin{}, err
  }

  return models.TokenAdmin{
    Admin: models.AdminDetailsResponse{
      ID: int(resData.AdminResponseData.ID), 
      Name: resData.AdminResponseData.Name,
      Email: resData.AdminResponseData.Email,

    },
    AccessToken: resData.AccessToken,
  }, nil
}


func (u *usersvcClient)  GetUserDetails(id int) (models.UserDetailsResponse, error){

  resData, err := u.Client.GetUserDetails(context.Background(), &usersvc.GetUserDetailsRequest{

    UserId: int64(id),

  })
  if err != nil {
    return models.UserDetailsResponse{}, err
  }

  return models.UserDetailsResponse{

    Id: int(resData.ResData.Id),
    Name: resData.ResData.Name,
    Email: resData.ResData.Email,
    Phone: resData.ResData.Phone,
    ReferralID: resData.ResData.ReferralID,
  }, nil
}


func (u *usersvcClient)EditUserDetails(id int, details models.EditUserDetails) error{


  _,err :=u.Client.EditUserDetails(context.Background(), &usersvc.EditUserDetailsRequest{
    UserId: int64(id) ,
    Details: &usersvc.EditUserDetailsModel{
    Name: details.Name,
    Email: details.Email,
    Phone: details.Phone, 
    Password: details.Password,
  }})
  if err!=nil{
    return err
  }
  return nil 
}


func (u *usersvcClient)GetUsers()([]models.UserDetailsAtAdmin,error){


  res,err :=u.Client.GetUsers(context.Background(), &usersvc.GetUsersRequest{})
  if err!=nil{
    return []models.UserDetailsAtAdmin{},err
  }
  
    var users []models.UserDetailsAtAdmin

    for _, user := range res.Details {
        users = append(users, models.UserDetailsAtAdmin{
            Id:      uint(user.Id),
            Name:    user.Name,
            Email:   user.Email,
            Phone:   user.Phone,
            Blocked: user.Blocked,
        })
    }

    return users, nil
}


func (u *usersvcClient)  AddAddress(id int, address models.AddAddress) error{

  _,err :=u.Client.AddAddress(context.Background(),&usersvc.AddAddressReq{
     UserID: int64(id),
    Details: &usersvc.Address{
      Name: address.Name ,       
      Address: address.Address,
      LandMark: address.LandMark,
      City: address.City,
      Pincode: address.Pincode,
      State: address.State,
      Phone: address.Phone,

    }, 
  })
  if err!=nil{
    return err
  }

    return nil
}


func (u *usersvcClient)  BlockUser(id int) error{

  _,err :=u.Client.BlockUser(context.Background(),&usersvc.BlockRequest{
    Id: int64(id),
  }) 
  if err!=nil{
    return err
  }

    return nil
}

func (u *usersvcClient)  UnBlockUser(id int) error{

  _,err :=u.Client.UnBlockUser(context.Background(),&usersvc.UnBlockRequest{
    Id: int64(id),
  }) 
  if err!=nil{
    return err
  }

    return nil

}


func (u *usersvcClient)ChangeAdminPassword(passChange models.AdminPasswordChange, id int) error{

  _,err :=u.Client.ChangeAdminPassword(context.Background(),&usersvc.ChangeAdminPasswordRequest{
    Details: &usersvc.ChangeAdminPassword{
      CurrentPassword: passChange.CurrentPassword,
      NewPassword: passChange.NewPassword,
      ConfirmPassword: passChange.ConfirmPassword,
    },
     Id: int64(id),

  },

    ) 
  if err!=nil{
    return err
  }

    return nil
}

func (u *usersvcClient)DeleteAddress(addressID int, userID int) error{

  _,err :=u.Client.DeleteAddress(context.Background(),&usersvc.DeleteAddressReq{
        AddressId: int64(addressID),
        UserId: int64(userID), 
  },
    
    ) 
  if err!=nil{
    return err
  }

    return nil
}

func (u *usersvcClient)  EditAddress(id int, userid uint, address models.EditAddress) error{

  _,err :=u.Client.EditAddress(context.Background(),&usersvc.EditAddressReq{
    ID: int64(id),
    Userid: int64(userid),
    EditAddressDetails: &usersvc.Address{
        Name: address.Name,
      Address: address.Address,
      LandMark: address.LandMark,
      City: address.City,
      Pincode: address.Pincode,
      Default: address.Default,
    },
  },
    
    ) 
  if err!=nil{
    return err
  }

    return nil
}


func (u *usersvcClient) GetAddressess(id int) ([]models.Address, error){

  res,err :=u.Client.GetAddress(context.Background(),&usersvc.GetAddressReq{
      UserID: int64(id),    
  })
  if err!=nil{
    return []models.Address{},err
  }
  
  var addresses []models.Address

    for _, address := range res.Details {
        addresses = append(addresses, models.Address{
      UserID: int(address.UserID),
        Name: address.Name,
      LandMark: address.LandMark,
      City: address.City,
      Pincode: address.Pincode,
      State: address.State,
      Phone: address.Phone, 

        })
    }

    return addresses, nil  
}

///////////////////////////

func (u *usersvcClient) GetWallet(userID int) (models.GetWallet, error){

  res,err :=u.Client.GetWallet(context.Background(),&usersvc.GetWalletReq{
    ID: int64(userID),
  })
  if err!=nil{
    return models.GetWallet{},err
  }
  

    return models.GetWallet{
    UserID: int(res.UserID) ,
    Username:res.Username ,
    TotalAmount: res.TotalAmount,

  }, nil  
}
