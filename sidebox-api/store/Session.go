package store

import (
	"context"
	"time"

	"alles/boxes/modules/random"

	"github.com/google/uuid"
)

type Session struct {
	Id        string
	AccountId string
	Token     string
	Address   string
	UserAgent string
	CreatedAt time.Time
}

func (s Store) SessionGetByToken(ctx context.Context, token string) (Session, error) {
	var session Session
	err := s.Conn.QueryRow(ctx, "select id, account_id, token, address, user_agent, created_at from session where token=$1", token).
		Scan(&session.Id, &session.AccountId, &session.Token, &session.Address, &session.UserAgent, &session.CreatedAt)
	return session, err
}

func (s Store) SessionCreate(ctx context.Context, accountId string, address string, userAgent string) (Session, error) {
	var session Session
	err := s.Conn.QueryRow(ctx, "insert into session (id, account_id, token, address, user_agent) "+
		"values ($1, $2, $3, $4, $5) "+
		"returning id, account_id, token, address, user_agent, created_at",
		uuid.New(), accountId, random.String(32), address, userAgent).
		Scan(&session.Id, &session.AccountId, &session.Token, &session.Address, &session.UserAgent, &session.CreatedAt)
	return session, err
}
