package hsdp

import (
	"fmt"

	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/dip-software/go-dip-api/iron"
)

func init() {
	gautocloud.RegisterConnector(NewIronRawConnector())
}

type IronPlan iron.Config

type IronRawConnector struct{}

func (i IronRawConnector) Id() string {
	return "hsdp:iron-raw"
}

func (i IronRawConnector) Name() string {
	return ".*iron.*"
}

func (i IronRawConnector) Tags() []string {
	return []string{"iron.*", "hsdp-iron", "iron.io"}
}

func (i IronRawConnector) Load(schema interface{}) (interface{}, error) {
	fSchema, ok := schema.(iron.Config)
	if !ok {
		return nil, fmt.Errorf("no iron.Config detected")
	}
	if fSchema.ProjectID == "" {
		return nil, fmt.Errorf("invalid Iron schema. Missing ProjectID")
	}
	return IronPlan(fSchema), nil
}

func (i IronRawConnector) Schema() interface{} {
	return iron.Config{}
}

func NewIronRawConnector() connectors.Connector {
	return &IronRawConnector{}
}
