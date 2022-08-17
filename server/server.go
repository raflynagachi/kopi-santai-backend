package server

import (
	"log"
)

func Init() {
	router := NewRouter(&RouterConfig{})
	log.Fatalln(router.Run())
}
