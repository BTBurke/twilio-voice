package config

import (
	"fmt"
	"os"
)

type Config struct {
	MailgunAPIKey     string
	NotificationEmail string
	ForwardingNumber  string
	VoicemailScript   string
	VoicemailFile     string
}

func (cfg *Config) Validate() (errors []error) {
	if len(cfg.MailgunAPIKey) == 0 {
		errors = append(errors, fmt.Errorf("set MAILGUN_API_KEY environment variable to receive voicemail notifications"))
	}
	if len(cfg.NotificationEmail) == 0 {
		errors = append(errors, fmt.Errorf("set NOTIFICATION_EMAIL environment variable to receive voicemail notifications"))
	}
	if len(cfg.ForwardingNumber) == 0 {
		errors = append(errors, fmt.Errorf("set FORWARDING_NUMBER environment variable to connect your incoming calls to your phone"))
	}
	if _, err := os.Stat(cfg.VoicemailFile); os.IsNotExist(err) && len(cfg.VoicemailScript) == 0 {
		errors = append(errors, fmt.Errorf("set VOICEMAIL_SCRIPT or VOICEMAIL_FILE environment variables to receive voicemail notifications"))
	}
	return
}
