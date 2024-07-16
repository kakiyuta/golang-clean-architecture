package controller

import (
	"net/http"

	api "github.com/kakiyuta/golang-clean-architecture/app/gen"
	"github.com/kakiyuta/golang-clean-architecture/app/usecase"
	"github.com/kakiyuta/golang-clean-architecture/app/usecase/input"
	"github.com/kakiyuta/golang-clean-architecture/app/usecase/output"
	"github.com/labstack/echo/v4"
)

func (c *Controller) GetV1Products(ctx echo.Context, params api.GetV1ProductsParams) error {
	// products := []*api.Product{
	// 	{
	// 		Id:   Int64Ptr(1),
	// 		Name: StringPtr("product1"),
	// 		Validations: &[]api.Validation{
	// 			{
	// 				Id:    Int64Ptr(1),
	// 				Name:  StringPtr("validation1"),
	// 				Price: IntPtr(100),
	// 			},
	// 			{
	// 				Id:    Int64Ptr(2),
	// 				Name:  StringPtr("validation2"),
	// 				Price: IntPtr(100),
	// 			},
	// 		},
	// 	},
	// 	{
	// 		Id:   Int64Ptr(2),
	// 		Name: StringPtr("product2"),
	// 		Validations: &[]api.Validation{
	// 			{
	// 				Id:    Int64Ptr(3),
	// 				Name:  StringPtr("validation3"),
	// 				Price: IntPtr(100),
	// 			},
	// 		},
	// 	},
	// }

	input := input.NewGetProducts(params.Limit, params.Offset)
	output, err := usecase.NewProducts().GetProducts(*input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	products := convertProducts(output)

	return ctx.JSON(http.StatusOK, products)
}

// convertProducts converts output.ProductsGetProducts to []*api.Product
func convertProducts(output *output.ProductsGetProducts) []*api.Product {
	products := make([]*api.Product, len(output.Products))

	for i, product := range output.Products {
		validations := make([]api.Validation, len(product.Validations))

		for j, validation := range product.Validations {
			validations[j] = api.Validation{
				Id:    Int64Ptr(int64(validation.ID)),
				Name:  StringPtr(validation.Name),
				Price: IntPtr(validation.Price),
			}
		}

		products[i] = &api.Product{
			Id:          Int64Ptr(int64(product.ID)),
			Name:        StringPtr(product.Name),
			Validations: &validations,
		}
	}

	return products
}
