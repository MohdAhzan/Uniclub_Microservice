package routes

// func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler, categoryHandler *handler.CategoryHandler,
// 	inventoryHandler *handler.InventaryHandler, OfferHandler *handler.OfferHandler, couponHandler *handler.CouponHandler) {
//
// 	engine.POST("/adminlogin", adminHandler.LoginHandler)
//
// 	engine.Use(middleware.AdminAuthMiddleware)
//
// 	{
//
// 		userManagement := engine.Group("/users")
// 		{
// 			userManagement.GET("", adminHandler.GetUsers)
// 			userManagement.PUT("/block", adminHandler.BlockUser)
// 			userManagement.PUT("/unblock", adminHandler.UnBlockUser)
// 		}
//
// 		categorymanagement := engine.Group("/category")
// 		{
// 			categorymanagement.GET("", categoryHandler.GetCategory)
// 			categorymanagement.POST("", categoryHandler.AddCategory)
// 			categorymanagement.PUT("", categoryHandler.UpdateCategory)
// 			categorymanagement.DELETE("", categoryHandler.DeleteCategory)
//
// 		}
//
// 		productmanagement := engine.Group("/products")
// 		{
// 			productmanagement.POST("", inventoryHandler.AddInventory)
// 			productmanagement.GET("", inventoryHandler.GetProductsForAdmin)
// 			productmanagement.DELETE("", inventoryHandler.DeleteInventory)
// 			productmanagement.PUT("/:id/edit_details", inventoryHandler.EditInventoryDetails)
// 		}
//
// 		couponManagement := engine.Group("/coupons")
// 		{
// 			couponManagement.POST("", couponHandler.CreateNewCoupon)
// 			couponManagement.GET("", couponHandler.GetAllCoupons)
// 			couponManagement.DELETE("", couponHandler.MakeCouponInvalid)
// 			couponManagement.PUT("", couponHandler.MakeCouponValid)
// 		}
//
// 		offerManagment := engine.Group("/offers")
// 		{
// 			offerManagment.POST("/category", OfferHandler.AddCategoryOffer)
// 			offerManagment.GET("/category", OfferHandler.GetAllCategoryOffers)
// 			offerManagment.PUT("/category", OfferHandler.EditCategoryOffer)
// 			offerManagment.DELETE("/category", OfferHandler.ValidorInvalidCategoryOffers)
//
// 			offerManagment.POST("/product", OfferHandler.AddInventoryOffer)
// 			offerManagment.GET("/product", OfferHandler.GetInventoryOffers)
// 			offerManagment.PUT("/product", OfferHandler.EditInventoryOffer)
// 			offerManagment.DELETE("/product", OfferHandler.ValidorInvalidInventoryOffers)
// 		}
//
// 		payment := engine.Group("/payment-methods")
// 		{
// 			payment.GET("", adminHandler.GetPaymentMethods)
// 			payment.POST("", adminHandler.NewPaymentHandler)
// 			payment.DELETE("", adminHandler.DeletePaymentMethod)
//
// 		}
//
// 		ordermanagement := engine.Group("/orders")
// 		{
// 			ordermanagement.GET("", adminHandler.GetAllOrderDetails)
// 			ordermanagement.PUT("/payment-status", adminHandler.MakePaymentStatusAsPaid)
// 			ordermanagement.PUT("/status", adminHandler.EditOrderStatus)
// 			ordermanagement.PUT("/return", adminHandler.OrderReturnApprove)
// 		}
//
// 		engine.PUT("/change_password", adminHandler.ChangePassword)
// 		engine.GET("/printsales", adminHandler.PrintSalesByDate)
// 		engine.GET("/currentsalesreport", adminHandler.FilteredSalesReport)
//
// 	}
// }
