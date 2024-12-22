package mysql

import (
	"github.com/kakiyuta/golang-clean-architecture/app/domain/model"
	"gorm.io/gorm"
)

type Product struct {
	Con *gorm.DB
}

func (p *Product) GetProducts(limit int, offset int) ([]model.Product, error) {
	var products []model.Product
	p.Con.Limit(limit).Offset(offset).Find(&products).Where("deleted_at IS NULL")
	return products, nil
}

func (p *Product) GetProductsWithVariation(limit int, offset int) ([]model.Product, error) {
	var products []model.Product
	p.Con.Limit(limit).Offset(offset).Preload("Variants").Find(&products).Where("deleted_at IS NULL")
	return products, nil
}

func (p *Product) GetProductByID(id int) (model.Product, error) {
	var product model.Product
	p.Con.First(&product, id)
	return product, nil
}
