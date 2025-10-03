package Tests

import (
	"GreatProject/Business"
	"GreatProject/Data/Entity"
	"GreatProject/Infrastructure/Logger"
	"testing"
)

func TestTaskService_CreateTask(t *testing.T) {
	mockRepo := &MockTaskRepository{
		tasks:  make(map[int]Entity.Task),
		nextID: 1,
		logger: Logger.NewMockLogger(),
	}

	service := Business.NewTaskService(mockRepo)

	task, err := service.CreateTask("Test Task", "Test Description")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if task.Name != "Test Task" {
		t.Errorf("Expected 'Test Task', got '%s'", task.Name)
	}
}

func TestTaskService_CreateTaskInvalid(t *testing.T) {
	mockRepo := &MockTaskRepository{
		tasks:  make(map[int]Entity.Task),
		nextID: 1,
		logger: Logger.NewMockLogger(),
	}

	service := Business.NewTaskService(mockRepo)

	_, err := service.CreateTask("", "Test Description")

	if err == nil {
		t.Error("Expected validation error for empty name")
	}
}

func TestTaskService_GetAllTasks(t *testing.T) {
	mockRepo := &MockTaskRepository{
		tasks:  make(map[int]Entity.Task),
		nextID: 1,
		logger: Logger.NewMockLogger(),
	}

	service := Business.NewTaskService(mockRepo)

	mockRepo.tasks[1] = Entity.Task{ID: 1, Name: "Task 1", Description: "Desc 1"}
	mockRepo.tasks[2] = Entity.Task{ID: 2, Name: "Task 2", Description: "Desc 2"}

	tasks := service.GetAllTasks()

	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
}

func TestTaskService_CompleteTask(t *testing.T) {
	mockRepo := &MockTaskRepository{
		tasks:  make(map[int]Entity.Task),
		nextID: 1,
		logger: Logger.NewMockLogger(),
	}

	service := Business.NewTaskService(mockRepo)

	task := Entity.Task{ID: 1, Name: "Test Task", Description: "Test Description", Completed: false}
	mockRepo.tasks[1] = task

	err := service.CompleteTask(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	updatedTask, _ := mockRepo.GetByID(1)
	if !updatedTask.Completed {
		t.Error("Expected task to be completed")
	}
}

func TestTaskService_CompleteTaskNotFound(t *testing.T) {
	mockRepo := &MockTaskRepository{
		tasks:  make(map[int]Entity.Task),
		nextID: 1,
		logger: Logger.NewMockLogger(),
	}

	service := Business.NewTaskService(mockRepo)

	err := service.CompleteTask(999)

	if err == nil {
		t.Error("Expected error for non-existent task")
	}
}

func TestTaskService_GetCompletedTasks(t *testing.T) {
	mockRepo := &MockTaskRepository{
		tasks:  make(map[int]Entity.Task),
		nextID: 1,
		logger: Logger.NewMockLogger(),
	}

	service := Business.NewTaskService(mockRepo)

	mockRepo.tasks[1] = Entity.Task{ID: 1, Name: "Task 1", Description: "Desc 1", Completed: true}
	mockRepo.tasks[2] = Entity.Task{ID: 2, Name: "Task 2", Description: "Desc 2", Completed: false}

	completedTasks := service.GetCompletedTasks()

	if len(completedTasks) != 1 {
		t.Errorf("Expected 1 completed task, got %d", len(completedTasks))
	}

	if completedTasks[0].Name != "Task 1" {
		t.Errorf("Expected 'Task 1', got '%s'", completedTasks[0].Name)
	}
}

func TestTaskService_GetPendingTasks(t *testing.T) {
	mockRepo := &MockTaskRepository{
		tasks:  make(map[int]Entity.Task),
		nextID: 1,
		logger: Logger.NewMockLogger(),
	}

	service := Business.NewTaskService(mockRepo)

	mockRepo.tasks[1] = Entity.Task{ID: 1, Name: "Task 1", Description: "Desc 1", Completed: true}
	mockRepo.tasks[2] = Entity.Task{ID: 2, Name: "Task 2", Description: "Desc 2", Completed: false}

	pendingTasks := service.GetPendingTasks()

	if len(pendingTasks) != 1 {
		t.Errorf("Expected 1 pending task, got %d", len(pendingTasks))
	}

	if pendingTasks[0].Name != "Task 2" {
		t.Errorf("Expected 'Task 2', got '%s'", pendingTasks[0].Name)
	}
}

type MockTaskRepository struct {
	tasks  map[int]Entity.Task
	nextID int
	logger Logger.ILogger
}

func (m *MockTaskRepository) GetAll() []Entity.Task {
	tasks := make([]Entity.Task, 0, len(m.tasks))
	for _, task := range m.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (m *MockTaskRepository) GetByID(id int) (Entity.Task, bool) {
	task, exists := m.tasks[id]
	return task, exists
}

func (m *MockTaskRepository) Create(task Entity.Task) Entity.Task {
	if err := task.Validate(); err != nil {
		return Entity.Task{}
	}

	task.ID = m.nextID
	m.tasks[m.nextID] = task
	m.nextID++
	return task
}

func (m *MockTaskRepository) Update(id int, task Entity.Task) (Entity.Task, bool) {
	if _, exists := m.tasks[id]; !exists {
		return Entity.Task{}, false
	}

	task.ID = id
	m.tasks[id] = task
	return task, true
}

func (m *MockTaskRepository) Delete(id int) bool {
	if _, exists := m.tasks[id]; !exists {
		return false
	}

	delete(m.tasks, id)
	return true
}
