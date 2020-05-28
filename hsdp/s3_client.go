package hsdp

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
)

func init() {
	gautocloud.RegisterConnector(NewS3ClientConnector())
}

type S3ClientConnector struct {
	s3RawConnector connectors.Connector
}

type S3Client struct {
	*s3.S3
	S3Credentials
}

func (v S3ClientConnector) Id() string {
	return "hsdp:s3-client"
}

func (v S3ClientConnector) Name() string {
	return v.s3RawConnector.Name()
}

func (v S3ClientConnector) Tags() []string {
	return v.s3RawConnector.Tags()
}

func (v S3ClientConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := v.s3RawConnector.Load(schema)
	if err != nil {
		return nil, err
	}
	fSchema := schema.(S3Credentials)

	region := "us-east-1" // Not sure if this is a good default
	if str, ok := fSchema.LocationConstraint.(string); ok {
		region = str
	}
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(fSchema.APIKey, fSchema.SecretKey, ""),
	})
	// Create S3 service client
	svc := s3.New(sess)

	S3Client := &S3Client{
		S3:            svc,
		S3Credentials: fSchema,
	}
	return S3Client, nil
}

func (v S3ClientConnector) Schema() interface{} {
	return v.s3RawConnector.Schema()
}

func NewS3ClientConnector() connectors.Connector {
	return &S3ClientConnector{
		s3RawConnector: NewS3RawConnector(),
	}
}
