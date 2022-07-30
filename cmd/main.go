package main

import (
	todo "github.com/cucumberjaye/go_todo"
	"github.com/cucumberjaye/go_todo/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}