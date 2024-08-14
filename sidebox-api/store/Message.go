package store

import (
	"context"
	"time"
)

type Message struct {
	Id          string
	InboxId     string
	MessageId   string
	FromName    string
	FromAddress string
	Subject     string
	Body        string
	Html        bool
	Date        time.Time
}

type MessageWithoutBody struct {
	Id          string
	InboxId     string
	MessageId   string
	FromName    string
	FromAddress string
	Subject     string
	Html        bool
	Date        time.Time
}

func (s Store) MessageGet(ctx context.Context, id string) (Message, error) {
	var message Message
	err := s.Conn.QueryRow(ctx, "select id, inbox_id, message_id, from_name, from_address, subject, body, html, date from message where id=$1", id).
		Scan(&message.Id, &message.InboxId, &message.MessageId, &message.FromName, &message.FromAddress, &message.Subject, &message.Body, &message.Html, &message.Date)
	return message, err
}

func (s Store) MessageCreate(ctx context.Context, data Message) error {
	_, err := s.Conn.Exec(ctx, "insert into message (id, inbox_id, message_id, from_name, from_address, subject, body, html) values ($1, $2, $3, $4, $5, $6, $7, $8)",
		data.Id, data.InboxId, data.MessageId, data.FromName, data.FromAddress, data.Subject, data.Body, data.Html)
	return err
}

func (s Store) MessageList(ctx context.Context, inboxId string) ([]MessageWithoutBody, error) {
	messages := []MessageWithoutBody{}
	rows, err := s.Conn.Query(ctx, "select id, inbox_id, message_id, from_name, from_address, subject, html, date from message "+
		"where inbox_id=$1 "+
		"order by date desc",
		inboxId)
	if err != nil {
		return messages, err
	}
	defer rows.Close()

	for rows.Next() {
		var message MessageWithoutBody
		err = rows.Scan(&message.Id, &message.InboxId, &message.MessageId, &message.FromName, &message.FromAddress, &message.Subject, &message.Html, &message.Date)
		if err != nil {
			return messages, err
		}
		messages = append(messages, message)
	}
	return messages, err
}
