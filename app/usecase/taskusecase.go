package usecase

import (
	"fmt"
	"todo/app/task"
	"todo/db"
	"todo/errors"
)

type TaskUseCase interface {
	Create(taskName string) (task.Task, error)
	Select() ([]task.Task, error)
	Update(id int, taskName string, completed bool) (task.Task, error)
	FindById(id int) (task.Task, error)
	DeleteById(id int) error
}

type TaskUseCaseImpl struct {
	db        db.DB
	tableName string
}

func NewTaskUseCaseImpl(db db.DB) *TaskUseCaseImpl {
	taskUseCaseImpl := &TaskUseCaseImpl{
		db:        db,
		tableName: "tasks",
	}
	return taskUseCaseImpl
}

func (taskUseCaseImpl *TaskUseCaseImpl) FindById(id int) (task.Task, error) {
	rows, err := taskUseCaseImpl.db.SelectWithCondition(
		taskUseCaseImpl.tableName,
		fmt.Sprintf("id = %v", id))

	var task task.Task
	if err != nil {
		return task, err
	}
	count := 0
	for rows.Next() {
		err = rows.Scan(&task.Id, &task.Name, &task.Completed)
		if err != nil {
			return task, err
		}
		count++
	}
	if count == 0 {
		return task, errors.NoTaskError()
	}
	return task, nil
}

func (taskUseCaseImpl *TaskUseCaseImpl) Create(taskName string) (task.Task, error) {
	task := task.Task{
		Name:      taskName,
		Completed: false,
	}
	return task, taskUseCaseImpl.db.Insert(
		taskUseCaseImpl.tableName,
		map[string]string{
			"name":      fmt.Sprintf("'%v'", taskName),
			"completed": "false",
		})
}

func (taskUseCaseImpl *TaskUseCaseImpl) Select() ([]task.Task, error) {
	rows, err := taskUseCaseImpl.db.Select(taskUseCaseImpl.tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	count := 0
	tasks := []task.Task{}
	for rows.Next() {
		var task task.Task
		err := rows.Scan(&task.Id, &task.Name, &task.Completed)
		if err != nil {
			return nil, err
		}
		count++
		tasks = append(tasks, task)
	}
	if count == 0 {
		return tasks, errors.NoTaskError()
	}
	return tasks, nil
}

func (taskUseCaseImpl *TaskUseCaseImpl) Update(id int, taskName string, completed bool) (task.Task, error) {
	args := map[string]string{
		"name":      fmt.Sprintf("'%v'", taskName),
		"completed": fmt.Sprintf("%v", completed),
	}
	task := task.Task{
		Id:        id,
		Name:      taskName,
		Completed: completed,
	}
	err := taskUseCaseImpl.db.Update(
		taskUseCaseImpl.tableName,
		fmt.Sprintf("id = %v", id),
		args,
	)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (taskUseCaseImpl *TaskUseCaseImpl) DeleteById(id int) error {
	err := taskUseCaseImpl.db.DeleteWithCondition(taskUseCaseImpl.tableName, fmt.Sprintf("id = %v", id))
	if err != nil {
		return errors.DbError()
	}
	return nil
}
