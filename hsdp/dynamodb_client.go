package hsdp

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func init() {
	gautocloud.RegisterConnector(NewDynamoDBClientConnector())
}

type DynamoDBClient struct {
	*dynamodb.DynamoDB
	TableName string
}

type DynamoDBClientConnector struct {
	wrapRawConn connectors.Connector
}

func NewDynamoDBClientConnector() connectors.Connector {
	return &DynamoDBClientConnector{
		wrapRawConn: NewDynamoDBRawConnector(),
	}
}

func (c DynamoDBClientConnector) Id() string {
	return "hsdp:dynamodb-client"
}

func (c DynamoDBClientConnector) Name() string {
	return ".*dynamodb.*"
}

func (c DynamoDBClientConnector) Tags() []string {
	return []string{"DynamoDB.*"}
}

func (c DynamoDBClientConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := c.wrapRawConn.Load(schema)
	if err != nil {
		return nil, err
	}
	fSchema := schema.(DynamoDBSchema)

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(fSchema.AWSRegion),
		Credentials: credentials.NewStaticCredentials(fSchema.AWSKey, fSchema.AWSSecret, ""),
	})
	if err != nil {
		return nil, err
	}
	// Create DynamoDB client
	client := dynamodb.New(sess)

	return &DynamoDBClient{
		DynamoDB:  client,
		TableName: fSchema.TableName,
	}, nil
}

func (c DynamoDBClientConnector) Schema() interface{} {
	return c.wrapRawConn.Schema()
}
