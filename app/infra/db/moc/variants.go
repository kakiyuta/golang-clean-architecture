package moc

import "github.com/kakiyuta/golang-clean-architecture/app/domain/model"

type Variants struct{}

func (v *Variants) GetVariants(productID int) ([]model.Variant, error) {
	return []model.Variant{
		{
			ID:        1,
			ProductID: 1,
			Name:      "variant1",
			Price:     1000,
		},
		{
			ID:        2,
			ProductID: 1,
			Name:      "variant2",
			Price:     2000,
		},
	}, nil
}
