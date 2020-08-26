package hsdp

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func init() {
	gautocloud.RegisterConnector(NewS3MinioClientConnector())
}

type S3MinioClientConnector struct {
	s3RawConnector connectors.Connector
}

type S3MinioClient struct {
	*minio.Client
	S3Credentials
}

func (v S3MinioClientConnector) Id() string {
	return "hsdp:s3-minio-client"
}

func (v S3MinioClientConnector) Name() string {
	return v.s3RawConnector.Name()
}

func (v S3MinioClientConnector) Tags() []string {
	return v.s3RawConnector.Tags()
}

func (v S3MinioClientConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := v.s3RawConnector.Load(schema)
	if err != nil {
		return nil, err
	}
	fSchema := schema.(S3Credentials)

	// Initialize minio client object.
	minioClient, err := minio.New(fSchema.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(fSchema.APIKey, fSchema.SecretKey, ""),
		Secure: true,
	})
	S3MinioClient := &S3MinioClient{
		Client:        minioClient,
		S3Credentials: fSchema,
	}
	return S3MinioClient, nil
}

func (v S3MinioClientConnector) Schema() interface{} {
	return v.s3RawConnector.Schema()
}

func NewS3MinioClientConnector() connectors.Connector {
	return &S3MinioClientConnector{
		s3RawConnector: NewS3RawConnector(),
	}
}
