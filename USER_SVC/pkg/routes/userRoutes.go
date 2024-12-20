package routes

//
// import (
// 	"github.com/MohdAhzan/Uniclub_ecommerce_Microservice_project/pkg/api/handler"
// 	"github.com/MohdAhzan/Uniclub_ecommerce_Microservice_project/pkg/api/middleware"
// 	"github.com/gin-gonic/gin"
// )
//
// func UserRoutes(engine *gin.RouterGroup,
// 	userHandler *handler.UserHandler,
// 	otpHandler *handler.OtpHandler, inventoryHandler *handler.InventaryHandler,
// 	cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler,
// 	paymentHandler *handler.PaymentHandler, wishlistHandler *handler.WishlistHandler, couponHandler *handler.CouponHandler) {
//
// 	engine.POST("/signup", userHandler.UserSignUp)
// 	engine.POST("/login", userHandler.UserLoginHandler)
// 	engine.POST("/otplogin", otpHandler.SendOTPHandler)
// 	engine.POST("/verifyotp", otpHandler.VerifyOTPHandler)
//
// 	payment := engine.Group("/payment")
// 	{
// 		payment.GET("/razorpay", paymentHandler.MakePaymentFromRazorPay)
// 		payment.GET("/update-status", paymentHandler.VerifyPaymentFromRazorPay)
// 		payment.GET("/wallet", paymentHandler.PaymentFromWallet)
// 	}
//
// 	engine.Use(middleware.UserAuthMiddleware)
// 	{
// 		home := engine.Group("/home")
// 		{
// 			home.GET("", inventoryHandler.GetProductsForUsers)
// 			home.POST("/add_to_cart", cartHandler.AddtoCart)
// 			home.GET("/coupons", couponHandler.GetAllCoupons)
// 		}
//
// 		search := engine.Group("/search")
// 		{
// 			search.GET("", inventoryHandler.SearchProducts)
// 		}
//
// 		wishlist := engine.Group("/wishlist")
// 		{
// 			wishlist.POST("", wishlistHandler.AddToWishlist)
// 			wishlist.GET("", wishlistHandler.GetWishlist)
// 			wishlist.DELETE("", wishlistHandler.RemoveFromWishlist)
// 		}
// 		cart := engine.Group("/cart")
// 		{
// 			cart.GET("", cartHandler.GetCart)
// 			cart.PUT("", cartHandler.DecreaseCartQuantity)
// 			cart.DELETE("/remove", cartHandler.RemoveCart)
// 		}
//
// 		profile := engine.Group("/profile")
//
// 		{
// 			profile.GET("/details", userHandler.GetUserDetails)
// 			profile.GET("/address", userHandler.GetAddressess)
// 			profile.POST("/address", userHandler.AddAddressess)
// 			profile.DELETE("/address", userHandler.DeleteAddress)
//
// 			wallet := engine.Group("/wallet")
// 			{
// 				wallet.GET("", userHandler.GetWallet)
// 			}
//
// 			edit := engine.Group("/edit")
// 			{
// 				edit.PUT("/account", userHandler.EditUserDetails)
// 				edit.PUT("/address", userHandler.EditAddress)
//
// 				edit.PUT("/password", userHandler.ChangePassword)
// 			}
// 		}
//
// 	}
// }
