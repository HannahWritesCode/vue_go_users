package models

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"vue-go-users.com/db"
	"vue-go-users.com/helpers"

	"github.com/golang-jwt/jwt/v4"
)

type (
	UserPublic struct {
		ID        uint   `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	User struct {
		ID                uint      `gorm:"column:id;primary_key" json:"id"`
		Email             string    `gorm:"type:varchar(255);NOT_NULL" json:"email"`
		FirstName         string    `gorm:"type:varchar(255)" sql:"default: null" json:"first_name"`
		LastName          string    `gorm:"type:varchar(255)" sql:"default: null" json:"last_name"`
		Password          string    `gorm:"type:varchar(255)" sql:"default: null" json:"password"`
		Token             string    `gorm:"type:varchar(255)" sql:"default: null" json:"token"`
		ResetToken        string    `gorm:"type:varchar(255)" sql:"default: null" json:"reset_token"`
		VerificationToken string    `gorm:"type:varchar(255)" sql:"default: null" json:"verification_token"`
		Verified          bool      `gorm:"type:boolean" sql:"default: false" json:"verified"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
	}
)

// TableName sets the table name
func (User) TableName() string {
	return "user"
}

func (User) List(args UserListArgs) *Pagination[User] {
	var users []User
	var totalRows int64

	db.GetDB().
		Model(&User{}).
		Count(&totalRows).
		Order("last_name ASC").
		Limit(args.Count).
		Offset(args.Count * (args.Page - 1)).
		Find(&users)

	output := Paginate(totalRows, int64(args.Count), int64(args.Page), users)

	return &output
}

func (user *User) Create() error {

	// Generate hash from the entered password
	err := HashAndSetPassword(user)
	if err != nil {
		panic(err)
	}

	// Add token
	user.Token = helpers.RandomString(60)

	// Create user in database
	err = db.GetDB().Create(&user).Error
	return err
}

func (User) GetByID(id uint) (User, error) {
	user := User{}

	query := db.GetDB().
		Model(&User{}).
		Where("id = ?", id).
		First(&user)

	return user, query.Error
}

func (User) GetbyEmail(email string) (User, error) {
	user := User{}

	query := db.GetDB().
		Model(&User{}).
		Where("email = ?", email).
		First(&user)

	return user, query.Error
}

func (User) GetByAuthToken(id uint, token string) (User, error) {
	user := User{}

	query := db.GetDB().Debug().
		Model(&User{}).
		Where("token = ? AND id = ?", token, id).
		First(&user)

	return user, query.Error
}

func (User) GetByResetToken(token string) (User, error) {
	user := User{}

	query := db.GetDB().Debug().
		Model(&User{}).
		Where("reset_token = ?", token).
		First(&user)

	return user, query.Error
}

func (user User) BuildJWT() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"token": user.Token,
		"id":    user.ID,
	})

	return token.SignedString([]byte(os.Getenv("APP_KEY")))
}

func (user User) BuildHexToken() (string, error) {
	// Generate slice of bytes
	bin := make([]byte, 64)
	_, err := rand.Read(bin)

	// Encode as hexadecimal token
	hex_token := hex.EncodeToString(bin)
	return hex_token, err
}

func (user *User) Update(newdata User) error {

	// Check that new password exists before hashing it
	if newdata.Password != "" {
		err := HashAndSetPassword(&newdata)
		if err != nil {
			panic(err)
		}
	}

	database := db.GetDB()

	// Update user
	err := database.Model(user).
		Where("id = ?", user.ID).
		Updates(newdata).
		Error
	if err != nil {
		return err
	}

	return err
}

func (user *User) Delete() error {
	err := db.GetDB().Delete(&user).Error

	return err
}

func HashAndSetPassword(user *User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return err
}
