package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/pb"
	interfaces "github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/usecase/interface"
	"github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/utils/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
  adminUsecase interfaces.AdminUseCase
  userUsecase  interfaces.UserUseCase
  pb.UnimplementedUserServiceServer
}

func NewUserServer(admnUsecase interfaces.AdminUseCase, userusecase interfaces.UserUseCase) pb.UserServiceServer {

  return &UserServer{
    adminUsecase: admnUsecase,
    userUsecase:  userusecase,
  }
}

func (u *UserServer) AdminLogin(ctx context.Context, req *pb.AdminLoginRequest) (*pb.AdminLoginResponse, error) {

  loginresponseData, err := u.adminUsecase.LoginHandler(models.AdminLogin{Email: req.Email, Password: req.Password})

  if err != nil {
    return &pb.AdminLoginResponse{
      AdminResponseData: nil,
      Httpstatus:        http.StatusInternalServerError,
    },status.Error(codes.Internal,fmt.Sprintf("%v",err))

  }

  return &pb.AdminLoginResponse{
    AdminResponseData: &pb.AdminDetailsResponse{
      ID:    int64(loginresponseData.Admin.ID),
      Name:  loginresponseData.Admin.Name,
      Email: loginresponseData.Admin.Email,
    },
    AccessToken: loginresponseData.AccessToken,
    Httpstatus:  http.StatusOK,
    Message : "successsfully logined",
  }, nil
}

func (u *UserServer)UserLoginHandler(ctx context.Context ,req *pb.UserLoginRequest)(*pb.UserLoginResponse,error){

  resData,err:=u.userUsecase.UserLoginHandler(models.UserLogin{Email: req.Email, Password:  req.Password})
  if err!=nil{
    return &pb.UserLoginResponse{
      Httpstatus: http.StatusInternalServerError,
    },status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }
  return &pb.UserLoginResponse{
    Id: int64(resData.Users.Id) ,
    Name: resData.Users.Name,
    Email: resData.Users.Email,
    Phone: resData.Users.Phone,
    ReferralID: resData.Users.ReferralID,
    Httpstatus: int64(codes.OK),
    AccessToken: resData.AccessToken,
    RefreshToken: resData.RefreshToken,
    Message: "successfully logined",
  },nil

}




func (u *UserServer)GetUserDetails(ctx context.Context ,req *pb.GetUserDetailsRequest)(*pb.GetUserDetailsResponse,error){

  resData,err:=u.userUsecase.GetUserDetails(int(req.UserId))
  if err!=nil{
    return &pb.GetUserDetailsResponse{
      Httpstatus: http.StatusInternalServerError,
      Message: "error getting UserDetails",
    },status.Error(codes.Internal,fmt.Sprintf("%v",err))

  }

  return &pb.GetUserDetailsResponse{
    ResData: &pb.GetUserDetailsModel{
      Id: int64(resData.Id),
      Name: resData.Name,
      Email: resData.Email,
      Phone: resData.Phone,
      ReferralID: resData.ReferralID,       

    },
    Httpstatus: http.StatusOK,
    Message: fmt.Sprintf("successfully fetched UserDetails of userID : %d",req.UserId),
  },nil
}



func (u *UserServer)UserSignup(ctx context.Context ,req *pb.UserSignupRequest)(*pb.UserSignupResponse,error){

  resData,err:=u.userUsecase.UserSignup(models.UserDetails{
    Name: req.UserModel.Name,
    Email: req.UserModel.Email,
    Phone: req.UserModel.Phone,
    Password: req.UserModel.Password,
    ConfirmPassword: req.UserModel.ConfirmPassword,

  },req.ReferallCode)

  if err!=nil{
    return &pb.UserSignupResponse{
    },
    status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }

  return &pb.UserSignupResponse{
    ResponseData: &pb.UserSignupResponseModel{
      UserID: int64(resData.Users.Id),
      Name:resData.Users.Name,
      Email:resData.Users.Email ,
      Phone: resData.Users.Phone,
      ReferralID:  resData.Users.ReferralID,

    },
    Httpstatus: http.StatusOK,
    Message: "successfully signed in",
    AccessToken: resData.AccessToken,
    RefreshToken: resData.RefreshToken,
  },nil
}

func (u *UserServer)EditUserDetails(ctx context.Context, req *pb.EditUserDetailsRequest)( *pb.EditUserDetailsResponse, error  ) {


  err:=u.userUsecase.EditUserDetails(int(req.UserId),models.EditUserDetails{
    Name: req.Details.Name,
    Phone: req.Details.Phone,
    Email: req.Details.Email,
    Password: req.Details.Password,
  })
  if err!=nil{
    return nil ,status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }

  return nil,nil

}




func (u *UserServer)GetUsers(ctx context.Context, req *pb.GetUsersRequest)( *pb.GetUsersResponse, error  ) {

  resData,err:=u.adminUsecase.GetUsers()
  if err!=nil{
    return nil ,status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }
  
  
   var resUsers []*pb.UserDetailsAtModel 
  
  for  _,user:=range resData{
  
    resUsers = append(resUsers, &pb.UserDetailsAtModel{
        Id:int64(user.Id) ,
        Name: user.Name,
        Email: user.Email,
        Phone: user.Phone,
        Blocked: user.Blocked,
      })
      
  }
 

  return &pb.GetUsersResponse{
      Details:resUsers,
  }  ,nil

}

func (u *UserServer)AddAddress(ctx context.Context, req *pb.AddAddressReq)( *pb.AddAddressRes, error  ) {

 err:=u.userUsecase.AddAddress(int(req.UserID),models.AddAddress{
        Name: req.Details.Name,
    Address: req.Details.Address,
    LandMark: req.Details.LandMark,
    City: req.Details.City ,
    Pincode:req.Details.Pincode ,
    State: req.Details.State,
    Phone: req.Details.Phone,
  })
  if err!=nil{
    return nil ,status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }
  
  return nil,nil

}

func (u *UserServer)BlockUser(ctx context.Context, req *pb.BlockRequest)( *pb.BlockResponse, error  ) {

 err:=u.adminUsecase.BlockUser(int(req.Id))
  if err!=nil{
    return nil ,status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }
  
  return nil,nil

}
func (u *UserServer)UnBlockUser(ctx context.Context, req *pb.UnBlockRequest)( *pb.UnBlockResponse, error  ) {

 err:=u.adminUsecase.UnBlockUser(int(req.Id))
  if err!=nil{
    return nil ,status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }
  return nil,nil

}


func (u *UserServer)ChangeAdminPassword(ctx context.Context, req *pb.ChangeAdminPasswordRequest)( *pb.ChangeAdminPasswordResponse, error  ) {

 err:=u.adminUsecase.ChangePassword(models.AdminPasswordChange{
      CurrentPassword: req.Details.CurrentPassword,
    NewPassword: req.Details.NewPassword,
    ConfirmPassword: req.Details.ConfirmPassword,
  }, int(req.Id))

  if err!=nil{
    return nil ,status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }
  return nil,nil

}


func (u *UserServer)DeleteAddress(ctx context.Context, req *pb.DeleteAddressReq)( *pb.DeleteAddressRes, error  ) {

 err:=u.userUsecase.DeleteAddress(int(req.AddressId),int(req.UserId))
  if err!=nil{
    return nil ,status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }
  return nil,nil

}

func (u *UserServer)EditAddress(ctx context.Context, req *pb.EditAddressReq)( *pb.EditAddressRes, error  ) {

  err:=u.userUsecase.EditAddress(int(req.ID),uint(req.Userid),models.EditAddress{
    Name:req.EditAddressDetails.Name ,
     Address: req.EditAddressDetails.Address,
    LandMark:req.EditAddressDetails.LandMark ,
    City: req.EditAddressDetails.City,
    Pincode: req.EditAddressDetails.Pincode ,
    State: req.EditAddressDetails.State,
    Phone: req.EditAddressDetails.Phone,
    Default: req.EditAddressDetails.Default,
  })
  if err!=nil{
    return nil ,status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }
  return nil,nil


}

func (u *UserServer)GetAddressess(ctx context.Context, req *pb.GetAddressReq)( *pb.GetAddressRes, error  ) {


  addressess,err:=u.userUsecase.GetAddressess(int(req.UserID))
  if err!=nil{
    return nil ,status.Error(codes.Internal,fmt.Sprintf("%v",err))
  }

  var resData  []*pb.Address 
  
  for  _,ad:=range addressess{
  
    resData = append(resData, &pb.Address{
  
      UserID: int64(ad.UserID),
      Name: ad.Name,
      Address: ad.Address,
      LandMark: ad.LandMark,
      City: ad.City,
      Pincode:ad.Pincode ,
      State: ad.State, 
      Phone: ad.Phone,
      Default: ad.Default,
      })
  }

  return &pb.GetAddressRes{
    
      Details:resData,
  }  ,nil
}

func (u *UserServer)GetWallet(ctx context.Context ,req *pb.GetWalletReq)(*pb.GetWalletRes,error){

  resData,err:=u.userUsecase.GetWallet(int(req.ID))
  if err!=nil{
    return &pb.GetWalletRes{
    },status.Error(codes.Internal,fmt.Sprintf("error getting UserDetails %v",err))

  }

  return &pb.GetWalletRes{
          UserID: int64(resData.UserID), 
    Username: resData.Username,
    TotalAmount: resData.TotalAmount,
  },nil
}

