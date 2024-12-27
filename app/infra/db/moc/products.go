package moc

import "github.com/kakiyuta/golang-clean-architecture/app/domain/model"

type Product struct{}

func (p *Product) GetProducts(limit int, offset int) ([]model.Product, error) {
	return []model.Product{
		{
			ID:   1,
			Name: "product1",
		},
		{
			ID:   99,
			Name: "product2",
		},
	}, nil
}

func (p *Product) GetProductsWithVariation(limit int, offset int) ([]model.Product, error) {
	return []model.Product{
		{
			ID:   1,
			Name: "product1",
			Variants: []model.Variant{
				{
					ID:    1,
					Name:  "validation1",
					Price: 100,
				},
				{
					ID:    99,
					Name:  "validation2",
					Price: 100,
				},
			},
		},
		{
			ID:   2,
			Name: "product2",
			Variants: []model.Variant{
				{
					ID:    3,
					Name:  "validation3",
					Price: 100,
				},
			},
		},
	}, nil
}

func (p *Product) GetProductByID(id int) (model.Product, error) {
	return model.Product{
		ID:   1,
		Name: "product1",
	}, nil
}

func (p *Product) CreateProduct(produnct model.Product) (model.Product, error) {
	return model.Product{
		ID:   1,
		Name: produnct.Name,
	}, nil
}
