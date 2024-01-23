package main

import (
	"log"
	"todo/app/server"
	"todo/app/service"
	"todo/app/usecase"
	"todo/db"
)

func main() {

	db, err := db.NewPostgresDb()
	if err != nil {
		log.Fatal(err)
	}
	taskUseCase := usecase.NewTaskUseCaseImpl(db)

	taskService := service.NewTaskServiceImpl(taskUseCase)

	server := server.NewHttpServer(taskService)
	server.SetupRoutes()
	server.Initialize()
}
