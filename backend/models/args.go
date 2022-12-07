package models

type (

	// UserListArgs sets the "List" arguments
	UserListArgs struct {
		Count int `form:"count"`
		Page  int `form:"page"`
	}
)
