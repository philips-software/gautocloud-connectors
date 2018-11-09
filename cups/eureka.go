package cups

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/loafoe/go-eureka-client/eureka"
)

func init() {
	gautocloud.RegisterConnector(NewEurekaConnector())
}

type EurekaSchema struct {
	URI string `cloud:"uri"`
}

type EurekaConnector struct{}

func NewEurekaConnector() connectors.Connector {
	return &EurekaConnector{}
}

func (c EurekaConnector) Id() string {
	return "cups:eureka"
}
func (c EurekaConnector) Name() string {
	return ".*eureka.*"
}
func (c EurekaConnector) Tags() []string {
	return []string{"eureka.*"}
}
func (c EurekaConnector) Load(schema interface{}) (interface{}, error) {
	fSchema := schema.(EurekaSchema)

	client := eureka.NewClient([]string{
		fSchema.URI,
	})
	return client, nil
}

func (c EurekaConnector) Schema() interface{} {
	return EurekaSchema{}
}
