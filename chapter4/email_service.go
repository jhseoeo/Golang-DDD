package chapter4

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type EmailSender interface {
	SendEmail(ctx context.Context, to string, title string, body string) error
}

const emailUrl = "https://maindrillapp.com/api/1.0/messages/send\""

type MailChimp struct {
	apiKey     string
	from       string
	httpClient http.Client
}

type MailChimpReqBody struct {
	Key     string `json:"key"`
	Message struct {
		FromEmail string `json:"from_email"`
		Subject   string `json:"subject"`
		Text      string `json:"text"`
		To        []struct {
			Email string `json:"email"`
			Type  string `json:"type"`
		} `json:"to"`
	} `json:"message"`
}

func NewMailChimp(apiKey string, from string, httpClient http.Client) *MailChimp {
	return &MailChimp{
		apiKey:     apiKey,
		from:       from,
		httpClient: httpClient,
	}
}
func (m MailChimp) SendEmail(ctx context.Context, to string, title string, body string) error {
	mailBody := MailChimpReqBody{
		Key: m.apiKey,
		Message: struct {
			FromEmail string `json:"from_email"`
			Subject   string `json:"subject"`
			Text      string `json:"text"`
			To        []struct {
				Email string `json:"email"`
				Type  string `json:"type"`
			} `json:"to"`
		}{
			FromEmail: m.from,
			Subject:   title,
			Text:      body,
			To: []struct {
				Email string `json:"email"`
				Type  string `json:"type"`
			}{{Email: to, Type: "to"}},
		},
	}

	payload, err := json.Marshal(mailBody)
	if err != nil {
		return fmt.Errorf("failed to marshall body: %w", err)
	}
	req, err := http.NewRequest(http.MethodPost, emailUrl, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	_, err = m.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
