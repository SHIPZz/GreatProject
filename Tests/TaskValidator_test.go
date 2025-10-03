package Tests

import (
	"GreatProject/Data/Entity"
	"testing"
)

func TestTask_Validate(t *testing.T) {
	task := Entity.Task{
		ID:          1,
		Name:        "Valid Task",
		Description: "Valid Description",
		Completed:   false,
	}

	err := task.Validate()

	if err != nil {
		t.Errorf("Expected no error for valid task, got %v", err)
	}
}

func TestTask_ValidateEmptyName(t *testing.T) {
	task := Entity.Task{
		ID:          1,
		Name:        "",
		Description: "Valid Description",
		Completed:   false,
	}

	err := task.Validate()

	if err != Entity.ErrEmptyName {
		t.Errorf("Expected ErrEmptyName, got %v", err)
	}
}

func TestTask_ValidateEmptyDescription(t *testing.T) {
	task := Entity.Task{
		ID:          1,
		Name:        "Valid Task",
		Description: "",
		Completed:   false,
	}

	err := task.Validate()

	if err != Entity.ErrEmptyDescription {
		t.Errorf("Expected ErrEmptyDescription, got %v", err)
	}
}

func TestTask_ValidateNegativeID(t *testing.T) {
	task := Entity.Task{
		ID:          -1,
		Name:        "Valid Task",
		Description: "Valid Description",
		Completed:   false,
	}

	err := task.Validate()

	if err != Entity.ErrInvalidID {
		t.Errorf("Expected ErrInvalidID, got %v", err)
	}
}

func TestTask_IsValid(t *testing.T) {
	validTask := Entity.Task{
		ID:          1,
		Name:        "Valid Task",
		Description: "Valid Description",
		Completed:   false,
	}

	if !validTask.IsValid() {
		t.Error("Expected valid task to be valid")
	}

	invalidTask := Entity.Task{
		ID:          1,
		Name:        "",
		Description: "Valid Description",
		Completed:   false,
	}

	if invalidTask.IsValid() {
		t.Error("Expected invalid task to be invalid")
	}
}



