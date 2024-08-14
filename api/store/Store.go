package store

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	Conn *pgxpool.Pool
}
