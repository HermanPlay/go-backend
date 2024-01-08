package routes

import (
	"github.com/HermanPlay/backend/pkg/service"
	"github.com/gin-gonic/gin"
)

type AuthRoute interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type AuthRouteImpl struct {
	service service.AuthService
}

func (a AuthRouteImpl) RegisterUser(c *gin.Context) {
	a.service.RegisterUser(c)
}

func (a AuthRouteImpl) LoginUser(c *gin.Context) {
	a.service.LoginUser(c)
}

func NewAuthRoute(service service.AuthService) *AuthRouteImpl {
	return &AuthRouteImpl{
		service: service,
	}
}
