package repository

import "github.com/kakiyuta/golang-clean-architecture/app/domain/model"

type Variants interface {
	GetVariants(productID int) ([]model.Variant, error)
}
