package hsdp

import (
	"fmt"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
)

type PostgresSQLSchema struct {
	AutomatedSnapshotsPassword string `cloud:"automated_snapshots_password"`
	DBName                     string `cloud:"db_name"`
	Hostname                   string `cloud:"hostname"`
	Password                   string `cloud:"password"`
	Port                       int    `cloud:"port"`
	URI                        string `cloud:"uri"`
	Username                   string `cloud:"username"`
	Options                    string
}

func init() {
	gautocloud.RegisterConnector(NewPostgresSQLRawConnector())
}

type PostgresSQLRawConnector struct{}

func NewPostgresSQLRawConnector() connectors.Connector {
	return &PostgresSQLRawConnector{}
}

func (c PostgresSQLRawConnector) Id() string {
	return "hsdp:postgresql-raw"
}
func (c PostgresSQLRawConnector) Name() string {
	return ".*postgres.*"
}
func (c PostgresSQLRawConnector) Tags() []string {
	return []string{"PostgreSQL.*", "RDS"}
}
func (c PostgresSQLRawConnector) Load(schema interface{}) (interface{}, error) {
	fSchema, ok := schema.(PostgresSQLSchema)
	if !ok {
		return nil, fmt.Errorf("no PostgresSQLSchema detected")
	}
	if fSchema.DBName == "" || fSchema.Hostname == "" {
		return nil, fmt.Errorf("empty database name or hostname")
	}
	return fSchema, nil
}

func (c PostgresSQLRawConnector) Schema() interface{} {
	return PostgresSQLSchema{}
}
