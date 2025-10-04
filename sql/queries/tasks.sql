-- name: GetTask :one
SELECT id, name, description, completed, created_at, updated_at 
FROM tasks 
WHERE id = $1;

-- name: ListTasks :many
SELECT id, name, description, completed, created_at, updated_at 
FROM tasks 
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListTasksByStatus :many
SELECT id, name, description, completed, created_at, updated_at 
FROM tasks 
WHERE completed = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: CreateTask :one
INSERT INTO tasks (name, description, completed)
VALUES ($1, $2, $3)
RETURNING id, name, description, completed, created_at, updated_at;

-- name: UpdateTask :one
UPDATE tasks 
SET name = $2, description = $3, completed = $4
WHERE id = $1
RETURNING id, name, description, completed, created_at, updated_at;

-- name: CompleteTask :one
UPDATE tasks 
SET completed = true
WHERE id = $1
RETURNING id, name, description, completed, created_at, updated_at;

-- name: DeleteTask :exec
DELETE FROM tasks 
WHERE id = $1;

-- name: CountTasks :one
SELECT COUNT(*) FROM tasks;

-- name: CountTasksByStatus :one
SELECT COUNT(*) FROM tasks WHERE completed = $1;
