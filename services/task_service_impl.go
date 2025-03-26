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


func (s *TaskServiceImpl) GetTasks(status string, limit int, offset int) ([]models.Task, error) {
	if status != "" {
		return s.repo.GetTasksByStatus(status)
	} else {
		tasks, err := s.repo.GetTasks()
		if err != nil {
			return nil, err
		}

		if limit > 0 {
			if offset+limit > len(tasks) {
				return tasks[offset:], nil
			}
			return tasks[offset : offset+limit], nil
		}

		return tasks, nil
	}
}

func (s *TaskServiceImpl) UpdateTask(task *models.Task) error {
	return s.repo.UpdateTask(task)
}

func (s *TaskServiceImpl) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}