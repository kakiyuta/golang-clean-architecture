package mysql

import (
	"github.com/kakiyuta/golang-clean-architecture/app/domain/model"
	"github.com/kakiyuta/golang-clean-architecture/app/infra/db"
)

type Variants struct {
	Con *db.MySQLConnector
}

func (v *Variants) GetVariants(productID int) ([]model.Variant, error) {
	return []model.Variant{
		{
			ID:        1,
			ProductID: 1,
			Name:      "variant1",
			Price:     1000,
		},
		{
			ID:        2,
			ProductID: 1,
			Name:      "variant2",
			Price:     2000,
		},
	}, nil
}
