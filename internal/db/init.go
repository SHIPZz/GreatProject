package db

import (
	"GreatProject/internal/database"

	"github.com/jackc/pgx/v5"
)

func New(conn *pgx.Conn) *db.Queries {
	return db.New(conn)
}
