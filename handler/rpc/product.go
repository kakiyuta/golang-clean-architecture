package rpc

import (
	"context"

	"github.com/kakiyuta/golang-clean-architecture/app/domain/dto/input"
	"github.com/kakiyuta/golang-clean-architecture/app/domain/dto/output"
	"github.com/kakiyuta/golang-clean-architecture/app/gen/grpc"
	"github.com/kakiyuta/golang-clean-architecture/app/registry"
	"github.com/kakiyuta/golang-clean-architecture/app/usecase"
)

type ProductServer struct {
	grpc.UnimplementedGreetingServiceServer
	repo registry.RepositoryInterface
}

func NewProductServer(repo registry.RepositoryInterface) *ProductServer {
	return &ProductServer{
		repo: repo,
	}
}

func (s *ProductServer) Products(ctx context.Context, req *grpc.ProductsRequest) (*grpc.ProductsResponse, error) {
	// TODO 本来はここでバリデーションやエラーハンドリングを行う
	offset := int(req.Offset)
	limit := int(req.Limit)
	input := input.NewGetProducts(&limit, &offset)
	usecase := s.newProductsUseCase()
	output, err := usecase.GetProducts(*input)
	if err != nil {
		return nil, err
	}

	res := &grpc.ProductsResponse{
		Total:    int32(output.Total),
		Products: convertProducts(output),
	}

	return res, nil
}

// newProductsUseCase Productsユースケースを作成
func (s *ProductServer) newProductsUseCase() usecase.ProductsUsecase {
	return usecase.NewProductsUsecase(
		s.repo.GetDB(),
		s.repo.NewProducts(),
		s.repo.NewVariants(),
	)
}

// convertProducts converts output.ProductsGetProducts to []*api.Product
func convertProducts(output *output.ProductsGetProducts) []*grpc.Product {
	products := make([]*grpc.Product, len(output.Products))

	for i, product := range output.Products {
		variants := make([]*grpc.Variant, len(product.Variants))
		for j, variant := range product.Variants {
			variants[j] = &grpc.Variant{
				Id:    int32(variant.ID),
				Name:  variant.Name,
				Price: int32(variant.Price),
			}
		}

		products[i] = &grpc.Product{
			Id:       int32(product.ID),
			Name:     product.Name,
			Variants: variants,
		}
	}

	return products
}
