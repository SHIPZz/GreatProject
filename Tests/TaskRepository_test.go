package Tests

import (
	"GreatProject/Data/Entity"
	"GreatProject/Data/Repository"
	"testing"
)

func TestInMemoryTaskRepository_Create(t *testing.T) {
	repo := Repository.NewInMemoryTaskRepository()

	task := Entity.Task{
		Name:        "Test Task",
		Description: "Test Description",
		Completed:   false,
	}

	createdTask := repo.Create(task)

	if createdTask.ID != 1 {
		t.Errorf("Expected ID 1, got %d", createdTask.ID)
	}

	if createdTask.Name != "Test Task" {
		t.Errorf("Expected name 'Test Task', got '%s'", createdTask.Name)
	}
}

func TestInMemoryTaskRepository_CreateInvalidTask(t *testing.T) {
	repo := Repository.NewInMemoryTaskRepository()

	invalidTask := Entity.Task{
		Name:        "",
		Description: "Test Description",
		Completed:   false,
	}

	createdTask := repo.Create(invalidTask)

	if createdTask.ID != 0 {
		t.Errorf("Expected empty task for invalid input, got ID %d", createdTask.ID)
	}
}

func TestInMemoryTaskRepository_GetAll(t *testing.T) {
	repo := Repository.NewInMemoryTaskRepository()

	task1 := Entity.Task{Name: "Task 1", Description: "Desc 1"}
	task2 := Entity.Task{Name: "Task 2", Description: "Desc 2"}

	repo.Create(task1)
	repo.Create(task2)

	tasks := repo.GetAll()

	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
}

func TestInMemoryTaskRepository_GetByID(t *testing.T) {
	repo := Repository.NewInMemoryTaskRepository()

	task := Entity.Task{Name: "Test Task", Description: "Test Description"}
	createdTask := repo.Create(task)

	retrievedTask, exists := repo.GetByID(createdTask.ID)

	if !exists {
		t.Error("Expected task to exist")
	}

	if retrievedTask.Name != "Test Task" {
		t.Errorf("Expected 'Test Task', got '%s'", retrievedTask.Name)
	}
}

func TestInMemoryTaskRepository_GetByIDNotFound(t *testing.T) {
	repo := Repository.NewInMemoryTaskRepository()

	_, exists := repo.GetByID(999)

	if exists {
		t.Error("Expected task not to exist")
	}
}

func TestInMemoryTaskRepository_Update(t *testing.T) {
	repo := Repository.NewInMemoryTaskRepository()

	task := Entity.Task{Name: "Original Task", Description: "Original Description"}
	createdTask := repo.Create(task)

	updatedTask := Entity.Task{
		ID:          createdTask.ID,
		Name:        "Updated Task",
		Description: "Updated Description",
		Completed:   true,
	}

	result, success := repo.Update(createdTask.ID, updatedTask)

	if !success {
		t.Error("Expected update to succeed")
	}

	if result.Name != "Updated Task" {
		t.Errorf("Expected 'Updated Task', got '%s'", result.Name)
	}
}

func TestInMemoryTaskRepository_UpdateNotFound(t *testing.T) {
	repo := Repository.NewInMemoryTaskRepository()

	task := Entity.Task{Name: "Test Task", Description: "Test Description"}

	_, success := repo.Update(999, task)

	if success {
		t.Error("Expected update to fail for non-existent task")
	}
}

func TestInMemoryTaskRepository_Delete(t *testing.T) {
	repo := Repository.NewInMemoryTaskRepository()

	task := Entity.Task{Name: "Test Task", Description: "Test Description"}
	createdTask := repo.Create(task)

	success := repo.Delete(createdTask.ID)

	if !success {
		t.Error("Expected delete to succeed")
	}

	_, exists := repo.GetByID(createdTask.ID)
	if exists {
		t.Error("Expected task to be deleted")
	}
}

func TestInMemoryTaskRepository_DeleteNotFound(t *testing.T) {
	repo := Repository.NewInMemoryTaskRepository()

	success := repo.Delete(999)

	if success {
		t.Error("Expected delete to fail for non-existent task")
	}
}
