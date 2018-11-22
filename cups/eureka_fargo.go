package cups

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/hudl/fargo"
)

func init() {
	gautocloud.RegisterConnector(NewEurekaFargoConnector())
}

type EurekaFargoSchema struct {
	URI string `cloud:"uri"`
}

type EurekaFargoConnector struct{}

func NewEurekaFargoConnector() connectors.Connector {
	return &EurekaFargoConnector{}
}

func (c EurekaFargoConnector) Id() string {
	return "cups:eureka-fargo"
}
func (c EurekaFargoConnector) Name() string {
	return ".*eureka.*"
}
func (c EurekaFargoConnector) Tags() []string {
	return []string{"eureka.*"}
}
func (c EurekaFargoConnector) Load(schema interface{}) (interface{}, error) {
	fSchema := schema.(EurekaFargoSchema)

	client := fargo.NewConn(fSchema.URI)
	return client, nil
}

func (c EurekaFargoConnector) Schema() interface{} {
	return EurekaFargoSchema{}
}
