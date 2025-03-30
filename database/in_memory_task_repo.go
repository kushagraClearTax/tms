package database

import (
	"github.com/kushagragupta/tms/models"
	"errors"
	"github.com/google/uuid"
	"fmt"
	"time"
)

type InMemoryTaskRepository struct {
	tasks map[string]*models.Task
}

// Using Repository design pattern to set repo for now as no db added
func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[string]*models.Task),
	}
}

func (r *InMemoryTaskRepository) CreateTask(task *models.Task) error {
 task.ID = uuid.New().String()
 task.CreatedAt = time.Now()
 task.UpdatedAt = time.Now()
 r.tasks[task.ID] = task
 return nil
}

func (r *InMemoryTaskRepository) GetTask(id string) (*models.Task, error) {
	fmt.Printf("Fetching task with ID: %s\n", id)
	task, ok := r.tasks[id]
	if !ok {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (r *InMemoryTaskRepository) GetTasks() ([]models.Task, error) {
	tasks := make([]models.Task, 0)
	for _, task := range r.tasks {
		tasks = append(tasks, *task)
	}
	return tasks, nil
}

func (r *InMemoryTaskRepository) UpdateTask(task *models.Task) error {
 _, ok := r.tasks[task.ID]
 if !ok {
  return errors.New("task not found")
 }
 task.UpdatedAt = time.Now()
 r.tasks[task.ID] = task
 return nil
}
func (r *InMemoryTaskRepository) DeleteTask(id string) error {
	_, ok := r.tasks[id]
	if !ok {
		return errors.New("task not found")
	}
	delete(r.tasks, id)
	return nil
}

func (r *InMemoryTaskRepository) GetTasksByStatus(status string) ([]models.Task, error) {
	var tasks []models.Task

	for _, task := range r.tasks {
		if task.Completed == (status == "Completed") {
			tasks = append(tasks, *task)
		}
	}

	return tasks, nil
}