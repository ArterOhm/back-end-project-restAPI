package main

import (
	"fmt"
	"os"

	"github.com/ArterOhm/back-end-project-restAPI/config"
	"github.com/ArterOhm/back-end-project-restAPI/pkg/database"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())

	db := database.DbConnect(cfg.Db())
	defer db.Close()
	fmt.Println(cfg.Db().Url())

	fmt.Println(db)
}
