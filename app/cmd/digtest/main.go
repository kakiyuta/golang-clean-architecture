package main

// 通常とdigとfxの比較

import (
	"fmt"

	"go.uber.org/dig"
)

type Usecase interface {
	Use()
}

type Repository interface {
	RepoPrint()
}

type usecase struct {
	repo Repository
}

func NewUsecase(r Repository) Usecase {
	return &usecase{
		repo: r,
	}
}
func (u usecase) Use() {
	u.repo.RepoPrint()
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r repository) RepoPrint() {
	fmt.Println("1")
}

func main() {

	// 通常の依存関係注入
	repo := NewRepository()
	usecase := NewUsecase(repo)
	usecase.Use()

	// uber-go/digを使った依存関係注入
	c := dig.New()
	c.Provide(NewUsecase)
	c.Provide(NewRepository)
	c.Invoke(func(u Usecase) {
		u.Use()
	})
}
