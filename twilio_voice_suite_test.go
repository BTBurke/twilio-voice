package main

import (
	"log"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTwilioVoice(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "twilio-voice suite")
}

var _ = BeforeSuite(func() {
	log.SetOutput(GinkgoWriter)
})
