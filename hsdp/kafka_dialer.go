package hsdp

import (
	"crypto/tls"
	"crypto/x509"
	"time"

	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func init() {
	gautocloud.RegisterConnector(NewKafkaClientConnector())
}

type KafkaDialer struct {
	*kafka.Dialer
	KafkaCredentials
}

type KafkaDialerConnector struct {
	wrapRawConn connectors.Connector
}

func (k KafkaDialerConnector) Id() string {
	return "hsdp:kafka-client"
}

func (k KafkaDialerConnector) Name() string {
	return k.wrapRawConn.Name()
}

func (k KafkaDialerConnector) Tags() []string {
	return k.wrapRawConn.Tags()
}

func (k KafkaDialerConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := k.wrapRawConn.Load(schema)
	if err != nil {
		return nil, err
	}
	fSchema := schema.(KafkaCredentials)
	pemData, err := fSchema.PEMData()
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(pemData)
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}
	mechanism := plain.Mechanism{
		Username: fSchema.Username,
		Password: fSchema.Password,
	}
	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		TLS:           tlsConfig,
		SASLMechanism: mechanism,
	}
	return KafkaDialer{
		Dialer:           dialer,
		KafkaCredentials: fSchema,
	}, nil
}

func (k KafkaDialerConnector) Schema() interface{} {
	return k.wrapRawConn.Schema()
}

func NewKafkaClientConnector() connectors.Connector {
	return &KafkaDialerConnector{
		wrapRawConn: NewKafkaRawConnector(),
	}
}
