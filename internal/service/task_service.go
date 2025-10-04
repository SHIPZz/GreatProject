package service

import (
	"context"
	"errors"

	db "GreatProject/internal/database"
	"GreatProject/internal/repository"
)

var (
	ErrTaskNotFound    = errors.New("task not found")
	ErrInvalidTaskData = errors.New("invalid task data")
	ErrEmptyTaskName   = errors.New("task name cannot be empty")
)

type TaskService interface {
	GetAllTasks(ctx context.Context, limit, offset int32) ([]*db.Task, error)
	GetTaskByID(ctx context.Context, id int32) (*db.Task, error)
	CreateTask(ctx context.Context, name, description string) (*db.Task, error)
	UpdateTask(ctx context.Context, id int32, name, description string, completed bool) (*db.Task, error)
	DeleteTask(ctx context.Context, id int32) error
	CompleteTask(ctx context.Context, id int32) (*db.Task, error)
	GetCompletedTasks(ctx context.Context, limit, offset int32) ([]*db.Task, error)
	GetPendingTasks(ctx context.Context, limit, offset int32) ([]*db.Task, error)
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{
		repo: repo,
	}
}

func (s *taskService) GetAllTasks(ctx context.Context, limit, offset int32) ([]*db.Task, error) {
	return s.repo.GetAll(ctx, limit, offset)
}

func (s *taskService) GetTaskByID(ctx context.Context, id int32) (*db.Task, error) {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrTaskNotFound
	}
	return task, nil
}

func (s *taskService) CreateTask(ctx context.Context, name, description string) (*db.Task, error) {
	if name == "" {
		return nil, ErrEmptyTaskName
	}

	if len(name) > 255 {
		return nil, ErrInvalidTaskData
	}

	return s.repo.Create(ctx, name, description)
}

func (s *taskService) UpdateTask(ctx context.Context, id int32, name, description string, completed bool) (*db.Task, error) {
	if name == "" {
		return nil, ErrEmptyTaskName
	}

	if len(name) > 255 {
		return nil, ErrInvalidTaskData
	}

	task, err := s.repo.Update(ctx, id, name, description, completed)
	if err != nil {
		return nil, ErrTaskNotFound
	}
	return task, nil
}

func (s *taskService) DeleteTask(ctx context.Context, id int32) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return ErrTaskNotFound
	}
	return nil
}

func (s *taskService) CompleteTask(ctx context.Context, id int32) (*db.Task, error) {
	task, err := s.repo.Complete(ctx, id)
	if err != nil {
		return nil, ErrTaskNotFound
	}
	return task, nil
}

func (s *taskService) GetCompletedTasks(ctx context.Context, limit, offset int32) ([]*db.Task, error) {
	return s.repo.GetByStatus(ctx, true, limit, offset)
}

func (s *taskService) GetPendingTasks(ctx context.Context, limit, offset int32) ([]*db.Task, error) {
	return s.repo.GetByStatus(ctx, false, limit, offset)
}

