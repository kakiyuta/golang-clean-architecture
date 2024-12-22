package controller

import (
	"net/http"

	"github.com/kakiyuta/golang-clean-architecture/app/gen/api"
	"github.com/kakiyuta/golang-clean-architecture/app/infra/db"
	"github.com/kakiyuta/golang-clean-architecture/app/usecase"
	"github.com/kakiyuta/golang-clean-architecture/app/usecase/input"
	"github.com/kakiyuta/golang-clean-architecture/app/usecase/output"
	"github.com/labstack/echo/v4"
)

type ProductList struct {
	Products []*api.ProductVariants `json:"products"`
	Total    int                    `json:"total"`
}

func (c *Controller) GetV1Products(ctx echo.Context, params api.GetV1ProductsParams) error {
	input := input.NewGetProducts(params.Limit, params.Offset)
	usecase := c.newProductsUseCase()
	output, err := usecase.GetProducts(*input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	products := convertProducts(output)

	result := ProductList{
		Products: products,
		Total:    output.Total,
	}

	return ctx.JSON(http.StatusOK, result)
}

func (c *Controller) PostV1Products(ctx echo.Context) error {

	var product api.PostV1ProductsJSONRequestBody
	if err := ctx.Bind(&product); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	input := input.NewCreateProduct(product.Name)
	usecase := c.newProductsUseCase()
	output, err := usecase.CreateProduct(*input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	result := api.Prouct{
		Id:   Int64Ptr(int64(output.Product.ID)),
		Name: StringPtr(output.Product.Name),
	}

	return ctx.JSON(http.StatusOK, result)
}

// newProductsUseCase Productsユースケースを作成
func (c *Controller) newProductsUseCase() usecase.ProductsUsecase {
	conn, err := db.NewMySQLConnector()
	if err != nil {
		panic(err)
	}

	return usecase.NewProductsUsecase(
		c.repo.NewProducts(),
		c.repo.NewVariants(),
		conn,
	)
}

// convertProducts converts output.ProductsGetProducts to []*api.Product
func convertProducts(output *output.ProductsGetProducts) []*api.ProductVariants {
	products := make([]*api.ProductVariants, len(output.Products))

	for i, product := range output.Products {
		validations := make([]api.Variant, len(product.Variants))

		for j, validation := range product.Variants {
			validations[j] = api.Variant{
				Id:    Int64Ptr(int64(validation.ID)),
				Name:  StringPtr(validation.Name),
				Price: IntPtr(validation.Price),
			}
		}

		products[i] = &api.ProductVariants{
			Id:          Int64Ptr(int64(product.ID)),
			Name:        StringPtr(product.Name),
			Validations: &validations,
		}
	}

	return products
}
