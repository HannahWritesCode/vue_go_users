package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"vue-go-users.com/models"
	"vue-go-users.com/services/mail"
	"vue-go-users.com/utils"
)

type CreateParams struct {
	Email           string `json:"email" validate:"required,email"`
	FirstName       string `json:"first_name" validate:"required"`
	LastName        string `json:"last_name" validate:"required"`
	Password        string `json:"password" validate:"required,eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"password_confirmation" validate:"required"`
}

func Create(c *gin.Context) {
	// Bind json
	var input CreateParams
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

	// Create
	user := models.User{
		Email:     input.Email, // note: when creating another user with the same email, 500 error is returned
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Password:  input.Password,
	}
	if err := user.Create(); err != nil {
		panic(err)
	}

	// Generate verification token
	verificationToken, err := user.BuildHexToken()
	if err != nil {
		panic(err)
	}

	// Update the user
	newdata := models.User{
		VerificationToken: verificationToken,
	}
	if err := user.Update(newdata); err != nil {
		panic(err)
	}

	// userPublic, err := helpers.ConvertStruct[models.UserPublic](user)
	// if err != nil {
	// 	panic(err)
	// }

	// // build JWT
	// token, buildErr := user.BuildJWT()
	// if buildErr != nil {
	// 	panic(buildErr.Error())
	// }

	// output := gin.H{
	// 	"token": token,
	// 	"user":  userPublic,
	// }

	c.JSON(http.StatusOK, verificationToken)

	// Send password reset email
	sendEmailErr := mail.SendVerificationEmail(input.Email, verificationToken)
	if sendEmailErr != nil {
		panic(sendEmailErr)
	}
}
