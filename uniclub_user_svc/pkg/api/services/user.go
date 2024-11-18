package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/MohdAhzan/Uniclub_ecommerce_Microservice_project/pkg/pb"
	interfaces "github.com/MohdAhzan/Uniclub_ecommerce_Microservice_project/pkg/usecase/interface"
	"github.com/MohdAhzan/Uniclub_ecommerce_Microservice_project/pkg/utils/models"
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
		}, err
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

func (u *UserServer)LoginHandler(ctx context.Context ,req *pb.UserLoginRequest)(*pb.UserLoginResponse,error){
  
  resData,err:=u.userUsecase.UserLoginHandler(models.UserLogin{Email: req.Email, Password:  req.Password})
  if err!=nil{
      return &pb.UserLoginResponse{
      Httpstatus: http.StatusInternalServerError,
    },errors.New("Error User Login ")
  }
 
     return &pb.UserLoginResponse{
      Id: int64(resData.Users.Id) ,
      Name: resData.Users.Name,
      Email: resData.Users.Email,
      Phone: resData.Users.Phone,
      ReferralID: resData.Users.ReferralID,
      Httpstatus: http.StatusOK,
      AccessToken: resData.AccessToken,
      RefreshToken: resData.RefreshToken,
      Message: "successfully logined ",
    },nil
}




func (u *UserServer)GetUserDetails(ctx context.Context ,req *pb.GetUserDetailsRequest)(*pb.GetUserDetailsResponse,error){
 
  resData,err:=u.userUsecase.GetUserDetails(int(req.UserId))
  if err!=nil{
      return &pb.GetUserDetailsResponse{
      Httpstatus: http.StatusInternalServerError,
      Message: "error getting UserDetails",
    },err
  }
 
     return &pb.GetUserDetailsResponse{
    ResData: &pb.GetUserDetailsModel{
      Id: int64(resData.Id) ,
      Name: resData.Name,
      Email: resData.Email,
      Phone: resData.Phone,

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
    },err
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
    },nil
}



