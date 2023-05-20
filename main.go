package main

import (
	"fmt"

	"github.com/raflynagachi/kopi-santai-backend/config"
	"github.com/raflynagachi/kopi-santai-backend/db"
	"github.com/raflynagachi/kopi-santai-backend/server"
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
