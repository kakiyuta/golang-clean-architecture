package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kakiyuta/golang-clean-architecture/app/domain/model"
	"github.com/kakiyuta/golang-clean-architecture/app/domain/repository"
	"github.com/kakiyuta/golang-clean-architecture/app/infra/db"
	"github.com/kakiyuta/golang-clean-architecture/app/usecase/input"
)

func TestProductsUsecase_正常系(t *testing.T) {
	// モックを呼び出すための Controller を生成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックの生成
	mockProducts := repository.NewMockProducts(ctrl)
	mockVariants := repository.NewMockVariants(ctrl)
	mockDb := db.NewMockConnector(ctrl)

	// モックの振る舞いを定義
	mockProducts.EXPECT().GetProductsWithVariation(3, 0).Return(
		[]model.Product{
			{
				ID:   1,
				Name: "product1",
				Variants: []model.Variant{
					{
						ID:    1,
						Name:  "variant1",
						Price: 100,
					},
					{
						ID:    2,
						Name:  "variant2",
						Price: 200,
					},
				},
			},
			{
				ID:       2,
				Name:     "product2",
				Variants: []model.Variant{},
			},
		}, nil,
	)

	// テスト対象のインスタンスを生成
	p := NewProductsUsecase(mockDb, mockProducts, mockVariants)

	// テスト対象のメソッドを呼び出し
	_, err := p.GetProducts(input.ProductsGetProducts{
		Limit:  3,
		Offset: 0,
	})

	// 検証
	if err != nil {
		t.Errorf("error = %v, want nil", err)
	}
}
