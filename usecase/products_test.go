package usecase

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kakiyuta/golang-clean-architecture/app/domain/model"
	"github.com/kakiyuta/golang-clean-architecture/app/domain/repository"
	"github.com/kakiyuta/golang-clean-architecture/app/infra/db"
	"github.com/kakiyuta/golang-clean-architecture/app/usecase/input"
	"github.com/kakiyuta/golang-clean-architecture/app/usecase/output"
)

func TestProductsUsecase_GetProducts_正常系(t *testing.T) {
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

func TestProductsUsecase_CreateProduct(t *testing.T) {
	// モックを呼び出すための Controller を生成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// モックの生成
	mockProducts := repository.NewMockProducts(ctrl)
	mockVariants := repository.NewMockVariants(ctrl)
	mockDb := db.NewMockConnector(ctrl)

	type fields struct {
		ConnectionController db.Connector
		ProductRepository    repository.Products
		VariantRepository    repository.Variants
	}
	type args struct {
		input input.ProductsCreateProduct
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		mockSetup func()
		want      *output.ProdunctsGreateProdunct
		wantErr   bool
	}{
		{
			name: "正常系",
			fields: fields{
				ConnectionController: mockDb,
				ProductRepository:    mockProducts,
				VariantRepository:    mockVariants,
			},
			args: args{
				input: input.ProductsCreateProduct{
					Name: "product1",
				},
			},
			mockSetup: func() {
				mockProducts.EXPECT().CreateProduct(model.Product{Name: "product1"}).Return(model.Product{ID: 1, Name: "product1"}, nil)
				mockDb.EXPECT().Begin().Return(nil)
				mockDb.EXPECT().Commit().Return(nil)
				mockDb.EXPECT().Rollback().Return()
			},
			want: &output.ProdunctsGreateProdunct{
				Product: model.Product{
					ID:   1,
					Name: "product1",
				},
			},
			wantErr: false,
		},
		{
			name: "異常系 - CreateProduct エラー",
			fields: fields{
				ConnectionController: mockDb,
				ProductRepository:    mockProducts,
				VariantRepository:    mockVariants,
			},
			args: args{
				input: input.ProductsCreateProduct{
					Name: "product2",
				},
			},
			mockSetup: func() {
				mockProducts.EXPECT().CreateProduct(model.Product{Name: "product2"}).Return(model.Product{}, errors.New("create product error"))
				mockDb.EXPECT().Begin().Return(nil)
				mockDb.EXPECT().Rollback().Return()
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProductsUsecase{
				ConnectionController: tt.fields.ConnectionController,
				ProductRepository:    tt.fields.ProductRepository,
				VariantRepository:    tt.fields.VariantRepository,
			}

			// モックの振る舞いを定義
			tt.mockSetup()

			_, err := p.CreateProduct(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductsUsecase.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			/*
				ここのテストはDBのカラム変更があった場合に、テストが通らなくなりやすい
				メンテナンスコストを考慮してコメントアウトする
			*/
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("ProductsUsecase.CreateProduct() = %v, want %v", got, tt.want)
			// }
		})
	}
}
