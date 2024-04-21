package models

import "github.com/google/uuid"

// CartCreate ...
type CartCreate struct {
	CommonFields

	Name        string
	Description *string
	Products    []uuid.UUID `gorm:"many2many:carts_products;"`
}
