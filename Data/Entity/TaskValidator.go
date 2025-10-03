package Entity

import "errors"

var (
	ErrEmptyName        = errors.New("task name cannot be empty")
	ErrEmptyDescription = errors.New("task description cannot be empty")
	ErrInvalidID        = errors.New("task ID must be positive")
)

func (t *Task) Validate() error {
	if t.Name == "" {
		return ErrEmptyName
	}

	if t.Description == "" {
		return ErrEmptyDescription
	}

	if t.ID < 0 {
		return ErrInvalidID
	}
	return nil
}

func (t *Task) IsValid() bool {
	return t.Validate() == nil
}



