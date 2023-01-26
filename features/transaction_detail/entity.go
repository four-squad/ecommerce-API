package transaction_detail

import "time"

type CoreTransaction_Detail struct {
	ID           uint
	ProductID    uint
	NameBuyer    string
	NameSeller   string
	ProductName  string
	ProductPrice uint
	Qty          uint
	TotalPrice   uint
	Image        string
	UpdatedAt    time.Time
}
type Transaction_DetailHandler interface {
	// Register() echo.HandlerFunc
	// Login() echo.HandlerFunc
	// Profile() echo.HandlerFunc
	// Update() echo.HandlerFunc
	// Deactivate() echo.HandlerFunc
}

type Transaction_DetailService interface {
	// Register(newUser Core) error
	// Login(email, password string) (string, Core, error)
	// Profile(token interface{}) (interface{}, error)
	// Update(formHeader multipart.FileHeader, token interface{}, updatedProfile Core) (Core, error)
	// Deactivate(token interface{}) error
}

type Transaction_DetailData interface {
	// Register(newUser Core) error
	// Login(email string) (Core, error)
	// Profile(id uint) (interface{}, error)
	// Update(id uint, updatedProfile Core) (Core, error)
	// Deactivate(id uint) error
}
