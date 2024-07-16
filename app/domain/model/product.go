package model

import "time"

type Product struct {
	ID          int
	Name        string
	Validations []Validation
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}
