package services

import "github.com/ExcitingFrog/xuyu/internal/repository"

type IServices interface {
	IHello
}

type Services struct {
	repository repository.IRepository
}

func NewServices(repository repository.IRepository) *Services {
	return &Services{
		repository: repository,
	}
}
