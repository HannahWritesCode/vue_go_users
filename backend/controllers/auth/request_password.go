package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"vue-go-users.com/models"
	"vue-go-users.com/services/mail"
	"vue-go-users.com/utils"
)

type RequestParams struct {
	Email string `json:"email" validate:"required,email"`
}

func SendResetPasswordEmail(c *gin.Context) {
	// Bind json
	var input RequestParams
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Validate
	validate = validator.New()
	if err := validate.Struct(input); err != nil {
		verrs := err.(validator.ValidationErrors)
		c.JSON(http.StatusUnprocessableEntity, utils.ValidationMessages(verrs))
		return
	}

	// Get user
	user, getByEmailErr := models.User{}.GetbyEmail(input.Email)
	if getByEmailErr != nil {
		c.JSON(http.StatusUnprocessableEntity, []string{"Email not found."})
		return
	}

	// // Generate temp password of 10 random letters
	// tempPassword := helpers.RandomString(10)

	// // Update the user
	// newdata := models.User{
	// 	Password: string(tempPassword),
	// }
	// if err := user.Update(newdata); err != nil {
	// 	panic(err)
	// }

	// c.JSON(http.StatusOK, tempPassword)

	// Generate reset token
	resetToken, err := user.BuildResetToken()
	if err != nil {
		panic(err)
	}

	// Update the user
	newdata := models.User{
		ResetToken: resetToken,
	}
	if err := user.Update(newdata); err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, resetToken)

	// Send password reset email
	sendEmailErr := mail.SendResetEmail(input.Email)
	if sendEmailErr != nil {
		panic(sendEmailErr)
	}
}
