package repository

import (
	"log"

	"errors"

	"github.com/HermanPlay/backend/internal/api/http/util/token"
	"github.com/HermanPlay/backend/internal/config"
	"github.com/HermanPlay/backend/pkg/domain/dao"
	"github.com/HermanPlay/backend/pkg/domain/dto"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository interface {
	LoginUser(userInput dto.UserLogin) (string, error)
	CheckUserExist(email string) bool
}

type AuthRepositoryImpl struct {
	db  *gorm.DB
	cfg *config.Config
}

func (a AuthRepositoryImpl) CheckUserExist(email string) bool {
	var user dao.User
	err := a.db.Model(&user).Where("email = ?", email).Limit(1).Find(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("User not found")
			return false
		}
		panic(err)
	}
	return true
}

func verifyPassword(inputPassword, validPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(validPassword), []byte(inputPassword))
}

func (a AuthRepositoryImpl) LoginUser(userInput dto.UserLogin) (string, error) {
	var user dao.User
	err := a.db.Model(&user).Where("email = ?", userInput.Email).Limit(1).Find(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("email not found")
		}
		logrus.Error("Error getting user from db, err: ", err.Error())
		return "", err
	}

	err = verifyPassword(userInput.Password, user.Password)
	if err != nil {
		return "", err
	}

	token, err := token.GenerateToken(uint(user.ID), a.cfg)
	if err != nil {
		return "", err
	}
	return token, nil
}

func NewAuthRepository(db *gorm.DB, cfg *config.Config) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{
		db:  db,
		cfg: cfg,
	}
}
