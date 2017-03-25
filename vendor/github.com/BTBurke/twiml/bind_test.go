package twiml

import (
	"bytes"
	"net/http"
	"net/url"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func makeRequest(v map[string]string) *http.Request {
	apiURL := "https://test.com"
	data := url.Values{}
	for key, value := range v {
		data.Set(key, value)
	}

	d := data.Encode()
	r, _ := http.NewRequest("POST", apiURL, bytes.NewBufferString(d))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(d)))
	return r
}

var _ = Describe("Callback binding", func() {
	It("can bind a callback request", func() {
		values := map[string]string{
			"CallSid": "testsid",
			"From":    "+19999999999",
			"To":      "+19991111111",
		}
		exp := VoiceRequest{
			CallSid: "testsid",
			From:    "+19999999999",
			To:      "+19991111111",
		}
		r := makeRequest(values)
		var vr VoiceRequest
		err := Bind(&vr, r)
		Expect(err).ToNot(HaveOccurred())
		Expect(vr).To(Equal(exp))
	})

	It("can bind requests for non-string types", func() {
		values := map[string]string{
			"RecordingUrl":      "https://test.api",
			"RecordingDuration": "10",
		}
		exp := RecordActionRequest{
			RecordingURL:      "https://test.api",
			RecordingDuration: 10,
		}
		r := makeRequest(values)
		var vr RecordActionRequest
		err := Bind(&vr, r)
		Expect(err).ToNot(HaveOccurred())
		Expect(vr).To(Equal(exp))
	})
})
