package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	r := gin.Default()
	if os.Getenv("GIN_MODE") == "release" {
		r.SetTrustedProxies([]string{"172.18.0.3"})
	}

	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "home",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
