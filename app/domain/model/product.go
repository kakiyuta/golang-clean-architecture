package model

import "time"

type Product struct {
	ID          int
	Name        string
	Validations []Variant
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}
