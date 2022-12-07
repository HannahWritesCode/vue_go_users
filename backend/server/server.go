package server

import "vue-go-users.com/config"

// Init initializes the Gin server
func Init() {
	conf := config.GetConfig()
	r := NewRouter()
	r.Run(conf.GetString("SERVER_PORT"))
}
