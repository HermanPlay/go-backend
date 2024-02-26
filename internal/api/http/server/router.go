package server

import (
	"github.com/HermanPlay/backend/internal/api/http"
	"github.com/HermanPlay/backend/internal/api/http/middleware"
	"github.com/gin-gonic/gin"
)

func Init(init *http.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		dev := api.Group("/dev")
		dev.GET("/status", init.DevRoute.HealthCheck)

		auth := api.Group("/auth")
		auth.POST("/register", init.AuthRoute.RegisterUser)
		auth.POST("/login", init.AuthRoute.LoginUser)

		user := api.Group("/user")
		user.Use(middleware.JwtAuthMiddleware(init.Cfg))
		user.GET("", init.UserRoute.GetAllUserData)
		user.POST("", init.UserRoute.AddUserData)
		user.GET("/:userID", init.UserRoute.GetUserById)
		user.PATCH("/:userID", init.UserRoute.UpdateUserData)
		user.DELETE("/:userID", init.UserRoute.DeleteUser)
	}

	return router
}
