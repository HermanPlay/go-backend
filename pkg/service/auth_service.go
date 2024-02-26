package service

import (
	"net/http"

	"github.com/HermanPlay/backend/internal/api/http/constant"
	"github.com/HermanPlay/backend/internal/api/http/util"
	"github.com/HermanPlay/backend/pkg/domain/dao"
	"github.com/HermanPlay/backend/pkg/domain/dto"
	"github.com/HermanPlay/backend/pkg/repository"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type AuthServiceImpl struct {
	authRepository repository.AuthRepository
	userRepository repository.UserRepository
}

func (a AuthServiceImpl) RegisterUser(c *gin.Context) {
	var userInput dto.UserRegister
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	if a.userRepository.CheckUserExist(userInput.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email already exist"})
		return
	}
	user := dao.User{}
	user.Email = userInput.Email
	user.Password = userInput.Password
	data, error := a.userRepository.Save(&user)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error when saving data"})
		return
	}
	c.JSON(http.StatusOK, util.BuildResponse(constant.Success, data))
}

func (a AuthServiceImpl) LoginUser(c *gin.Context) {
	var userInput dto.UserLogin
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	if !a.userRepository.CheckUserExist(userInput.Email) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Email doesn't exist"})
		return
	}

	token, err := a.authRepository.LoginUser(userInput)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func NewAuthService(authRepository repository.AuthRepository, userRepository repository.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		authRepository: authRepository,
		userRepository: userRepository,
	}
}
