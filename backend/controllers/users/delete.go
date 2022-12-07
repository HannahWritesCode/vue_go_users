package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"vue-go-users.com/models"
)

// Delete a user
func Delete(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	// Delete
	if err := user.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
