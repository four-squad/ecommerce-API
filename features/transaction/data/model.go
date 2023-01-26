package data

import (
	c "ecommerce/features/cart/data"
	p "ecommerce/features/product/data"
	"ecommerce/features/transaction"

	x "ecommerce/features/transaction_detail/data"

	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model
	TotalPrice        uint
	Address           string
	PaymentUrl        string
	Status            string `gorm:"default:pending"`
	UserID            uint
	User              Users
	TransactionDetail []x.Transaction_Detail `gorm:"foreignKey:TransactionID"`
}

type Users struct {
	gorm.Model
	Avatar      string
	Name        string
	Email       string
	Address     string
	Password    string
	Cartss      []c.Carts      `gorm:"foreignKey:UserID"`
	Product     []p.Products   `gorm:"foreignKey:UserID"`
	Transaction []Transactions `gorm:"foreignKey:UserID"`
}

// type TransactionDetail struct {
// 	TransactionID uint
// 	ProductID     uint
// 	Title         string
// 	Price         uint
// 	Qty           uint
// 	TotalPrice    uint
// }

func ToCore(data Transactions) transaction.Core {
	return transaction.Core{
		ID: data.ID,
		// SellerName:   data.User.Name,
		// BuyerName:    data.User.Name,
		Address:      data.Address,
		ProductName:  "",
		ProductPrice: 0,
		Qty:          0,
		TotalPrice:   data.TotalPrice,
		PaymentUrl:   data.PaymentUrl,
		Status:       data.Status,
		UpdatedAt:    data.UpdatedAt,
	}
}

func CoreToData(data transaction.Core) Transactions {
	return Transactions{
		Model:      gorm.Model{ID: data.ID},
		TotalPrice: data.TotalPrice,
		Address:    data.Address,
		PaymentUrl: data.PaymentUrl,
		Status:     data.Status,
		UserID:     data.UserID,
		// User: Users{
		// 	Model: gorm.Model{ID: data.ID},
		// 	Name:  data.BuyerName,
		// },
	}
}
func CoreToData2(data transaction.Core) Transactions {
	return Transactions{
		Model:      gorm.Model{ID: data.ID},
		TotalPrice: data.TotalPrice,
		Address:    data.Address,
		PaymentUrl: data.PaymentUrl,
		Status:     data.Status,
		UserID:     data.UserID,
		// User: Users{
		// 	Model: gorm.Model{ID: data.ID},
		// 	Name:  data.SellerName,
		// },
	}
}
