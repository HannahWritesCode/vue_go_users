package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"vue-go-users.com/helpers"
	"vue-go-users.com/models"
	"vue-go-users.com/utils"
)

type UserUpdateParams struct {
	Email           string `json:"email" validate:"omitempty,email"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Password        string `json:"password" validate:"omitempty,eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"password_confirmation"`
}

func Update(c *gin.Context) {
	// Bind json
	var input UserUpdateParams
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user := c.MustGet("user").(models.User)

	// Validate
	validate = validator.New()
	if err := validate.Struct(input); err != nil {
		verrs := err.(validator.ValidationErrors)
		c.JSON(http.StatusUnprocessableEntity, utils.ValidationMessages(verrs))
		return
	}

	// Update the user
	newdata := models.User{
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
	if input.Password != "" {
		newdata.Password = input.Password
	}
	err := user.Update(newdata)
	if err != nil {
		panic(err)
	}

	userPublic, err := helpers.ConvertStruct[models.UserPublic](user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, userPublic)
}
