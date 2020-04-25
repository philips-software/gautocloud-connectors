package hsdp

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/kevinburke/twilio-go"
)

func init() {
	gautocloud.RegisterConnector(NewTwilioClientConnector())
}

type TwilioClientConnector struct {
	wrapRawConn connectors.Connector
}

func NewTwilioClientConnector() connectors.Connector {
	return &TwilioClientConnector{
		wrapRawConn: NewTwilioRawConnector(),
	}
}

func (c TwilioClientConnector) Id() string {
	return "hsdp:twilio-client"
}
func (c TwilioClientConnector) Name() string {
	return ".*twilio.*"
}
func (c TwilioClientConnector) Tags() []string {
	return []string{"twilio.*"}
}
func (c TwilioClientConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := c.wrapRawConn.Load(schema)
	if err != nil {
		return nil, err
	}
	fSchema := schema.(TwilioSubAccount)
	client := twilio.NewClient(fSchema.SID, fSchema.AuthToken, nil)
	return client, nil
}

func (c TwilioClientConnector) Schema() interface{} {
	return c.wrapRawConn.Schema()
}
