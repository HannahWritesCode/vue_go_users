package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"vue-go-users.com/helpers"
	"vue-go-users.com/models"
	"vue-go-users.com/utils"

	_ "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type CreateParams struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func Login(c *gin.Context) {
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

	// Authenticate user

	// user = GetUserByEmail
	user := models.User{}
	user, getByEmailErr := user.GetbyEmail(input.Email)
	if getByEmailErr != nil {
		c.JSON(http.StatusUnprocessableEntity, []string{"Email not found."})
		return
	}

	// compare input.Password with user.Password using CompareHashAndPassword
	matchErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if matchErr != nil {
		// if passwords don't match, return 422
		c.JSON(http.StatusUnprocessableEntity, []string{"Password does not match our records."})
		return
	}

	userPublic, err := helpers.ConvertStruct[models.UserPublic](user)
	if err != nil {
		panic(err)
	}

	// build JWT
	token, buildErr := user.BuildJWT()
	if buildErr != nil {
		panic(buildErr.Error())
	}

	output := gin.H{
		"token": token,
		"user":  userPublic,
	}

	// return 200 and JWT
	c.JSON(http.StatusOK, output)
}
