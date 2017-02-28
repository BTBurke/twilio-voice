package main

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/BTBurke/twiml"
	"gopkg.in/mailgun/mailgun-go.v1"
)

func Send(cfg Config, tcb twiml.TranscribeCallbackRequest) error {
	mg := mailgun.NewMailgun(cfg.MailgunDomain, cfg.MailgunSecretKey, cfg.MailgunPublicKey)

	emailTemplate, err := Asset("templates/voicemail.html")
	if err != nil {
		return err
	}
	tmpl, err := template.New("email").Parse(string(emailTemplate))
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, tcb); err != nil {
		return err
	}

	message := mailgun.NewMessage(
		fmt.Sprintf("voicemail@%s", cfg.MailgunDomain),
		fmt.Sprintf("New voicemail from %s", tcb.From),
		fmt.Sprintf("You have received a new voicemail from %s\n\nTranscript:\n%s\n\nVoicemail Link: %s\n", tcb.From, tcb.TranscriptionText, tcb.RecordingURL),
		cfg.NotificationEmail,
	)
	message.SetHtml(buf.String())
	if _, _, err := mg.Send(message); err != nil {
		return err
	}
	return nil
}
