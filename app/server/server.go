package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo/app/service"
	"todo/app/task"
)

type Server interface {
	Initialize() error
	SetupRoutes()
}

type ServerImpl struct {
	mux         *http.ServeMux
	taskService service.TaskService
}

func NewHttpServer(taskService service.TaskService) *ServerImpl {
	serverImpl := &ServerImpl{
		mux:         http.NewServeMux(),
		taskService: taskService,
	}
	return serverImpl
}

func (serverImpl *ServerImpl) Initialize() error {
	err := http.ListenAndServe(":8080", serverImpl.mux)
	return err
}

func (serverImpl *ServerImpl) SetupRoutes() {
	serverImpl.mux.HandleFunc("/hello", hello)

	serverImpl.mux.HandleFunc(
		"/task/all",
		func(w http.ResponseWriter, r *http.Request) {
			tasks, err := serverImpl.taskService.ShowTasks()
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
			encoder := json.NewEncoder(w)
			encoder.Encode(tasks)
		})

	serverImpl.mux.HandleFunc(
		"/task/add",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "POST" {
				taskName := r.FormValue("taskName")

				task, err := serverImpl.taskService.AddTask(taskName)
				if err != nil {
					fmt.Fprintln(w, err)
					return
				}
				encoder := json.NewEncoder(w)
				encoder.Encode(task)
			}
		})

	serverImpl.mux.HandleFunc(
		"/task/complete",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "POST" {
				idString := r.FormValue("id")
				id, _ := strconv.Atoi(idString)

				task, err := serverImpl.taskService.CompleteTask(id)
				if err != nil {
					fmt.Fprintln(w, err)
					return
				}
				encoder := json.NewEncoder(w)
				encoder.Encode(task)
			}
		})

	serverImpl.mux.HandleFunc(
		"/task/incomplete",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "POST" {
				idString := r.FormValue("id")
				id, _ := strconv.Atoi(idString)
				task, err := serverImpl.taskService.IncompleteTask(id)
				if err != nil {
					fmt.Fprintln(w, err)
					return
				}

				encoder := json.NewEncoder(w)
				encoder.Encode(task)

			}
		})

	serverImpl.mux.HandleFunc(
		"/task/delete",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "POST" {
				idString := r.FormValue("id")
				id, _ := strconv.Atoi(idString)

				encoder := json.NewEncoder(w)
				encoder.Encode(serverImpl.taskService.DeleteTask(id))

			}
		})

	serverImpl.mux.HandleFunc(
		"/task/list",
		func(w http.ResponseWriter, r *http.Request) {
			response, _ := http.Get("http://localhost:8080/task/all")
			decoder := json.NewDecoder(response.Body)
			var tasks []task.Task
			for decoder.More() {
				decoder.Decode(&tasks)
			}
			fmt.Println(tasks)
		})

	serverImpl.mux.HandleFunc(
		"/task",
		func(w http.ResponseWriter, r *http.Request) {
			id, _ := strconv.Atoi(r.Header.Get("id"))

			task, err := serverImpl.taskService.FindById(id)
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
			encoder := json.NewEncoder(w)
			encoder.Encode(task)
		})
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}
