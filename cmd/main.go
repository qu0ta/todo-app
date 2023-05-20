package main

import (
	"github.com/qu0ta/todo-app"
	"github.com/qu0ta/todo-app/pkg/handler"
	"github.com/qu0ta/todo-app/pkg/repository"
	"github.com/qu0ta/todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running server: %s", err.Error())
	}
}
