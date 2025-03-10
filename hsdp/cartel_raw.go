package hsdp

import (
	"fmt"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/dip-software/go-dip-api/cartel"
)

func init() {
	gautocloud.RegisterConnector(NewCartelRawConnector())
}

type CartelConfig cartel.Config

type CartelRawConnector struct {
}

func (c CartelRawConnector) Id() string {
	return "hsdp:cartel-raw"
}

func (c CartelRawConnector) Name() string {
	return "*.cartel.*"
}

func (c CartelRawConnector) Tags() []string {
	return []string{"cartel"}
}

func (c CartelRawConnector) Load(schema interface{}) (interface{}, error) {
	fSchema, ok := schema.(cartel.Config)
	if !ok {
		return nil, fmt.Errorf("no cartel.Config detected")
	}
	if fSchema.Token == "" || fSchema.Secret == "" || fSchema.Host == "" {
		return nil, fmt.Errorf("invalid Cartel schema")
	}
	return CartelConfig(fSchema), nil
}

func (c CartelRawConnector) Schema() interface{} {
	return cartel.Config{}
}

func NewCartelRawConnector() connectors.Connector {
	return &CartelRawConnector{}
}
