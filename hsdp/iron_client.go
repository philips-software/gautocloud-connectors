package hsdp

import (
	"fmt"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/philips-software/go-hsdp-api/iron"
)

func init() {
	gautocloud.RegisterConnector(NewIronClientConnector())
}

type IronClient struct {
	iron.Config
	*iron.Client
}

type IronClientConnector struct {
	wrapRawConn connectors.Connector
}

func (i IronClientConnector) Id() string {
	return "hsdp:iron-client"
}

func (i IronClientConnector) Name() string {
	return i.wrapRawConn.Name()
}

func (i IronClientConnector) Tags() []string {
	return i.wrapRawConn.Tags()
}

func (i IronClientConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := i.wrapRawConn.Load(schema)
	if err != nil {
		return nil, err
	}
	fSchema := schema.(IronPlan)
	if fSchema.ProjectID == "" {
		return nil, fmt.Errorf("invalid ProjectID")
	}
	if fSchema.Token == "" {
		return nil, fmt.Errorf("missing token")
	}
	config := iron.Config(fSchema)
	client, err := iron.NewClient(&config)
	if err != nil {
		return nil, err
	}
	return &IronClient{
		Client: client,
		Config: config,
	}, nil
}

func (i IronClientConnector) Schema() interface{} {
	return i.wrapRawConn.Schema()
}

func NewIronClientConnector() connectors.Connector {
	return &IronClientConnector{
		wrapRawConn: NewIronRawConnector(),
	}
}
