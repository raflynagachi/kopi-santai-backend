package main

import (
	"fmt"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/config"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/db"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/server"
)

func main() {
	fmt.Println(config.Config)
	err := db.Connect()
	if err != nil {
		fmt.Println("failed to connect to DB")
		return
	}

	server.Init()
}
