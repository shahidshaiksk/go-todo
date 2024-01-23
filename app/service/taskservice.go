package service

import (
	"todo/app/task"
	"todo/app/usecase"
)

type TaskService interface {
	FindById(id int) (task.Task, error)
	ShowTasks() ([]task.Task, error)
	AddTask(taskName string) (task.Task, error)
	CompleteTask(id int) (task.Task, error)
	IncompleteTask(id int) (task.Task, error)
	DeleteTask(id int) string
}

type TaskServiceImpl struct {
	taskUseCase usecase.TaskUseCase
}

func NewTaskServiceImpl(taskUseCase usecase.TaskUseCase) *TaskServiceImpl {
	taskServiceImpl := &TaskServiceImpl{
		taskUseCase: taskUseCase,
	}
	return taskServiceImpl
}

func (taskServiceImpl *TaskServiceImpl) ShowTasks() ([]task.Task, error) {
	tasks, err := taskServiceImpl.taskUseCase.Select()
	return tasks, err
}

func (taskServiceImpl *TaskServiceImpl) AddTask(taskName string) (task.Task, error) {
	task, err := taskServiceImpl.taskUseCase.Create(taskName)
	return task, err
}

func (taskServiceImpl *TaskServiceImpl) DeleteTask(id int) string {
	err := taskServiceImpl.taskUseCase.DeleteById(id)
	if err != nil {
		return err.Error()
	}
	return "Task deleted successfully"
}

func (taskServiceImpl *TaskServiceImpl) CompleteTask(id int) (task.Task, error) {
	task, err := taskServiceImpl.taskUseCase.FindById(id)
	if err != nil {
		return task, err
	}
	task, err = taskServiceImpl.taskUseCase.Update(task.Id, task.Name, true)
	return task, err
}

func (taskServiceImpl *TaskServiceImpl) IncompleteTask(id int) (task.Task, error) {
	task, err := taskServiceImpl.taskUseCase.FindById(id)
	if err != nil {
		return task, err
	}
	task, err = taskServiceImpl.taskUseCase.Update(task.Id, task.Name, false)
	return task, err
}

func (taskServiceImpl *TaskServiceImpl) FindById(id int) (task.Task, error) {
	task, err := taskServiceImpl.taskUseCase.FindById(id)
	return task, err
}
