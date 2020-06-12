package hsdp

import (
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
	return []string{"PostgreSQL.*"}
}
func (c PostgresSQLRawConnector) Load(schema interface{}) (interface{}, error) {
	fSchema := schema.(PostgresSQLSchema)
	return fSchema, nil
}

func (c PostgresSQLRawConnector) Schema() interface{} {
	return PostgresSQLSchema{}
}
