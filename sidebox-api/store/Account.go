package store

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id        string
	Name      string
	Email     string
	GoogleId  string
	CreatedAt time.Time
}

func (s Store) AccountGet(ctx context.Context, id string) (Account, error) {
	var account Account
	err := s.Conn.QueryRow(ctx, "select id, name, email, google_id, created_at from account where id=$1", id).
		Scan(&account.Id, &account.Name, &account.Email, &account.GoogleId, &account.CreatedAt)
	return account, err
}

func (s Store) AccountCreate(ctx context.Context, name string, email string, googleId string, ref string) (Account, error) {
	var account Account
	err := s.Conn.QueryRow(ctx, "insert into account (id, name, email, google_id, ref) "+
		"values ($1, $2, $3, $4, $5) "+
		"on conflict (google_id) do update set name=$2, email=$3 "+
		"returning id, name, email google_id, created_at",
		uuid.New(), name, email, googleId, ref).
		Scan(&account.Id, &account.Name, &account.Email, &account.CreatedAt)
	return account, err
}

func (s Store) AccountSetLastUsedAt(ctx context.Context, id string) error {
	_, err := s.Conn.Exec(ctx, "update account set last_used_at=now() where id=$1", id)
	return err
}
