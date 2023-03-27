package main

import (
	docs "API/docs"
	"API/internal"
	"API/internal/controllers"
	"API/internal/controllers/auth"
	"API/internal/controllers/user"
	"API/internal/middlewares"
	"API/internal/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	if os.Getenv("GIN_MODE") == "release" {
		r.SetTrustedProxies([]string{os.Getenv("PROXY_IP")})
		panic("Database not setup for prod")
	}

	db := internal.GetDB()
	db.AutoMigrate(&models.User{})

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", controllers.Helloworld)
		v1.GET("/docs", func(c *gin.Context) {
			c.Redirect(http.StatusSeeOther, "/api/v1/docs/index.html")
		})
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		users := v1.Group("/users")
		{
			users.POST("/:username", user.Register)
			users.GET("/:username", middlewares.JwtAuthMiddleware(), user.Fetch)
		}
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/login", auth.Login)
		}
	}

	r.Run()
}
