package hsdp

import (
	"fmt"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	_ "github.com/cloudfoundry-community/gautocloud/connectors"
)

type RedshiftSchema struct {
	DatabaseName string `cloud:"db_name"`
	Hostname       string `cloud:"hostname"`
	Username string `cloud:"username"`
	Password string `cloud:"password"`
	Port     int  `cloud:"port"`
	URI      string `cloud:"uri"`
}

type RedshiftCredentials RedshiftSchema

type RedshiftRawConnector struct {}

func (r RedshiftRawConnector) Id() string {
	return "hsdp:redshift-raw"
}

func (r RedshiftRawConnector) Name() string {
	return ".*redshift.*"
}

func (r RedshiftRawConnector) Tags() []string {
	return []string{"Redshift.*", "redshift.*"}
}

func (r RedshiftRawConnector) Load(schema interface{}) (interface{}, error) {
	fSchema, ok := schema.(RedshiftSchema)
	if !ok || fSchema.DatabaseName == "" || fSchema.Hostname == "" {
		return nil, fmt.Errorf("empty database name or hostname")
	}
	return RedshiftCredentials(fSchema), nil
}

func (r RedshiftRawConnector) Schema() interface{} {
	return RedshiftCredentials{}
}

func init() {
	gautocloud.RegisterConnector(NewRedshiftRawConnector())
}

func NewRedshiftRawConnector() connectors.Connector {
	return &RedshiftRawConnector{}
}