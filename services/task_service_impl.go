package services

import (
	"github.com/kushagragupta/tms/database"
	"github.com/kushagragupta/tms/models"
)

type TaskServiceImpl struct {
	repo database.TaskRepository
}

func NewTaskServiceImpl(repo database.TaskRepository) *TaskServiceImpl {
	return &TaskServiceImpl{
		repo: repo,
	}
}

func (s *TaskServiceImpl) CreateTask(task *models.Task) error {
	return s.repo.CreateTask(task)
}

func (s *TaskServiceImpl) GetTask(id string) (*models.Task, error) {
	return s.repo.GetTask(id)
}

func (s *TaskServiceImpl) GetTasks() ([]models.Task, error) {
	return s.repo.GetTasks()
}

func (s *TaskServiceImpl) UpdateTask(task *models.Task) error {
	return s.repo.UpdateTask(task)
}

func (s *TaskServiceImpl) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}