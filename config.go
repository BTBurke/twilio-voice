package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

type Config struct {
	MailgunPublicKey   string
	MailgunSecretKey   string
	MailgunDomain      string
	NotificationEmail  string
	ForwardingNumber   string
	VoicemailScript    string
	VoicemailFile      string
	EnableCustomPrompt bool
	ServeDirectory     string
	VoiceFileName      string
}

func (cfg *Config) Validate() (errors []error) {
	workingDir, err := os.Getwd()
	if err != nil {
		workingDir = ""
	}
	fullVoicemailPath := path.Join(workingDir, cfg.VoicemailFile)

	if (len(cfg.MailgunPublicKey) == 0) || (len(cfg.MailgunSecretKey) == 0) || (len(cfg.MailgunDomain) == 0) {
		errors = append(errors, fmt.Errorf("set MAILGUN_PUBLIC_KEY, MAILGUN_SECRET_KEY, MAILGUN_DOMAIN environment variables to receive voicemail notifications"))
	}
	if len(cfg.NotificationEmail) == 0 {
		errors = append(errors, fmt.Errorf("set NOTIFICATION_EMAIL environment variable to receive voicemail notifications"))
	}
	if len(cfg.ForwardingNumber) == 0 {
		errors = append(errors, fmt.Errorf("set FORWARDING_NUMBER environment variable to connect your incoming calls to your phone"))
	}
	// If no voicemail file is accessible and no script is set, falls back to generic voicemail prompt
	if stat, err := os.Stat(fullVoicemailPath); os.IsNotExist(err) || stat.IsDir() {
		log.Printf("Voicemail file not found, falling back to voice prompt")
		cfg.VoicemailFile = ""
		if len(cfg.VoicemailScript) == 0 {
			cfg.VoicemailScript = "Please leave a message"
		}
	}
	if len(cfg.VoicemailFile) > 0 {
		cfg.EnableCustomPrompt = true
		cfg.ServeDirectory, cfg.VoiceFileName = path.Split(fullVoicemailPath)
	}
	return
}
