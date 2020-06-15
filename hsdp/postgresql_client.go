package hsdp

import (
	"database/sql"
	"fmt"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	_ "github.com/lib/pq"
)

func init() {
	gautocloud.RegisterConnector(NewPostgresSQLClientConnector())
}

type PostgreSQLClient struct {
	*sql.DB
}

type PostgreSQLClientConnector struct {
	wrapRawConn connectors.Connector
}

func NewPostgresSQLClientConnector() connectors.Connector {
	return &PostgreSQLClientConnector{
		wrapRawConn: NewPostgresSQLRawConnector(),
	}
}

func (c PostgreSQLClientConnector) Id() string {
	return "hsdp:postgresql-client"
}

func (c PostgreSQLClientConnector) Name() string {
	return c.wrapRawConn.Name()
}

func (c PostgreSQLClientConnector) Tags() []string {
	return c.wrapRawConn.Tags()
}

func (c PostgreSQLClientConnector) GetConnString(schema PostgresSQLSchema) string {
	connString := "postgres://" + schema.Username
	if schema.Password != "" {
		connString += ":" + schema.Password
	}
	connString += fmt.Sprintf("@%s:%d/%s", schema.Hostname, schema.Port, schema.DBName)
	if schema.Options != "" {
		connString += "?" + schema.Options
	}
	return connString
}

func (c PostgreSQLClientConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := c.wrapRawConn.Load(schema)
	if err != nil {
		return nil, err
	}
	fSchema := schema.(PostgresSQLSchema)

	db, err := sql.Open("postgres", c.GetConnString(fSchema))
	if err != nil {
		return db, err
	}

	if err != nil {
		return nil, err
	}
	return &PostgreSQLClient{
		DB: db,
	}, nil
}

func (c PostgreSQLClientConnector) Schema() interface{} {
	return c.wrapRawConn.Schema()
}
