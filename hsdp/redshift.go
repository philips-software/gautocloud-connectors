package hsdp

import (
	"database/sql"
	"fmt"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
	_ "github.com/lib/pq"
)

func init() {
	gautocloud.RegisterConnector(NewRedshiftConnector())
}

type RedshiftConnector struct {
	redshiftRawConnector connectors.Connector
}

func (r RedshiftConnector) Id() string {
	return "hsdp:redshift"
}

func (r RedshiftConnector) Name() string {
	return r.redshiftRawConnector.Name()
}

func (r RedshiftConnector) Tags() []string {
	return r.redshiftRawConnector.Tags()
}

func (r RedshiftConnector) GetConnString(schema RedshiftSchema) string {
	connString := "postgres://" + schema.Username
	if schema.Password != "" {
		connString += ":" + schema.Password
	}
	connString += fmt.Sprintf("@tcp(%s:%d)/%s", schema.Hostname, schema.Port, schema.DatabaseName)
	schema.Options = "sslmode=prefer"
	if schema.Options != "" {
		connString += "?" + schema.Options
	}
	return connString
}

func (r RedshiftConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := r.redshiftRawConnector.Load(schema)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", r.GetConnString(schema.(RedshiftSchema)))
	if err != nil {
		return db, err
	}
	return &dbtype.PostgresqlDB{db}, nil
}

func (r RedshiftConnector) Schema() interface{} {
	return r.redshiftRawConnector.Schema()
}

func NewRedshiftConnector() connectors.Connector {
	return &RedshiftConnector{
		redshiftRawConnector: NewRedshiftRawConnector(),
	}
}
