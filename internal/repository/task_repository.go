package repository

import (
	"context"

	db "GreatProject/internal/database"

	"github.com/jackc/pgx/v5/pgtype"
)

type TaskRepository interface {
	GetAll(ctx context.Context, limit, offset int32) ([]*db.Task, error)
	GetByID(ctx context.Context, id int32) (*db.Task, error)
	Create(ctx context.Context, name, description string) (*db.Task, error)
	Update(ctx context.Context, id int32, name, description string, completed bool) (*db.Task, error)
	Delete(ctx context.Context, id int32) error
	Complete(ctx context.Context, id int32) (*db.Task, error)
	GetByStatus(ctx context.Context, completed bool, limit, offset int32) ([]*db.Task, error)
}

type taskRepository struct {
	queries *db.Queries
}

func NewTaskRepository(queries *db.Queries) TaskRepository {
	return &taskRepository{
		queries: queries,
	}
}

func (r *taskRepository) GetAll(ctx context.Context, limit, offset int32) ([]*db.Task, error) {
	return r.queries.ListTasks(ctx, db.ListTasksParams{
		Limit:  limit,
		Offset: offset,
	})
}

func (r *taskRepository) GetByID(ctx context.Context, id int32) (*db.Task, error) {
	return r.queries.GetTask(ctx, id)
}

func (r *taskRepository) Create(ctx context.Context, name, description string) (*db.Task, error) {
	return r.queries.CreateTask(ctx, db.CreateTaskParams{
		Name:        name,
		Description: pgtype.Text{String: description, Valid: description != ""},
		Completed:   pgtype.Bool{Bool: false, Valid: true},
	})
}

func (r *taskRepository) Update(ctx context.Context, id int32, name, description string, completed bool) (*db.Task, error) {
	return r.queries.UpdateTask(ctx, db.UpdateTaskParams{
		ID:          id,
		Name:        name,
		Description: pgtype.Text{String: description, Valid: description != ""},
		Completed:   pgtype.Bool{Bool: completed, Valid: true},
	})
}

func (r *taskRepository) Delete(ctx context.Context, id int32) error {
	return r.queries.DeleteTask(ctx, id)
}

func (r *taskRepository) Complete(ctx context.Context, id int32) (*db.Task, error) {
	return r.queries.CompleteTask(ctx, id)
}

func (r *taskRepository) GetByStatus(ctx context.Context, completed bool, limit, offset int32) ([]*db.Task, error) {
	return r.queries.ListTasksByStatus(ctx, db.ListTasksByStatusParams{
		Completed: pgtype.Bool{Bool: completed, Valid: true},
		Limit:     limit,
		Offset:    offset,
	})
}

