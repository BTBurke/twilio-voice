package twiml

import (
	"encoding/xml"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTwiml(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TwiML Suite")
}

func x(s string, n int) string {
	return strings.Join([]string{strings.Repeat(" ", n), s, "\n"}, "")
}

func buildResponse(s ...string) string {
	t := []string{xml.Header, "<Response>\n"}
	for _, verb := range s {
		t = append(t, verb)
	}
	t = append(t, "</Response>")
	return strings.Join(t, "")
}

var _ = Describe("TwiML responses", func() {
	It("to error on an empty response", func() {
		r := NewResponse()
		_, err := r.String()
		Expect(err).To(HaveOccurred())
	})

	It("can encode a basic verb in XML", func() {
		r := NewResponse()
		d := Dial{
			Action: "https://testurl.com",
			Number: "415-999-9999",
		}
		exp := buildResponse(
			x("<Dial action=\"https://testurl.com\">415-999-9999</Dial>", 2),
		)
		r.Add(&d)
		resp, err := r.String()
		Expect(err).ToNot(HaveOccurred())
		Expect(resp).To(Equal(exp))
	})

	It("can encode nested verbs and nouns", func() {
		r := NewResponse()
		d := Dial{
			Number: "415-999-9999",
		}
		c := Client{
			Name: "test",
		}
		exp := buildResponse(
			x("<Dial>415-999-9999", 2),
			x("<Client>test</Client>", 4),
			x("</Dial>", 2),
		)

		d.Add(&c)
		r.Add(&d)
		s, err := r.String()
		Expect(err).ToNot(HaveOccurred())
		Expect(s).To(Equal(exp))
	})

	It("will validate markup and prevent encoding on errors", func() {
		d := Dial{Method: "TEST"}
		err := d.Validate()
		Expect(err).To(HaveOccurred())
	})
})
