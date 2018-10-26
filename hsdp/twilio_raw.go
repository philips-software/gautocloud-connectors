package hsdp

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
)

func init() {
	gautocloud.RegisterConnector(NewTwilioRawConnector())
}

type TwilioSchema struct {
	TwilioAuthToken string `cloud:"twilio_auth_token"`
	TwilioSID       string `cloud:"twilio_sid"`
}

type TwilioSubAccount struct {
	SID       string
	AuthToken string
}

type TwilioRawConnector struct{}

func NewTwilioRawConnector() connectors.Connector {
	return &TwilioRawConnector{}
}

func (c TwilioRawConnector) Id() string {
	return "hsdp:twilio-raw"
}
func (c TwilioRawConnector) Name() string {
	return ".*twilio.*"
}
func (c TwilioRawConnector) Tags() []string {
	return []string{"twilio.*"}
}
func (c TwilioRawConnector) Load(schema interface{}) (interface{}, error) {
	fSchema := schema.(TwilioSchema)
	return TwilioSubAccount{
		SID:       fSchema.TwilioSID,
		AuthToken: fSchema.TwilioAuthToken,
	}, nil
}

func (c TwilioRawConnector) Schema() interface{} {
	return TwilioSchema{}
}
