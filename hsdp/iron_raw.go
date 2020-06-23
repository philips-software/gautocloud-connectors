package hsdp

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
)

func init() {
	gautocloud.RegisterConnector(NewIronRawConnector())
}

type IronSchema struct {
	ClusterInfo []struct {
		ClusterID   string `cloud:"cluster_id"`
		ClusterName string `cloud:"cluster_name"`
		Pubkey      string `cloud:"pubkey"`
		UserID      string `cloud:"user_id"`
	} `cloud:"cluster_info"`
	Email     string `cloud:"email"`
	Password  string `cloud:"password"`
	Project   string `cloud:"project"`
	ProjectID string `cloud:"project_id"`
	Token     string `cloud:"token"`
	UserID    string `cloud:"user_id"`
}

type IronRawConnector struct{}

func (i IronRawConnector) Id() string {
	return "hsdp:iron-raw"
}

func (i IronRawConnector) Name() string {
	return ".*iron.*"
}

func (i IronRawConnector) Tags() []string {
	return []string{"iron.*", "hsdp-iron", "iron.io"}
}

func (i IronRawConnector) Load(schema interface{}) (interface{}, error) {
	fSchema := schema.(IronSchema)
	return fSchema, nil
}

func (i IronRawConnector) Schema() interface{} {
	return IronSchema{}
}

func NewIronRawConnector() connectors.Connector {
	return &IronRawConnector{}
}
