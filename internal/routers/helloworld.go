package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloworldD struct {
	Message string
}

// Helloworld godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} HelloworldD
// @Router /ping [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{"message": "hi"})
}
