package controllers

import (
	"github.com/ExcitingFrog/xuyu/internal/services"
)

type Controllers struct {
	services services.IServices
}

func NewControllers(services services.IServices) *Controllers {
	return &Controllers{
		services: services,
	}
}
