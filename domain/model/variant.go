package model

import "time"

type Variant struct {
	ID        int
	ProductID int
	Name      string
	Price     int
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
