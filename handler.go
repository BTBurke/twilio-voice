package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BTBurke/twiml"
)

// CallRequest will return XML to connect to the forwarding number
func CallRequest(cfg Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var cr twiml.VoiceRequest
		if err := twiml.Bind(&cr, r); err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		res := twiml.NewResponse()

		switch status := cr.CallStatus; status {
		case twiml.InProgress:
			w.WriteHeader(200)
			return
		case twiml.Ringing, twiml.Queued:
			d := twiml.Dial{
				Number:   cfg.ForwardingNumber,
				Action:   "action/",
				Timeout:  15,
				CallerID: cr.To,
			}
			res.Add(&d)
			b, err := res.Encode()
			if err != nil {
				http.Error(w, http.StatusText(502), 502)
				return
			}
			if _, err := w.Write(b); err != nil {
				http.Error(w, http.StatusText(502), 502)
				return
			}
			w.Header().Set("Content-Type", "application/xml")
			w.WriteHeader(200)
			return
		default:
			res.Add(&twiml.Hangup{})
			b, err := res.Encode()
			if err != nil {
				http.Error(w, http.StatusText(502), 502)
				return
			}
			if _, err := w.Write(b); err != nil {
				http.Error(w, http.StatusText(502), 502)
				return
			}
			w.Header().Set("Content-Type", "application/xml")
			w.WriteHeader(200)
			return
		}
	}
}

// DialAction will forward to voicemail if the call is not connected
func DialAction(cfg Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var ca twiml.DialActionRequest
		if err := twiml.Bind(&ca, r); err != nil {
			log.Printf("%v", err)
			http.Error(w, http.StatusText(400), 400)
			return
		}
		res := twiml.NewResponse()
		switch status := ca.DialCallStatus; status {
		case twiml.NoAnswer, twiml.Failed, twiml.Busy:
			if cfg.EnableCustomPrompt {
				p := twiml.Play{URL: fmt.Sprintf("/prompt/%s", cfg.VoiceFileName)}
				res.Add(&p)
			} else {
				s := twiml.Say{
					Voice: "woman",
					Text:  cfg.VoicemailScript,
				}
				res.Add(&s)
			}

			rec := twiml.Record{
				TranscribeCallback: "/voicemail",
			}
			res.Add(&rec)

			b, err := res.Encode()
			if err != nil {
				http.Error(w, http.StatusText(502), 502)
				return
			}
			if _, err := w.Write(b); err != nil {
				http.Error(w, http.StatusText(502), 502)
				return
			}
			w.Header().Set("Content-Type", "application/xml")
			w.WriteHeader(200)
			return
		default:
			w.WriteHeader(200)
			return

		}
	}
}

// Voicemail handles the TranscriptionCallback which lets you know that transcription is done and the
// voicemail is available.  If Mailgun is set, it will email a copy of the transcription text and
// a link to the voicemail to your email address
func Voicemail(cfg Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var tcb twiml.TranscribeCallbackRequest
		if err := twiml.Bind(&tcb, r); err != nil {
			log.Printf("%v", err)
			http.Error(w, http.StatusText(400), 400)
			return
		}
		log.Printf("Call from: %s\n\nTranscription follows:\n%s\n\nVoicemail Link: %s\n", tcb.From, tcb.TranscriptionText, tcb.RecordingURL)
		if err := Send(cfg, tcb); err != nil {
			log.Printf("Unable to send notification email due to error: %s\n\nVoicemail available at: %s", err, tcb.RecordingURL)
		}
		w.WriteHeader(200)
	}
}

// Status receives in-progress status events.  It is outside the mail control loop.  In this case,
// acknowledging the status to continue the call is the right thing to do.
func Status(cfg Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}
}
