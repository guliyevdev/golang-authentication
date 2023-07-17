package app

import (
	"auth/app/handlers"
	"auth/config"
	"auth/database"
	"auth/domain/repositories"
	"auth/middleware"
	"auth/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	config.LoadEnvVariables()
	database.ConnectToDb()
	database.SyncDatabase()
}

var (
	router = gin.Default()
)

func Start() {
	authRepository := repositories.NewAuthRepository(database.DB)
	ah := handlers.NewAuthHandler(service.NewLoginService(authRepository))

	router.POST("/signup", ah.SignUp)
	router.POST("/login", ah.Login)
	router.Any("/test/*path", func(c *gin.Context) {
		targetURL := c.Request.URL.Path
		c.JSON(http.StatusOK, gin.H{
			"message": "hello I am listening everything",
			"path":    targetURL,
		})
	})
	router.GET("/validate", middleware.RequireAuth, handlers.Validate)
	router.Run() // listen and serve
}
