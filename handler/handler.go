package handler

import (
	"log"
	"net/http"

	"github.com/BTBurke/twilio-forwarder/config"
	"github.com/BTBurke/twiml"
)

// CallRequest will return XML to connect to the forwarding number
func CallRequest(cfg config.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var cr twiml.VoiceRequest
		if err := twiml.Bind(&cr, r); err != nil {
			log.Printf("%v", err)
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
			if err := res.Add(d); err != nil {
				http.Error(w, http.StatusText(502), 502)
				return
			}
			if err := res.Write(w); err != nil {
				http.Error(w, http.StatusText(502), 502)
				return
			}
			w.Header().Set("Content-Type", "application/xml")
			w.WriteHeader(200)
			return
		default:
			if err := res.Add(twiml.Hangup{}); err != nil {
				http.Error(w, http.StatusText(502), 502)
				return
			}
			if err := res.Write(w); err != nil {
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
func DialAction(cfg config.Config) func(http.ResponseWriter, *http.Request) {
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
			s := twiml.Say{
				Voice: "woman",
				Text:  cfg.VoicemailScript,
			}
			rec := twiml.Record{
				TranscribeCallback: "/voicemail",
			}
			if err := res.Add(s, rec); err != nil {
				http.Error(w, http.StatusText(502), 502)
				return
			}
			if err := res.Write(w); err != nil {
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
