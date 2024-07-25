package controller

import (
	registry "github.com/kakiyuta/golang-clean-architecture/app/registory"
)

type Controller struct {
	repo registry.RepositoryInterface
}

func NewController(repo registry.RepositoryInterface) Controller {
	return Controller{
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
