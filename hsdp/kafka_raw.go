package hsdp

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	_ "github.com/cloudfoundry-community/gautocloud/connectors"
)

type KafkaSchema struct {
	Username  string   `json:"username,omitempty"`
	Password  string   `json:"password,omitempty"`
	Hostname  string   `json:"hostname,omitempty"`
	Port      int      `json:"port,omitempty"`
	Hostnames []string `json:"hostnames"`
	URI       string   `json:"uri"`
	CACert    string   `json:"ca_cert,omitempty"`
}

type KafkaCredentials KafkaSchema

type KafkaRawConnector struct{}

func (k KafkaCredentials) PEMData() ([]byte, error) {
	pemData, err := base64.StdEncoding.DecodeString(k.CACert)
	if err != nil {
		return []byte(""), err
	}
	return pemData, nil
}

func (k KafkaCredentials) CACertificate() (*x509.Certificate, error) {
	pemData, err := k.PEMData()
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("invalid PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("ParseCertificate: %w", err)
	}
	return cert, nil
}

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
	if !strings.Contains(fSchema.URI, "kafka") {
		return nil, fmt.Errorf("no hostnames or invalid URI [%s]", fSchema.URI)
	}
	return KafkaCredentials(fSchema), nil
}

func (r KafkaRawConnector) Schema() interface{} {
	return KafkaSchema{}
}

func NewKafkaRawConnector() connectors.Connector {
	return &KafkaRawConnector{}
}
