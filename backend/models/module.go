package models

import (
	"time"

	"vue-go-users.com/db"
)

type (
	Module struct {
		ID        uint      `gorm:"column:id;primary_key" json:"id"`
		Name      string    `gorm:"type:varchar(255);NOT_NULL" json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		UserIds   []uint    `gorm:"-" json:"user_ids"`
		Users     []User    `gorm:"many2many:module_user;" json:"users"`
	}
)

// TableName sets the table name
func (Module) TableName() string {
	return "module"
}

func (Module) List() []Module {
	var modules []Module

	db.GetDB().
		Select("id, name").
		Model(&Module{}).
		Find(&modules)

	return modules
}

func (module *Module) Create() error {
	err := db.GetDB().Create(&module).Error
	return err
}

func (Module) GetByID(id uint) (Module, error) {
	module := Module{}

	query := db.GetDB().
		Model(&Module{}).
		Where("id = ?", id).
		First(&module)

	return module, query.Error
}

func (module *Module) Update(newdata Module) error {
	database := db.GetDB()

	// Update module
	err := database.Model(module).
		Where("id = ?", module.ID).
		Updates(newdata).
		Error
	if err != nil {
		return err
	}

	// Update users
	users := []User{}
	for _, u := range newdata.UserIds {
		users = append(users, User{ID: u})
	}

	database.Model(module).
		Association("User").
		Replace(users)

	return err
}

func (module *Module) Delete() error {
	err := db.GetDB().Delete(&module).Error

	return err
}
