package hsdp

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
)

func init() {
	gautocloud.RegisterConnector(NewElastic7ClientConnector())
}

type Elastic7ClientConnector struct {
	elasticRawConnector connectors.Connector
}

type Elastic7Client struct {
	*elasticsearch7.Client
	ElasticCredentials
}

func (v Elastic7ClientConnector) Id() string {
	return "hsdp:elastic7-client"
}

func (v Elastic7ClientConnector) Name() string {
	return v.elasticRawConnector.Name()
}

func (v Elastic7ClientConnector) Tags() []string {
	return v.elasticRawConnector.Tags()
}

func (v Elastic7ClientConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := v.elasticRawConnector.Load(schema)
	if err != nil {
		return nil, err
	}
	fSchema := schema.(ElasticCredentials)

	// Initialize elastic client object.
	es7, err := elasticsearch7.NewClient(elasticsearch7.Config{
		Username:  fSchema.Username,
		Password:  fSchema.Password,
		Addresses: []string{fSchema.Server()},
	})
	if err != nil {
		return nil, err
	}
	client := &Elastic7Client{
		Client:             es7,
		ElasticCredentials: fSchema,
	}
	return client, nil
}

func (v Elastic7ClientConnector) Schema() interface{} {
	return v.elasticRawConnector.Schema()
}

func NewElastic7ClientConnector() connectors.Connector {
	return &Elastic7ClientConnector{
		elasticRawConnector: NewElasticRawConnector(),
	}
}
