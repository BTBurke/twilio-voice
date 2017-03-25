package twiml

import (
	"encoding/xml"
	"fmt"
)

// Client TwiML
type Client struct {
	XMLName xml.Name `xml:"Client"`
	Method  string   `xml:"method,attr,omitempty"`
	URL     string   `xml:"URL,omitempty"`
	Name    string   `xml:",chardata"`
}

func (c *Client) Validate() error {
	ok := Validate(
		AllowedMethod(c.Method),
		Required(c.Name),
	)
	if !ok {
		return fmt.Errorf("Client markup failed validation")
	}
	return nil
}

func (c *Client) Type() string {
	return "Client"
}

// Conference TwiML
type Conference struct {
	XMLName                       xml.Name `xml:"Conference"`
	Muted                         bool     `xml:"muted,attr,omitempty"`
	Beep                          string   `xml:"beep,attr,omitempty"`
	StartConferenceOnEnter        bool     `xml:"startConferenceOnEnter,attr,omitempty"`
	EndConferenceOnExit           bool     `xml:"endConferenceOnExit,attr,omitempty"`
	WaitURL                       string   `xml:"waitUrl,attr,omitempty"`
	WaitMethod                    string   `xml:"waitMethod,attr,omitempty"`
	MaxParticipants               int      `xml:"maxParticipants,attr,omitempty"`
	Record                        string   `xml:"record,attr,omitempty"`
	Trim                          string   `xml:"trim,attr,omitempty"`
	StatusCallbackEvent           string   `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallback                string   `xml:"statusCallback,attr,omitempty"`
	StatusCallbackMethod          string   `xml:"statusCallbackMethod,attr,omitempty"`
	RecordingStatusCallback       string   `xml:"recordingStatusCallback,attr,omitempty"`
	RecordingStatusCallbackMethod string   `xml:"recordingStatusCallbackMethod,attr,omitempty"`
	EventCallbackURL              string   `xml:"eventCallbackUrl,attr,omitempty"`
}

func (c *Conference) Validate() error {
	ok := Validate(
		OneOfOpt(c.Beep, "true", "false", "onEnter", "onExit"),
		AllowedMethod(c.WaitMethod),
		OneOfOpt(c.Record, "do-not-record", "record-from-start"),
		OneOfOpt(c.Trim, "trim-silence", "do-not-trim"),
		OneOfOpt(c.StatusCallbackEvent, "start", "end", "join", "leave", "mute", "hold"),
		AllowedMethod(c.StatusCallbackMethod),
		AllowedMethod(c.RecordingStatusCallbackMethod),
	)
	if !ok {
		return fmt.Errorf("Conference markup failed validation")
	}
	return nil
}

func (c *Conference) Type() string {
	return "Conference"
}

// Dial TwiML
type Dial struct {
	XMLName      xml.Name `xml:"Dial"`
	Action       string   `xml:"action,attr,omitempty"`
	Method       string   `xml:"method,attr,omitempty"`
	Timeout      int      `xml:"timeout,attr,omitempty"`
	HangupOnStar bool     `xml:"hangupOnStar,attr,omitempty"`
	TimeLimit    int      `xml:"timeLimit,attr,omitempty"`
	CallerID     string   `xml:"callerId,attr,omitempty"`
	Record       bool     `xml:"record,attr,omitempty"`
	Number       string   `xml:",chardata"`
	Children     []Markup `xml:",omitempty"`
}

func (d *Dial) Validate() error {
	var errs []error
	for _, s := range d.Children {
		switch t := s.Type(); t {
		default:
			return fmt.Errorf("Not a valid verb under Dial: '%T'", s)
		case "Client", "Conference", "Number", "Queue", "Sip":
			if childErr := s.Validate(); childErr != nil {
				errs = append(errs, childErr)
			}
		}
	}

	ok := Validate(
		OneOfOpt(d.Method, "GET", "POST"),
		Required(d.Number),
	)
	if !ok {
		errs = append(errs, fmt.Errorf("Dial did not pass validation"))
	}

	if len(errs) > 0 {
		return ValidationError{errs}
	}
	return nil
}

// Add adds noun structs to a Dial response as children
func (d *Dial) Add(ml ...Markup) {
	for _, s := range ml {
		d.Children = append(d.Children, s)
	}
	return
}

func (d *Dial) Type() string {
	return "Dial"
}

// Enqueue TwiML
type Enqueue struct {
	XMLName       xml.Name `xml:"Enqueue"`
	Action        string   `xml:"action,attr,omitempty"`
	Method        string   `xml:"method,attr,omitempty"`
	WaitURL       string   `xml:"waitUrl,attr,omitempty"`
	WaitURLMethod string   `xml:"waiUrlMethod,attr,omitempty"`
	WorkflowSid   string   `xml:"workflowSid,attr,omitempty"`
	QueueName     string   `xml:",chardata"`
}

func (e *Enqueue) Validate() error {
	ok := Validate(
		AllowedMethod(e.Method),
		AllowedMethod(e.WaitURLMethod),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", e.Type())
	}
	return nil
}

func (e *Enqueue) Type() string {
	return "Enqueue"
}

// Hangup TwiML
type Hangup struct {
	XMLName xml.Name `xml:"Hangup"`
}

func (h *Hangup) Validate() error {
	return nil
}

func (h *Hangup) Type() string {
	return "Hangup"
}

// Leave TwiML
type Leave struct {
	XMLName xml.Name `xml:"Leave"`
}

func (l *Leave) Validate() error {
	return nil
}

func (l *Leave) Type() string {
	return "Leave"
}

// Sms TwiML sends an SMS message. Text is required.  See the Twilio docs
// for an explanation of the default values of to and from.
type Sms struct {
	XMLName        xml.Name `xml:"Message"`
	To             string   `xml:"to,attr,omitempty"`
	From           string   `xml:"from,attr,omitempty"`
	Action         string   `xml:"action,attr,omitempty"`
	Method         string   `xml:"method,attr,omitempty"`
	StatusCallback string   `xml:"statusCallback,attr,omitempty"`
	Text           string   `xml:",chardata"`
}

func (s *Sms) Validate() error {
	ok := Validate(
		AllowedMethod(s.Method),
		Required(s.Text),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", s.Type())
	}
	return nil
}

func (s *Sms) Type() string {
	return "Sms"
}

// Number TwiML
type Number struct {
	XMLName    xml.Name `xml:"Number"`
	SendDigits string   `xml:"sendDigits,attr,omitempty"`
	URL        string   `xml:"url,attr,omitempty"`
	Method     string   `xml:"method,attr,omitempty"`
	Number     string   `xml:",chardata"`
}

func (n *Number) Validate() error {
	ok := Validate(
		NumericOpt(n.SendDigits),
		AllowedMethod(n.Method),
		Required(n.Number),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", n.Type())
	}
	return nil
}

func (n *Number) Type() string {
	return "Number"
}

// Pause TwiML
type Pause struct {
	XMLName xml.Name `xml:"Pause"`
	Length  int      `xml:"length,attr,omitempty"`
}

func (p *Pause) Validate() error {
	return nil
}

func (p *Pause) Type() string {
	return "Pause"
}

// Play TwiML
type Play struct {
	XMLName xml.Name `xml:"Play"`
	Loop    int      `xml:"loop,attr,omitempty"`
	Digits  int      `xml:"digits,attr,omitempty"`
	URL     string   `xml:",chardata"`
}

func (p *Play) Validate() error {

	ok := Validate(
		Required(p.URL),
	)

	if !ok {
		return fmt.Errorf("%s markup failed validation", p.Type())
	}
	return nil
}

func (p *Play) Type() string {
	return "Play"
}

// Queue TwiML
type Queue struct {
	XMLName             xml.Name `xml:"Queue"`
	URL                 string   `xml:"url,attr,omitempty"`
	Method              string   `xml:"method,attr,omitempty"`
	ReservationSid      string   `xml:"reservationSid,attr,omitempty"`
	PostWorkActivitySid string   `xml:"postWorkActivitySid,attr,omitempty"`
	Name                string   `xml:",chardata"`
}

func (q *Queue) Validate() error {
	ok := Validate(
		AllowedMethod(q.Method),
		Required(q.Name),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", q.Type())
	}
	return nil
}

func (q *Queue) Type() string {
	return "Queue"
}

// Record TwiML
type Record struct {
	XMLName                       xml.Name `xml:"Record"`
	Action                        string   `xml:"action,attr,omitempty"`
	Method                        string   `xml:"method,attr,omitempty"`
	Timeout                       int      `xml:"timeout,attr,omitempty"`
	FinishOnKey                   string   `xml:"finishOnKey,attr,omitempty"`
	MaxLength                     int      `xml:"maxLength,attr,omitempty"`
	PlayBeep                      bool     `xml:"playBeep,attr,omitempty"`
	Trim                          string   `xml:"trim,attr,omitempty"`
	RecordingStatusCallback       string   `xml:"recordingStatusCallback,attr,omitempty"`
	RecordingStatusCallbackMethod string   `xml:"recordingStatusCallbackMethod,attr,omitempty"`
	Transcribe                    bool     `xml:"transcribe,attr,omitempty"`
	TranscribeCallback            string   `xml:"transcribeCallback,attr,omitempty"`
}

func (r *Record) Validate() error {
	ok := Validate(
		AllowedMethod(r.Method),
		OneOfOpt(r.Trim, TrimSilence, DoNotTrim),
		AllowedMethod(r.RecordingStatusCallbackMethod),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", r.Type())
	}
	return nil
}

func (r *Record) Type() string {
	return "Record"
}

// Redirect TwiML
type Redirect struct {
	XMLName xml.Name `xml:"Redirect"`
	Method  string   `xml:"method,attr,omitempty"`
	URL     string   `xml:",chardata"`
}

func (r *Redirect) Validate() error {
	ok := Validate(
		AllowedMethod(r.Method),
		Required(r.URL),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", r.Type())
	}
	return nil
}

func (r *Redirect) Type() string {
	return "Redirect"
}

// Reject TwiML
type Reject struct {
	XMLName xml.Name `xml:"Reject"`
	Reason  string   `xml:"reason,attr,omitempty"`
}

func (r *Reject) Validate() error {
	ok := Validate(
		OneOfOpt(r.Reason, "rejected", "busy"),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", r.Type())
	}
	return nil
}

func (r *Reject) Type() string {
	return "Reject"
}

// Say TwiML
type Say struct {
	XMLName  xml.Name `xml:"Say"`
	Voice    string   `xml:"voice,attr,omitempty"`
	Language string   `xml:"language,attr,omitempty"`
	Loop     int      `xml:"loop,attr,omitempty"`
	Text     string   `xml:",chardata"`
}

func (s *Say) Validate() error {
	ok := Validate(
		OneOfOpt(s.Voice, Man, Woman, Alice),
		AllowedLanguage(s.Voice, s.Language),
		Required(s.Text),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", s.Type())
	}
	return nil

}

func (s *Say) Type() string {
	return "Say"
}

// Sip TwiML
type Sip struct {
	XMLName              xml.Name `xml:"Sip"`
	Username             string   `xml:"username,attr,omitempty"`
	Password             string   `xml:"password,attr,omitempty"`
	URL                  string   `xml:"url,attr,omitempty"`
	Method               string   `xml:"method,attr,omitempty"`
	StatusCallbackEvent  string   `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallback       string   `xml:"statusCallback,attr,omitempty"`
	StatusCallbackMethod string   `xml:"statusCallbackMethod,attr,omitempty"`
	Address              string   `xml:",chardata"`
}

// TODO: Needs helpers to construct the SIP URL (specifying transport
// and headers) See https://www.twilio.com/docs/api/twiml/sip

func (s *Sip) Validate() error {
	//TODO: needs a custom validator type for statusCallbackEvent when set
	//because valid values can be concatenated
	ok := Validate(
		AllowedMethod(s.StatusCallbackMethod),
		AllowedCallbackEvent(s.StatusCallbackEvent),
		Required(s.Address),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", s.Type())
	}
	return nil
}

func (s *Sip) Type() string {
	return "Sip"
}

// Gather TwiML
type Gather struct {
	XMLName     xml.Name `xml:"Gather"`
	Action      string   `xml:"action,attr,omitempty"`
	Method      string   `xml:"method,attr,omitempty"`
	Timeout     int      `xml:"timeout,attr,omitempty"`
	FinishOnKey string   `xml:"finishOnKey,attr,omitempty"`
	NumDigits   int      `xml:"numDigits,attr,omitempty"`
	Children    []Markup `valid:"-"`
}

func (g *Gather) Validate() error {
	var errs []error

	for _, s := range g.Children {
		switch t := s.Type(); t {
		default:
			return fmt.Errorf("Not a valid verb as child of Gather: '%T'", s)
		case "Say", "Play", "Pause":
			if childErr := s.Validate(); childErr != nil {
				errs = append(errs, childErr)
			}
		}
	}
	ok := Validate(
		AllowedMethod(g.Method),
	)
	if !ok {
		errs = append(errs, fmt.Errorf("Gather failed validation"))
		return ValidationError{errs}
	}
	if len(errs) > 0 {
		return ValidationError{errs}
	}
	return nil
}

// Add collects digits a caller enter by pressing the keypad to an existing Gather verb.
// Valid nested verbs: Say, Pause, Play
func (g *Gather) Add(ml ...Markup) {
	for _, s := range ml {
		g.Children = append(g.Children, s)
	}
	return
}

func (g *Gather) Type() string {
	return "Gather"
}
