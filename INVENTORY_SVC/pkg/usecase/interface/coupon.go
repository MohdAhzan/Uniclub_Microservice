package interfaces

import (
	"github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/utils/domain"
	"github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/utils/models"
)

type CouponUseCase interface {
	CreateNewCoupon(coupon models.Coupons) error
	GetAllCoupons() ([]domain.Coupons, error)
	MakeCouponInvalid(couponID int) error
	MakeCouponValid(couponID int) error
}
