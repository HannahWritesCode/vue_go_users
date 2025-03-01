package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"vue-go-users.com/models"
	"vue-go-users.com/utils"
)

type PasswordUpdateParams struct {
	// CurrentPassword string `json:"current_password" validate:"required"`
	ResetToken      string `json:"reset_token" validate:"required"`
	Password        string `json:"password" validate:"required,eqfield=ConfirmPassword"` // min length?
	ConfirmPassword string `json:"password_confirmation" validate:"required"`
}

func ResetPassword(c *gin.Context) {
	// Bind json
	var input PasswordUpdateParams
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Get current user
	// user := c.MustGet("user").(models.User)

	// Get user by reset token
	user, err := models.User{}.GetByResetToken(input.ResetToken)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, []string{"Reset token not found."})
		return
	}

	// Validate
	validate = validator.New()
	if err := validate.Struct(input); err != nil {
		verrs := err.(validator.ValidationErrors)
		c.JSON(http.StatusUnprocessableEntity, utils.ValidationMessages(verrs))
		return
	}

	// // Validate user's current password
	// matchErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.CurrentPassword))
	// if matchErr != nil {
	// 	// if passwords don't match, return 422
	// 	c.JSON(http.StatusUnprocessableEntity, []string{"Password does not match our records."})
	// 	return
	// }

	// Update the user
	newdata := models.User{
		Password:   input.Password,
		ResetToken: " ",
	}
	if err := user.Update(newdata); err != nil {
		panic(err)
	}

}
