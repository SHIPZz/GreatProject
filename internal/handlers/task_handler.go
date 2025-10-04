package handlers

import (
	"context"
	"errors"
	"net/http"

	db "GreatProject/internal/database"
	"GreatProject/internal/generated"
	"GreatProject/internal/service"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service service.TaskService
}

func NewTaskHandler(svc service.TaskService) *TaskHandler {
	return &TaskHandler{
		service: svc,
	}
}

// GetHealth проверка здоровья сервиса
func (h *TaskHandler) GetHealth(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"service": "todo-api",
	})
}

// GetTasks получить все задачи
func (h *TaskHandler) GetTasks(ctx echo.Context, params generated.GetTasksParams) error {
	limit := int32(50)
	offset := int32(0)

	if params.Limit != nil {
		limit = int32(*params.Limit)
	}
	if params.Offset != nil {
		offset = int32(*params.Offset)
	}

	tasks, err := h.service.GetAllTasks(context.Background(), limit, offset)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, generated.Error{
			Code:    "INTERNAL_ERROR",
			Message: "Failed to fetch tasks",
		})
	}

	// Конвертируем в формат API
	apiTasks := make([]generated.Task, len(tasks))
	for i, task := range tasks {
		apiTasks[i] = h.convertToAPITask(*task)
	}

	return ctx.JSON(http.StatusOK, apiTasks)
}

// PostTasks создать новую задачу
func (h *TaskHandler) PostTasks(ctx echo.Context) error {
	var req generated.CreateTaskRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.Error{
			Code:    "INVALID_REQUEST",
			Message: "Invalid request body",
		})
	}

	// Создаем задачу через сервис (валидация внутри)
	task, err := h.service.CreateTask(context.Background(), req.Name, req.Description)
	if err != nil {
		if errors.Is(err, service.ErrEmptyTaskName) {
			return ctx.JSON(http.StatusBadRequest, generated.Error{
				Code:    "VALIDATION_ERROR",
				Message: "Task name is required",
			})
		}
		if errors.Is(err, service.ErrInvalidTaskData) {
			return ctx.JSON(http.StatusBadRequest, generated.Error{
				Code:    "VALIDATION_ERROR",
				Message: "Task name is too long",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, generated.Error{
			Code:    "INTERNAL_ERROR",
			Message: "Failed to create task",
		})
	}

	return ctx.JSON(http.StatusCreated, h.convertToAPITask(*task))
}

// GetTasksCompleted получить выполненные задачи
func (h *TaskHandler) GetTasksCompleted(ctx echo.Context, params generated.GetTasksCompletedParams) error {
	limit := int32(50)
	offset := int32(0)

	if params.Limit != nil {
		limit = int32(*params.Limit)
	}
	if params.Offset != nil {
		offset = int32(*params.Offset)
	}

	tasks, err := h.service.GetCompletedTasks(context.Background(), limit, offset)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, generated.Error{
			Code:    "INTERNAL_ERROR",
			Message: "Failed to fetch completed tasks",
		})
	}

	apiTasks := make([]generated.Task, len(tasks))
	for i, task := range tasks {
		apiTasks[i] = h.convertToAPITask(*task)
	}

	return ctx.JSON(http.StatusOK, apiTasks)
}

// GetTasksPending получить невыполненные задачи
func (h *TaskHandler) GetTasksPending(ctx echo.Context, params generated.GetTasksPendingParams) error {
	limit := int32(50)
	offset := int32(0)

	if params.Limit != nil {
		limit = int32(*params.Limit)
	}
	if params.Offset != nil {
		offset = int32(*params.Offset)
	}

	tasks, err := h.service.GetPendingTasks(context.Background(), limit, offset)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, generated.Error{
			Code:    "INTERNAL_ERROR",
			Message: "Failed to fetch pending tasks",
		})
	}

	apiTasks := make([]generated.Task, len(tasks))
	for i, task := range tasks {
		apiTasks[i] = h.convertToAPITask(*task)
	}

	return ctx.JSON(http.StatusOK, apiTasks)
}

// GetTasksId получить задачу по ID
func (h *TaskHandler) GetTasksId(ctx echo.Context, id int) error {
	task, err := h.service.GetTaskByID(context.Background(), int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, generated.Error{
			Code:    "TASK_NOT_FOUND",
			Message: "Task not found",
		})
	}

	return ctx.JSON(http.StatusOK, h.convertToAPITask(*task))
}

// PutTasksId обновить задачу
func (h *TaskHandler) PutTasksId(ctx echo.Context, id int) error {
	var req generated.UpdateTaskRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.Error{
			Code:    "INVALID_REQUEST",
			Message: "Invalid request body",
		})
	}

	// Обновляем задачу через сервис (валидация внутри)
	task, err := h.service.UpdateTask(context.Background(), int32(id), req.Name, req.Description, req.Completed)
	if err != nil {
		if errors.Is(err, service.ErrEmptyTaskName) || errors.Is(err, service.ErrInvalidTaskData) {
			return ctx.JSON(http.StatusBadRequest, generated.Error{
				Code:    "VALIDATION_ERROR",
				Message: err.Error(),
			})
		}
		return ctx.JSON(http.StatusNotFound, generated.Error{
			Code:    "TASK_NOT_FOUND",
			Message: "Task not found",
		})
	}

	return ctx.JSON(http.StatusOK, h.convertToAPITask(*task))
}

// DeleteTasksId удалить задачу
func (h *TaskHandler) DeleteTasksId(ctx echo.Context, id int) error {
	err := h.service.DeleteTask(context.Background(), int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, generated.Error{
			Code:    "TASK_NOT_FOUND",
			Message: "Task not found",
		})
	}

	return ctx.NoContent(http.StatusNoContent)
}

// PatchTasksIdComplete отметить задачу выполненной
func (h *TaskHandler) PatchTasksIdComplete(ctx echo.Context, id int) error {
	task, err := h.service.CompleteTask(context.Background(), int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, generated.Error{
			Code:    "TASK_NOT_FOUND",
			Message: "Task not found",
		})
	}

	return ctx.JSON(http.StatusOK, h.convertToAPITask(*task))
}

// PatchTasksIdUncomplete снять отметку выполнения с задачи
func (h *TaskHandler) PatchTasksIdUncomplete(ctx echo.Context, id int) error {
	task, err := h.service.UncompleteTask(context.Background(), int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, generated.Error{
			Code:    "TASK_NOT_FOUND",
			Message: "Task not found",
		})
	}

	return ctx.JSON(http.StatusOK, h.convertToAPITask(*task))
}

// convertToAPITask конвертирует модель БД в API модель
func (h *TaskHandler) convertToAPITask(task db.Task) generated.Task {
	description := ""
	if task.Description.Valid {
		description = task.Description.String
	}

	completed := false
	if task.Completed.Valid {
		completed = task.Completed.Bool
	}

	return generated.Task{
		Id:          int(task.ID),
		Name:        task.Name,
		Description: description,
		Completed:   completed,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}
