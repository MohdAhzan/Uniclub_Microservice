package domain

type Users struct {
	ID         uint   `json:"id" gorm:"unique;not null"`
	Name       string `json:"name"`
	Email      string `json:"email" validate:"email"`
	Password   string `json:"password" validate:"min=8,max=20"`
	Phone      string `json:"phone"`
	Blocked    bool   `json:"blocked" gorm:"default:false"`
	ReferralID string `json:"referral_id" gorm:"unique"`
}

type Address struct {
	ID       uint   `json:"id" gorm:"unique;not null"`
	UserID   uint   `json:"user_id"`
	Users    Users  `json:"-"  gorm:"foreignkey:UserID"`
	Name     string `json:"name" validate:"required"`
	Address  string `json:"address" validate:"required"`
	LandMark string `json:"landmark"`
	City     string `json:"city" validate:"required"`
	Pincode  string `json:"pincode" validate:"required,len=6"`
	State    string `json:"state" validate:"required"`
	Phone    string `json:"phone" gorm:"phone"`
	Default  bool   `json:"default" gorm:"default:false"`
}

type Wallet struct {
	WalletID        uint    `json:"wallet_id" gorm:"primarykey;autoIncrement"`
	UserID          uint    `json:"user_id"`
	Users           Users   `json:"-" gorm:"foreignkey : UserID"`
	WalletAmount    float64 `json:"wallet_amount" gorm:"default:0"`
	TransactionType string  `json:"transaction_type" gorm:"transaction_type:4;check:transaction_type IN ('REFERAL','PDT_CANCELLED','PDT_RETURNED','DEBIT')"`
}

