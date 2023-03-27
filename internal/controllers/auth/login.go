package auth

import (
	"API/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginInput struct {
	Username string `json:"username" binding:"required,lte=50"`
	Password string `json:"password" binding:"required,lte=75"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.GetOrError(input.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	token, err := models.LoginCheck(user, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
