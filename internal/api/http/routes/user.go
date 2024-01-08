package routes

import (
	"github.com/HermanPlay/backend/pkg/service"
	"github.com/gin-gonic/gin"
)

type UserRoute interface {
	GetAllUserData(c *gin.Context)
	AddUserData(c *gin.Context)
	GetUserById(c *gin.Context)
	UpdateUserData(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserRouteImpl struct {
	service service.UserService
}

func (u UserRouteImpl) GetAllUserData(c *gin.Context) {
	u.service.GetAllUser(c)
}

func (u UserRouteImpl) AddUserData(c *gin.Context) {
	u.service.AddUserData(c)
}

func (u UserRouteImpl) GetUserById(c *gin.Context) {
	u.service.GetUserById(c)
}

func (u UserRouteImpl) UpdateUserData(c *gin.Context) {
	u.service.UpdateUserData(c)
}

func (u UserRouteImpl) DeleteUser(c *gin.Context) {
	u.service.DeleteUser(c)
}

func NewUserRoute(userService service.UserService) *UserRouteImpl {
	return &UserRouteImpl{
		service: userService,
	}
}
