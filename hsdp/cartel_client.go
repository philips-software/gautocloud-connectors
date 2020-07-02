package hsdp

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/philips-software/go-hsdp-api/cartel"
)

func init() {
	gautocloud.RegisterConnector(NewCartelClientConnector())
}

type CartelClient struct {
	cartel.Config
	*cartel.Client
}

type CartelClientConnector struct {
	wrapRawConn connectors.Connector
}

func (c CartelClientConnector) Id() string {
	return "hsdp:cartel-client"
}

func (c CartelClientConnector) Name() string {
	return c.wrapRawConn.Name()
}

func (c CartelClientConnector) Tags() []string {
	return c.wrapRawConn.Tags()
}

func (c CartelClientConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := c.wrapRawConn.Load(schema)
	if err != nil {
		return nil, err
	}
	fSchema := schema.(CartelConfig)
	config := cartel.Config(fSchema)
	client, err := cartel.NewClient(nil, &config)
	if err != nil {
		return nil, err
	}
	return &CartelClient{
		Client: client,
		Config: config,
	}, nil
}

func (c CartelClientConnector) Schema() interface{} {
	return c.wrapRawConn.Schema()
}

func NewCartelClientConnector() connectors.Connector {
	return &CartelClientConnector{
		wrapRawConn: NewCartelRawConnector(),
	}
}
