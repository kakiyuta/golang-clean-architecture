package input

const (
	GetProductsLimitDefault  = 10
	GetProductsOffsetDefault = 0
)

type ProductsGetProducts struct {
	Limit  int
	Offset int
}

// Constructor for GetProducts
func NewGetProducts(limit *int, offset *int) *ProductsGetProducts {
	gp := &ProductsGetProducts{
		Limit:  GetProductsLimitDefault,
		Offset: GetProductsOffsetDefault,
	}
	if limit != nil {
		gp.Limit = *limit
	}
	if offset != nil {
		gp.Offset = *offset
	}
	return gp
}

// ProductsCreateProduct is the input struct for CreateProduct
type ProductsCreateProduct struct {
	Name string
}

// Constructor for CreateProduct
func NewCreateProduct(name string) *ProductsCreateProduct {
	return &ProductsCreateProduct{
		Name: name,
	}
}
