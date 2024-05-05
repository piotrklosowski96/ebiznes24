package models

import "time"

// CommonFields ...
type CommonFields struct {
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
	SoftDelete bool      `json:"soft_delete" bson:"soft_delete"`
}
