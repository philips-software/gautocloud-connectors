package hsdp

import (
	"fmt"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	_ "github.com/cloudfoundry-community/gautocloud/connectors"
)

type S3Schema struct {
	APIKey             string      `json:"api_key"`
	Bucket             string      `json:"bucket"`
	Endpoint           string      `json:"endpoint"`
	LocationConstraint interface{} `json:"location_constraint"`
	SecretKey          string      `json:"secret_key"`
	URI                string      `json:"uri"`
}

type S3Credentials S3Schema

type S3RawConnector struct{}

func (r S3RawConnector) Id() string {
	return "hsdp:s3-raw"
}

func (r S3RawConnector) Name() string {
	return ".*s3.*"
}

func (r S3RawConnector) Tags() []string {
	return []string{"S3"}
}

func (r S3RawConnector) Load(schema interface{}) (interface{}, error) {
	fSchema, ok := schema.(S3Schema)
	if !ok {
		return nil, fmt.Errorf("no S3Schema detected")
	}
	if fSchema.Bucket == "" || fSchema.Endpoint == "" {
		return nil, fmt.Errorf("empty bucket or endpoint")
	}
	return S3Credentials(fSchema), nil
}

func (r S3RawConnector) Schema() interface{} {
	return S3Schema{}
}

func init() {
	gautocloud.RegisterConnector(NewS3RawConnector())
}

func NewS3RawConnector() connectors.Connector {
	return &S3RawConnector{}
}
