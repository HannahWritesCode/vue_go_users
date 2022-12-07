package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"

	"vue-go-users.com/config"
	"vue-go-users.com/db"
	"vue-go-users.com/db/cache"
	"vue-go-users.com/db/migrations"
	"vue-go-users.com/server"
)

func main() {
	env := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	// Config init
	config.Init(*env)
	conf := config.GetConfig()

	// Sentry init
	err := sentry.Init(sentry.ClientOptions{
		Dsn: conf.GetString("SENTRY_DSN"),
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	// DB init
	db.Init()
	cache.Init()
	migrations.Migrate()
	server.Init()

	defer db.Close()
}
