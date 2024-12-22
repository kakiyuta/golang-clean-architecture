package output

import "github.com/kakiyuta/golang-clean-architecture/app/domain/model"

type ProductsGetProducts struct {
	Total    int
	Products []model.Product
}

type ProdunctsGreateProdunct struct {
	Product model.Product
}
