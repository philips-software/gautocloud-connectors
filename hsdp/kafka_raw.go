package hsdp

import (
	"fmt"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	_ "github.com/cloudfoundry-community/gautocloud/connectors"
	"strings"
)

type KafkaSchema struct {
	Username  string   `json:"username,omitempty"`
	Password  string   `json:"password,omitempty"`
	Hostname  string   `json:"hostname,omitempty"`
	Port      int      `json:"port,omitempty"`
	Hostnames []string `json:"hostnames"`
	URI       string   `json:"uri"`
}

type KafkaCredentials KafkaSchema

type KafkaRawConnector struct{}

func init() {
	gautocloud.RegisterConnector(NewKafkaRawConnector())
}

func (r KafkaRawConnector) Id() string {
	return "hsdp:kafka-raw"
}

func (r KafkaRawConnector) Name() string {
	return ".*kafka.*"
}

func (r KafkaRawConnector) Tags() []string {
	return []string{"Kafka", "kafka"}
}

func (r KafkaRawConnector) Load(schema interface{}) (interface{}, error) {
	fSchema, ok := schema.(KafkaSchema)
	if !ok {
		return nil, fmt.Errorf("no KafkaSchema detected")
	}
	if len(fSchema.Hostnames) == 0 || !strings.Contains(fSchema.URI, "kafka://") {
		return nil, fmt.Errorf("no hostnames or invalid URI")
	}
	return KafkaCredentials(fSchema), nil
}

func (r KafkaRawConnector) Schema() interface{} {
	return KafkaSchema{}
}

func NewKafkaRawConnector() connectors.Connector {
	return &KafkaRawConnector{}
}
