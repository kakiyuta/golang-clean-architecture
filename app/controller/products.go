package controller

import (
	"net/http"

	api "github.com/kakiyuta/golang-clean-architecture/app/gen"
	"github.com/labstack/echo/v4"
)

func (c *Controller) GetV1Products(ctx echo.Context, params api.GetV1ProductsParams) error {
	products := []*api.Product{
		{
			Id:   Int64Ptr(1),
			Name: StringPtr("product1"),
			Validations: &[]api.Validation{
				{
					Id:    Int64Ptr(1),
					Name:  StringPtr("validation1"),
					Price: IntPtr(100),
				},
				{
					Id:    Int64Ptr(2),
					Name:  StringPtr("validation2"),
					Price: IntPtr(100),
				},
			},
		},
		{
			Id:   Int64Ptr(2),
			Name: StringPtr("product2"),
			Validations: &[]api.Validation{
				{
					Id:    Int64Ptr(3),
					Name:  StringPtr("validation3"),
					Price: IntPtr(100),
				},
			},
		},
	}

	return ctx.JSON(http.StatusOK, products)
}
