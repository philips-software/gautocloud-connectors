package hsdp

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	elasticsearch6 "github.com/elastic/go-elasticsearch/v6"
)

func init() {
	gautocloud.RegisterConnector(NewElastic6ClientConnector())
}

type Elastic6ClientConnector struct {
	elasticRawConnector connectors.Connector
}

type Elastic6Client struct {
	*elasticsearch6.Client
	ElasticCredentials
}

func (v Elastic6ClientConnector) Id() string {
	return "hsdp:elastic6-client"
}

func (v Elastic6ClientConnector) Name() string {
	return v.elasticRawConnector.Name()
}

func (v Elastic6ClientConnector) Tags() []string {
	return v.elasticRawConnector.Tags()
}

func (v Elastic6ClientConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := v.elasticRawConnector.Load(schema)
	if err != nil {
		return nil, err
	}
	fSchema := schema.(ElasticCredentials)

	// Initialize elastic client object.
	es6, err := elasticsearch6.NewClient(elasticsearch6.Config{
		Username:  fSchema.Username,
		Password:  fSchema.Password,
		Addresses: []string{fSchema.Server()},
	})
	if err != nil {
		return nil, err
	}
	client := &Elastic6Client{
		Client:             es6,
		ElasticCredentials: fSchema,
	}
	return client, nil
}

func (v Elastic6ClientConnector) Schema() interface{} {
	return v.elasticRawConnector.Schema()
}

func NewElastic6ClientConnector() connectors.Connector {
	return &Elastic6ClientConnector{
		elasticRawConnector: NewElasticRawConnector(),
	}
}
