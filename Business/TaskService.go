package Business

import (
	"GreatProject/Data/Entity"
	"GreatProject/Data/Repository"
)

type ITaskService interface {
	CreateTask(name, description string) (Entity.Task, error)
	GetAllTasks() []Entity.Task
	GetTaskByID(id int) (Entity.Task, bool)
	UpdateTask(id int, name, description string, completed bool) (Entity.Task, error)
	DeleteTask(id int) error
	CompleteTask(id int) error
	GetCompletedTasks() []Entity.Task
	GetPendingTasks() []Entity.Task
}

type TaskService struct {
	repository Repository.ITaskRepository
}

func NewTaskService(repository Repository.ITaskRepository) ITaskService {
	return &TaskService{
		repository: repository,
	}
}

func (ts *TaskService) CreateTask(name, description string) (Entity.Task, error) {
	task := Entity.Task{
		Name:        name,
		Description: description,
		Completed:   false,
	}

	if err := task.Validate(); err != nil {
		return Entity.Task{}, err
	}

	createdTask := ts.repository.Create(task)
	return createdTask, nil
}

func (ts *TaskService) GetAllTasks() []Entity.Task {
	return ts.repository.GetAll()
}

func (ts *TaskService) GetTaskByID(id int) (Entity.Task, bool) {
	return ts.repository.GetByID(id)
}

func (ts *TaskService) UpdateTask(id int, name, description string, completed bool) (Entity.Task, error) {
	task := Entity.Task{
		ID:          id,
		Name:        name,
		Description: description,
		Completed:   completed,
	}

	if err := task.Validate(); err != nil {
		return Entity.Task{}, err
	}

	updatedTask, success := ts.repository.Update(id, task)
	if !success {
		return Entity.Task{}, Entity.ErrInvalidID
	}

	return updatedTask, nil
}

func (ts *TaskService) DeleteTask(id int) error {
	success := ts.repository.Delete(id)
	if !success {
		return Entity.ErrInvalidID
	}
	return nil
}

func (ts *TaskService) CompleteTask(id int) error {
	task, exists := ts.repository.GetByID(id)
	if !exists {
		return Entity.ErrInvalidID
	}

	task.Completed = true
	_, success := ts.repository.Update(id, task)
	if !success {
		return Entity.ErrInvalidID
	}

	return nil
}

func (ts *TaskService) GetCompletedTasks() []Entity.Task {
	allTasks := ts.repository.GetAll()
	var completedTasks []Entity.Task

	for _, task := range allTasks {
		if task.Completed {
			completedTasks = append(completedTasks, task)
		}
	}

	return completedTasks
}

func (ts *TaskService) GetPendingTasks() []Entity.Task {
	allTasks := ts.repository.GetAll()
	var pendingTasks []Entity.Task

	for _, task := range allTasks {
		if !task.Completed {
			pendingTasks = append(pendingTasks, task)
		}
	}

	return pendingTasks
}



