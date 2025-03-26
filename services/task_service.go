package services

import (
	"github.com/kushagragupta/tms/models"
)

type TaskService interface {
	CreateTask(task *models.Task) error
	GetTask(id string) (*models.Task, error)
	GetTasks(status string, limit int, offset int) ([]models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(id string) error
}