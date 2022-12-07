package main

import (
	"flag"
	"fmt"
	"os"

	"vue-go-users.com/config"
	"vue-go-users.com/db"
	"vue-go-users.com/db/migrations"
)

func main() {
	env := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	config.Init(*env)

	db.Init()

	migrations.Migrate()
}
