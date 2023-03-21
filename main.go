package main

import (
	docs "API/docs"
	"API/internal/routers"
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

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	if os.Getenv("GIN_MODE") == "release" {
		r.SetTrustedProxies([]string{"172.18.0.3"})
	}

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", routers.Helloworld)
		v1.GET("/docs", func(c *gin.Context) {
			c.Redirect(http.StatusSeeOther, "/api/v1/docs/index.html")
		})
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	r.Run()
}
