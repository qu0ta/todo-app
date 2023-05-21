package main

import (
	"github.com/joho/godotenv"
	"github.com/qu0ta/todo-app"
	"github.com/qu0ta/todo-app/pkg/handler"
	"github.com/qu0ta/todo-app/pkg/repository"
	"github.com/qu0ta/todo-app/pkg/service"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs, %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables, %s", err.Error())

	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("host"),
		Port:     viper.GetString("port"),
		Username: viper.GetString("username"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   viper.GetString("dbname"),
		SSLMode:  viper.GetString("sslmode"),
	})
	if err != nil {
		log.Fatalf("error create postgres db: %s", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running server: %s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
