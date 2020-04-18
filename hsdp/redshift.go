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

func (r RedshiftConnector) GetConnString(creds RedshiftCredentials) string {
	connString := "user=" + creds.Username
	if creds.Password != "" {
		connString += " password=" + creds.Password
	}
	connString += fmt.Sprintf(" host=%s port=%d dbname=%s", creds.Hostname, creds.Port, creds.DatabaseName)
	creds.Options = "sslmode=require"
	if creds.Options != "" {
		connString += " " + creds.Options
	}
	return connString
}

func (r RedshiftConnector) Load(schema interface{}) (interface{}, error) {
	creds, err := r.redshiftRawConnector.Load(schema)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", r.GetConnString(creds.(RedshiftCredentials)))
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
