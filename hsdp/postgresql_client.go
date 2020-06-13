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

type PostgresSQLClient struct {
	*sql.DB
}

type PostgresSQLClientConnector struct {
	wrapRawConn connectors.Connector
}

func NewPostgresSQLClientConnector() connectors.Connector {
	return &PostgresSQLClientConnector{
		wrapRawConn: NewPostgresSQLRawConnector(),
	}
}

func (c PostgresSQLClientConnector) Id() string {
	return "hsdp:postgresql-client"
}

func (c PostgresSQLClientConnector) Name() string {
	return c.wrapRawConn.Name()
}

func (c PostgresSQLClientConnector) Tags() []string {
	return c.wrapRawConn.Tags()
}

func (c PostgresSQLClientConnector) GetConnString(schema PostgresSQLSchema) string {
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

func (c PostgresSQLClientConnector) Load(schema interface{}) (interface{}, error) {
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
	return &PostgresSQLClient{
		DB: db,
	}, nil
}

func (c PostgresSQLClientConnector) Schema() interface{} {
	return c.wrapRawConn.Schema()
}
