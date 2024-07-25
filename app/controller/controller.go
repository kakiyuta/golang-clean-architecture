package controller

import (
	api "github.com/kakiyuta/golang-clean-architecture/app/gen"
	"github.com/kakiyuta/golang-clean-architecture/app/registry"
)

type Controller struct {
	repo registry.RepositoryInterface
}

func NewController(repo registry.RepositoryInterface) api.ServerInterface {
	return &Controller{
		repo: repo,
	}
}

func Int64Ptr(i int64) *int64 {
	return &i
}

func StringPtr(s string) *string {
	return &s
}

func IntPtr(i int) *int {
	return &i
}
