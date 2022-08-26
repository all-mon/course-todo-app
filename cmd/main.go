package main

import (
	todo "github.com/m0n7h0ff/course-todo-app"
	"github.com/m0n7h0ff/course-todo-app/pkg/handler"
	"github.com/m0n7h0ff/course-todo-app/pkg/repository"
	"github.com/m0n7h0ff/course-todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}
