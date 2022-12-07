package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"vue-go-users.com/helpers"
	"vue-go-users.com/models"
)

func Get(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	userPublic, err := helpers.ConvertStruct[models.UserPublic](user)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, userPublic)
}
