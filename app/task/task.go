package task

import "fmt"

type Task struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

func (task Task) String() string {
	return fmt.Sprintf("Task { id = %v, name = %v, completed = %v}\n", task.Id, task.Name, task.Completed)
}
