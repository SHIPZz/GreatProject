package Repository

import (
	"GreatProject/Data/Entity"
	"GreatProject/Infrastructure/Logger"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ITaskRepository interface {
	GetAll() []Entity.Task
	GetByID(id int) (Entity.Task, bool)
	Create(task Entity.Task) Entity.Task
	Update(id int, task Entity.Task) (Entity.Task, bool)
	Delete(id int) bool
}

type InMemoryTaskRepository struct {
	Tasks  map[int]Entity.Task
	NextID int
	Logger Logger.ILogger
}

func (i *InMemoryTaskRepository) GetAll() []Entity.Task {
	i.Logger.Debug("Getting all tasks")
	tasks := make([]Entity.Task, 0, len(i.Tasks))
	for _, task := range i.Tasks {
		tasks = append(tasks, task)
	}
	i.Logger.Info(fmt.Sprintf("Retrieved %d tasks", len(tasks)))
	return tasks
}

func (i *InMemoryTaskRepository) GetByID(id int) (Entity.Task, bool) {
	i.Logger.Debug(fmt.Sprintf("Getting task by ID: %d", id))
	task, exists := i.Tasks[id]
	if exists {
		i.Logger.Info(fmt.Sprintf("Task found: %s", task.Name))
	} else {
		i.Logger.Warn(fmt.Sprintf("Task not found with ID: %d", id))
	}
	return task, exists
}

func (i *InMemoryTaskRepository) Create(task Entity.Task) Entity.Task {
	i.Logger.Debug(fmt.Sprintf("Creating task: %s", task.Name))
	if err := task.Validate(); err != nil {
		i.Logger.Error(fmt.Sprintf("Task validation failed: %v", err))
		return Entity.Task{}
	}

	task.ID = i.NextID
	i.Tasks[i.NextID] = task
	i.NextID++
	i.Logger.Info(fmt.Sprintf("Task created with ID: %d", task.ID))
	return task
}

func (i *InMemoryTaskRepository) Update(id int, task Entity.Task) (Entity.Task, bool) {
	i.Logger.Debug(fmt.Sprintf("Updating task with ID: %d", id))
	if _, exists := i.Tasks[id]; !exists {
		i.Logger.Warn(fmt.Sprintf("Task not found for update with ID: %d", id))
		return Entity.Task{}, false
	}

	task.ID = id
	i.Tasks[id] = task
	i.Logger.Info(fmt.Sprintf("Task updated: %s", task.Name))
	return task, true
}

func (i *InMemoryTaskRepository) Delete(id int) bool {
	i.Logger.Debug(fmt.Sprintf("Deleting task with ID: %d", id))
	if _, exists := i.Tasks[id]; !exists {
		i.Logger.Warn(fmt.Sprintf("Task not found for deletion with ID: %d", id))
		return false
	}

	delete(i.Tasks, id)
	i.Logger.Info(fmt.Sprintf("Task deleted with ID: %d", id))
	return true
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		Tasks:  make(map[int]Entity.Task),
		NextID: 1,
		Logger: Logger.NewLogger(),
	}
}

type PostgresTaskRepository struct {
	DB     *sqlx.DB
	Logger Logger.ILogger
}

func (p *PostgresTaskRepository) Close() error {
	return p.DB.Close()
}

func (p *PostgresTaskRepository) GetAll() []Entity.Task {
	p.Logger.Debug("Getting all tasks from PostgreSQL")
	var tasks []Entity.Task
	err := p.DB.Select(&tasks, "SELECT id, name, description, completed FROM tasks")
	if err != nil {
		p.Logger.Error(fmt.Sprintf("Error getting all tasks: %v", err))
		return []Entity.Task{}
	}
	p.Logger.Info(fmt.Sprintf("Retrieved %d tasks from database", len(tasks)))
	return tasks
}

func (p *PostgresTaskRepository) GetByID(id int) (Entity.Task, bool) {
	p.Logger.Debug(fmt.Sprintf("Getting task by ID: %d from PostgreSQL", id))
	var task Entity.Task
	err := p.DB.Get(&task, "SELECT id, name, description, completed FROM tasks WHERE id = $1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			p.Logger.Warn(fmt.Sprintf("Task not found with ID: %d", id))
			return Entity.Task{}, false
		}
		p.Logger.Error(fmt.Sprintf("Error getting task by ID: %v", err))
		return Entity.Task{}, false
	}
	p.Logger.Info(fmt.Sprintf("Task found: %s", task.Name))
	return task, true
}

func (p *PostgresTaskRepository) Create(task Entity.Task) Entity.Task {
	p.Logger.Debug(fmt.Sprintf("Creating task: %s in PostgreSQL", task.Name))
	if err := task.Validate(); err != nil {
		p.Logger.Error(fmt.Sprintf("Task validation failed: %v", err))
		return Entity.Task{}
	}

	var id int
	err := p.DB.QueryRow(
		"INSERT INTO tasks (name, description, completed) VALUES ($1, $2, $3) RETURNING id",
		task.Name, task.Description, task.Completed,
	).Scan(&id)

	if err != nil {
		p.Logger.Error(fmt.Sprintf("Error creating task: %v", err))
		return Entity.Task{}
	}

	task.ID = id
	p.Logger.Info(fmt.Sprintf("Task created with ID: %d", task.ID))
	return task
}

func (p *PostgresTaskRepository) Update(id int, task Entity.Task) (Entity.Task, bool) {
	p.Logger.Debug(fmt.Sprintf("Updating task with ID: %d in PostgreSQL", id))

	result, err := p.DB.Exec(
		"UPDATE tasks SET name = $1, description = $2, completed = $3 WHERE id = $4",
		task.Name, task.Description, task.Completed, id,
	)

	if err != nil {
		p.Logger.Error(fmt.Sprintf("Error updating task: %v", err))
		return Entity.Task{}, false
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		p.Logger.Warn(fmt.Sprintf("Task not found for update with ID: %d", id))
		return Entity.Task{}, false
	}

	task.ID = id
	p.Logger.Info(fmt.Sprintf("Task updated: %s", task.Name))
	return task, true
}

func (p *PostgresTaskRepository) Delete(id int) bool {
	p.Logger.Debug(fmt.Sprintf("Deleting task with ID: %d from PostgreSQL", id))

	result, err := p.DB.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		p.Logger.Error(fmt.Sprintf("Error deleting task: %v", err))
		return false
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		p.Logger.Warn(fmt.Sprintf("Task not found for deletion with ID: %d", id))
		return false
	}

	p.Logger.Info(fmt.Sprintf("Task deleted with ID: %d", id))
	return true
}

func NewPostgresTaskRepository(connectionString string) (*PostgresTaskRepository, error) {
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &PostgresTaskRepository{
		DB:     db,
		Logger: Logger.NewLogger(),
	}, nil
}
