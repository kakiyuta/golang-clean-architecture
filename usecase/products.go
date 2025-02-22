package usecase

import (
	"github.com/kakiyuta/golang-clean-architecture/app/domain/dto/input"
	"github.com/kakiyuta/golang-clean-architecture/app/domain/dto/output"
	"github.com/kakiyuta/golang-clean-architecture/app/domain/model"
	"github.com/kakiyuta/golang-clean-architecture/app/domain/repository"
	"github.com/kakiyuta/golang-clean-architecture/app/infra/db"
)

type ProductsUsecase struct {
	ConnectionController db.Connector
	ProductRepository    repository.Products
	VariantRepository    repository.Variants
}

func NewProductsUsecase(cc db.Connector, p repository.Products, v repository.Variants) ProductsUsecase {
	return ProductsUsecase{
		ConnectionController: cc,
		ProductRepository:    p,
		VariantRepository:    v,
	}
}

// GetProducts 商品一覧を取得
func (p *ProductsUsecase) GetProducts(input input.ProductsGetProducts) (*output.ProductsGetProducts, error) {
	productsVariants, err := p.ProductRepository.GetProductsWithVariation(input.Limit, input.Offset)
	if err != nil {
		return nil, err
	}

	// return nil, weberrors.New(500, "Internal Server Error hoge hoge")

	output := &output.ProductsGetProducts{
		Total:    len(productsVariants),
		Products: productsVariants,
	}

	return output, nil
}

func (p *ProductsUsecase) CreateProduct(input input.ProductsCreateProduct) (*output.ProdunctsGreateProdunct, error) {

	// トランザクションの動作確認
	err := p.ConnectionController.Begin()
	if err != nil {
		return nil, err
	}
	defer p.ConnectionController.Rollback()

	// 商品を登録
	product := model.Product{
		Name: input.Name,
	}
	newProduct, err := p.ProductRepository.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	err = p.ConnectionController.Commit()
	if err != nil {
		return nil, err
	}

	output := &output.ProdunctsGreateProdunct{
		Product: newProduct,
	}
	return output, nil
}
