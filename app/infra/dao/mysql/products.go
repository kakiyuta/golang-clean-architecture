package mysql

import (
	"github.com/kakiyuta/golang-clean-architecture/app/domain/model"
	"github.com/kakiyuta/golang-clean-architecture/app/infra/db"
)

type Product struct {
	// Con *gorm.DB
	Con *db.MySQLConnector
}

func (p *Product) GetProducts(limit int, offset int) ([]model.Product, error) {
	db := p.Con.GetSlave()

	var products []model.Product
	db.Limit(limit).Offset(offset).Find(&products).Where("deleted_at IS NULL")
	return products, nil
}

func (p *Product) GetProductsWithVariation(limit int, offset int) ([]model.Product, error) {
	db := p.Con.GetSlave()

	var products []model.Product
	db.Limit(limit).Offset(offset).Preload("Variants").Find(&products).Where("deleted_at IS NULL")
	return products, nil
}

func (p *Product) GetProductByID(id int) (model.Product, error) {
	db := p.Con.GetSlave()
	var product model.Product
	db.First(&product, id)
	return product, nil
}

func (p *Product) CreateProduct(produnct model.Product) (model.Product, error) {
	db := p.Con.GetMaster()
	db.Create(&produnct)
	return produnct, nil
}
