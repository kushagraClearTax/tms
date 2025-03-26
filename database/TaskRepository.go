package database

import (
	"github.com/kushagragupta/tms/models"
)

type TaskRepository interface {
	CreateTask(task *models.Task) error
	GetTask(id string) (*models.Task, error)
	GetTasks() ([]models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(id string) error
	GetTasksByStatus(status string) ([]models.Task, error)
}