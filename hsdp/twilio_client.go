package hsdp

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/kevinburke/twilio-go"
)

func init() {
	gautocloud.RegisterConnector(NewTwilioClientConnector())
}

type TwilioCientConnector struct {
	wrapRawConn connectors.Connector
}

func NewTwilioClientConnector() connectors.Connector {
	return &TwilioCientConnector{
		wrapRawConn: NewTwilioRawConnector(),
	}
}

func (c TwilioCientConnector) Id() string {
	return "hsdp:twilio-client"
}
func (c TwilioCientConnector) Name() string {
	return ".*twilio.*"
}
func (c TwilioCientConnector) Tags() []string {
	return []string{"twilio.*"}
}
func (c TwilioCientConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := c.wrapRawConn.Load(schema)
	if err != nil {
		return nil, err
	}
	fSchema := schema.(TwilioSubAccount)
	client := twilio.NewClient(fSchema.SID, fSchema.AuthToken, nil)
	return client, nil
}

func (c TwilioCientConnector) Schema() interface{} {
	return c.wrapRawConn.Schema()
}
