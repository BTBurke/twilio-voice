package twiml

import (
	"net/http"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

// Bind will marshal a callback request from the Twilio API
// into the cbRequest struct provided
func Bind(cbRequest interface{}, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	decoder.IgnoreUnknownKeys(true)
	if err := decoder.Decode(cbRequest, r.PostForm); err != nil {
		return err
	}
	return nil
}
