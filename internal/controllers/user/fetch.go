package user

import (
	"API/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FetchData model info
// @Description Registration request body
type FetchData struct {
	Username string `json:"username" binding:"required,lte=50"`
}

// Fetch godoc
// @Summary Fetch a user account
// @Schemes
// @Description Register a new user account
// @Tags User routes
// @Param username path string true "The username to register, max length 50."
// @Success 200 {object} FetchData
// @Router /users/{username}  [get]
// @Failure 400
// @Security Bearer
func Fetch(c *gin.Context) {
	name := c.Param("username")
	user, err := models.GetOrError(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Invalid username."})
		return
	}
	// TODO Ensure JWT username matches current user requested

	c.JSON(200, FetchData{Username: user.Username})
}
