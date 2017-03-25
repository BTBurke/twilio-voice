package twiml

// VoiceRequest represents the standard request format for callbacks.  This struct is
// embedded in other callback requests that return this common data format.
type VoiceRequest struct {
	CallSid       string
	AccountSid    string
	From          string
	To            string
	CallStatus    string
	APIVersion    string `schema:"ApiVersion"`
	Direction     string
	ForwardedFrom string
	CallerName    string
	FromCity      string
	FromState     string
	FromZip       string
	FromCountry   string
	ToCity        string
	ToState       string
	ToZip         string
	ToCountry     string
}

// DialActionRequest represents a request as a result of declaring an `action` URL on the Dial verb
type DialActionRequest struct {
	VoiceRequest
	DialCallStatus        string
	DialCallSid           string
	DialCallDuration      int
	RecordingURL          string `schema:"RecordingUrl"`
	QueueSid              string
	DequeueResult         string
	DequeuedCallSid       string
	DequeuedCallQueueTime int
	DequeuedCallDuration  int
}

// RecordActionRequest represents a request as a result of declaring an `action`
// URL on a Record verb
type RecordActionRequest struct {
	VoiceRequest
	RecordingURL      string `schema:"RecordingUrl"`
	RecordingDuration int
	Digits            string
}

// RecordingStatusCallbackRequest represents a request as a result of declaring
// a `recordingStatusCallback` on a Record verb
type RecordingStatusCallbackRequest struct {
	AccountSid        string
	CallSid           string
	RecordingSid      string
	RecordingURL      string `schema:"RecordingUrl"`
	RecordingStatus   string
	RecordingDuration int
	RecordingChannels int
	RecordingSource   string
}

// TranscribeCallbackRequest represents a request as a result of declaring
// a `transcribeCallback` on a Record verb
type TranscribeCallbackRequest struct {
	TranscriptionSid    string
	TranscriptionText   string
	TranscriptionStatus string
	TranscriptionURL    string `schema:"TranscriptionUrl"`
	RecordingSid        string
	RecordingURL        string `schema:"RecordingUrl"`
	CallSid             string
	AccountSid          string
	From                string
	To                  string
	CallStatus          string
	APIVersion          string `schema:"ApiVersion"`
	Direction           string
	ForwardedFrom       string
}
