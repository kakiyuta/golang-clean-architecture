package usecase

import (
	"github.com/kakiyuta/golang-clean-architecture/app/domain/model"
	"github.com/kakiyuta/golang-clean-architecture/app/usecase/input"
	"github.com/kakiyuta/golang-clean-architecture/app/usecase/output"
)

type Products struct{}

func NewProducts() *Products {
	return &Products{}
}

func (p *Products) GetProducts(input input.ProductsGetProducts) (*output.ProductsGetProducts, error) {
	output := &output.ProductsGetProducts{
		Total: 0,
		Products: []model.Product{
			{
				ID:   1,
				Name: "product1",
				Validations: []model.Validation{
					{
						ID:    1,
						Name:  "validation1",
						Price: 100,
					},
					{
						ID:    2,
						Name:  "validation2",
						Price: 100,
					},
				},
			},
			{
				ID:   2,
				Name: "product2",
				Validations: []model.Validation{
					{
						ID:    3,
						Name:  "validation3",
						Price: 100,
					},
				},
			},
		},
	}
	return output, nil
}
