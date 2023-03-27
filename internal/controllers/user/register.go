package user

import (
	"API/internal"
	"API/internal/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// RegisterBody model info
// @Description Registration request body
type RegisterBody struct {
	Password string `json:"password" binding:"required,lte=75"`
}

// Register godoc
// @Summary Register a new user account
// @Schemes
// @Description Register a new user account
// @Tags User routes
// @Accept json
// @Param data body RegisterBody true "Request body"
// @Param username path string true "The username to register, max length 50."
// @Success 204
// @Router /users/{username} [post]
// @Failure 400
// @Security Bearer
func Register(c *gin.Context) {
	body := RegisterBody{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters."})
		return
	}

	name := c.Param("username")
	if len(name) > 50 {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": "Invalid parameters."},
		)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "Internal server error."},
		)
		log.Fatal("Bcrypt error:", err.Error())
		return
	}
	user := models.User{Username: name, Password: string(hashedPassword)}
	db := internal.GetDB()
	result := db.FirstOrCreate(&user)
	if result.Error != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "Internal server error."},
		)
		return
	}

	if result.RowsAffected == 1 {
		c.Status(204)
		return
	}

	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters."})

}
