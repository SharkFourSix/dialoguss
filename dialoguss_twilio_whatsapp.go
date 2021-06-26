package main

/**
 * Support for Twilio's WhatsApp APIs for Dialoguss
 */
import (
	"github.com/sfreiberg/gotwilio"
)

const DialogussTwilioMode DialogussModeType = 3

type twilioRequest struct{}

func (t *twilioRequest) Send(appSid string, twilio *gotwilio.Twilio) {
	from := ""
	to := ""
	body := ""

	twilio.SendWhatsApp(from, to, body, "", appSid)

}
