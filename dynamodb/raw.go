package dynamodb

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
)

func init() {
	gautocloud.RegisterConnector(NewDynamoDBRawConnector())
}

type Schema struct {
	AWSKey    string `cloud:"aws_key"`
	AWSRegion string `cloud:"aws_region"`
	AWSSecret string `cloud:"aws_secret"`
	TableName string `cloud:"table_name"`
}

type DynamoDBRawConnector struct{}

func NewDynamoDBRawConnector() connectors.Connector {
	return &DynamoDBRawConnector{}
}

func (c DynamoDBRawConnector) Id() string {
	return "hsdp:dynamodb-raw"
}
func (c DynamoDBRawConnector) Name() string {
	return ".*dynamodb.*"
}
func (c DynamoDBRawConnector) Tags() []string {
	return []string{"DynamoDB.*"}
}
func (c DynamoDBRawConnector) Load(schema interface{}) (interface{}, error) {
	fSchema := schema.(Schema)
	return fSchema, nil
}

func (c DynamoDBRawConnector) Schema() interface{} {
	return Schema{}
}
