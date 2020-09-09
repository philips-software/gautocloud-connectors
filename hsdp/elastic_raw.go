package hsdp

import (
	"fmt"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
)

type ElasticSchema struct {
	Hostname    string `json:"hostname"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Port        int    `json:"port"`
	SecretToken string `json:"secret_token"`
	ServerUrls  string `json:"server_urls"`
	URI         string `json:"uri"`
	Username    string `json:"username"`
}

type ElasticCredentials ElasticSchema

func (e ElasticCredentials) Server() string {
	return fmt.Sprintf("https://%s:%d", e.Hostname, e.Port)
}

type ElasticRawConnector struct{}

func (r ElasticRawConnector) Id() string {
	return "hsdp:elastic-raw"
}

func (r ElasticRawConnector) Name() string {
	return ".*elastic.*"
}

func (r ElasticRawConnector) Tags() []string {
	return []string{"Elasticsearch", "elastic"}
}

func (r ElasticRawConnector) Load(schema interface{}) (interface{}, error) {
	fSchema, ok := schema.(ElasticSchema)
	if !ok {
		return nil, fmt.Errorf("no ElasticSchema detected")
	}
	if fSchema.URI == "" {
		return nil, fmt.Errorf("empty URI")
	}
	return ElasticCredentials(fSchema), nil
}

func (r ElasticRawConnector) Schema() interface{} {
	return ElasticSchema{}
}

func init() {
	gautocloud.RegisterConnector(NewElasticRawConnector())
}

func NewElasticRawConnector() connectors.Connector {
	return &ElasticRawConnector{}
}
