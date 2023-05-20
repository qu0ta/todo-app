package main

import (
	"fmt"
	"github.com/qu0ta/todo-app"
	"github.com/qu0ta/todo-app/pkg/handler"
	"log"
)

func main() {
	handler := new(handler.Handler)
	srv := new(todo.Server)
	fmt.Println("Start!")

	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error while running server: %s", err.Error())
	}
}