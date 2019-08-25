package structs

import "time"

type Product struct {
	ID           int
	ProductTitle string
	Price        int
	Quantity     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

func (Product) TableName() string {
	return "product"
}
