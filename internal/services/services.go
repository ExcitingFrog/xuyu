package services

import (
	"github.com/ExcitingFrog/xuyu/internal/repository"
	"github.com/ExcitingFrog/xuyu/internal/resources"
)

type IServices interface {
	IHello
}

type Services struct {
	repository repository.IRepository

	xuanwu *resources.Xuanwu
}

func NewServices(repository repository.IRepository, xuanwu *resources.Xuanwu) *Services {
	return &Services{
		repository: repository,
		xuanwu:     xuanwu,
	}
}
