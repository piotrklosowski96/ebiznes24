package models

// Cart ...
type Cart struct {
	CommonFields

	Name        string
	Description *string
	Products    []*Product `gorm:"many2many:carts_products;"`
}
