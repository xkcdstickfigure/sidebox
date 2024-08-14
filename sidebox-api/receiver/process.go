package receiver

import (
	"context"
	"io"
	"strings"

	"alles/boxes/modules/email"
	"alles/boxes/store"

	"github.com/google/uuid"
)

func process(ctx context.Context, db store.Store, input io.Reader) error {
	// get email data
	inputBytes, err := io.ReadAll(input)
	if err != nil {
		return err
	}
	inputString := strings.Join(strings.Split(string(inputBytes), "\n")[1:], "\n")

	// parse email message
	message, err := email.Parse(strings.NewReader(inputString))
	if err != nil {
		return err
	}

	// determine body
	html := message.HtmlBody != ""
	var body *string
	if html {
		body = &message.HtmlBody
	} else {
		body = &message.PlainBody
	}

	// get inbox
	to := message.Header.Get("Delivered-To")
	toUsername := strings.ToLower(strings.Split(to, "@")[0])
	inbox, err := db.InboxGetByCode(ctx, toUsername)
	if err != nil {
		return err
	}

	// create message
	id := uuid.New().String()
	err = db.MessageCreate(ctx, store.Message{
		Id:          id,
		InboxId:     inbox.Id,
		MessageId:   message.MessageId,
		FromName:    message.FromName,
		FromAddress: message.FromAddress,
		Subject:     message.Subject,
		Body:        *body,
		Html:        html,
	})
	if err != nil {
		return err
	}

	// mark inbox as read
	db.InboxSetUnread(ctx, inbox.Id, true)

	return nil
}
