package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/BTBurke/twilio-forwarder/config"
	"github.com/BTBurke/twilio-forwarder/handler"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

var cfg config.Config

func init() {
	cfg = config.Config{
		MailgunAPIKey:     os.Getenv("MAILGUN_API_KEY"),
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
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CloseNotify)
	r.Use(middleware.Timeout(10 * time.Second))

	r.Post("/call/", handler.CallRequest(cfg))
	r.Post("/call/action/", handler.DialAction(cfg))
	// r.Route("/call", func(r chi.Router) {
	// 	r.Post("/", CallRequest)
	// 	//r.Post("/action", CallActionCallback)
	// 	//r.Post("/voicemail", VoicemailActionCallback)
	// })
	http.ListenAndServe(":8080", r)
}
