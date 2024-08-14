package email

import (
	"errors"
	"io"
	"mime"
	"net/mail"
	"strings"
)

type Message struct {
	Header      mail.Header
	MessageId   string
	FromName    string
	FromAddress string
	Subject     string
	PlainBody   string
	HtmlBody    string
}

func Parse(email io.Reader) (Message, error) {
	wd := new(mime.WordDecoder)

	// parse message
	m, err := mail.ReadMessage(email)
	if err != nil {
		return Message{}, err
	}

	// message id
	messageId, err := mail.ParseAddress(m.Header.Get("message-id"))
	if err != nil {
		return Message{}, err
	}

	// subject
	subject, err := wd.DecodeHeader(m.Header.Get("subject"))
	if err != nil {
		return Message{}, err
	}

	// parse from
	from, err := mail.ParseAddress(m.Header.Get("from"))
	if err != nil {
		return Message{}, err
	}

	fromName := from.Name
	fromAddress := strings.ToLower(from.Address)
	if !ValidateAddress(fromAddress) {
		return Message{}, errors.New("invalid from address")
	}

	// parse body
	plainBody, htmlBody, err := parseBody(
		m.Header.Get("content-type"),
		m.Header.Get("content-transfer-encoding"),
		m.Body,
	)
	if err != nil {
		return Message{}, err
	}

	// return
	return Message{
		Header:      m.Header,
		MessageId:   messageId.Address,
		FromName:    fromName,
		FromAddress: fromAddress,
		Subject:     subject,
		PlainBody:   plainBody,
		HtmlBody:    htmlBody,
	}, nil
}
