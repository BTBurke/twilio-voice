package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

var cfg Config

func init() {
	cfg = Config{
		MailgunPublicKey:  os.Getenv("MAILGUN_PUBLIC_KEY"),
		MailgunSecretKey:  os.Getenv("MAILGUN_SECRET_KEY"),
		MailgunDomain:     os.Getenv("MAILGUN_DOMAIN"),
		NotificationEmail: os.Getenv("NOTIFICATION_EMAIL"),
		ForwardingNumber:  os.Getenv("FORWARDING_NUMBER"),
		VoicemailScript:   os.Getenv("VOICEMAIL_SCRIPT"),
		VoicemailFile:     os.Getenv("VOICEMAIL_FILE"),
	}
	if errs := cfg.Validate(); len(errs) > 0 {
		log.Fatalf("%v", errs)
	}
}

func main() {
	log.Printf("Forwarding calls to %s\n", cfg.ForwardingNumber)
	log.Printf("Voicemail notifications will be sent to %s\n", cfg.NotificationEmail)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CloseNotify)
	r.Use(middleware.Timeout(10 * time.Second))

	r.Post("/call/", CallRequest(cfg))
	r.Post("/call/action/", DialAction(cfg))
	r.Post("/voicemail", Voicemail(cfg))
	r.Post("/status", Status(cfg))
	if cfg.EnableCustomPrompt {
		log.Printf("Serving custom voicemail prompt from %s\n", cfg.VoicemailFile)
		r.FileServer("/prompt", http.Dir(cfg.ServeDirectory))
	}

	log.Println("Listening on 127.0.0.1:8080")
	http.ListenAndServe(":8080", r)
}
