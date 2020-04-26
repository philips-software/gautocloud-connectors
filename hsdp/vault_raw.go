package hsdp

import (
	"fmt"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
)

type VaultSchema struct {
	Endpoint           string `json:"endpoint"`
	OrgSecretPath      string `json:"org_secret_path"`
	RoleID             string `json:"role_id"`
	SecretID           string `json:"secret_id"`
	ServiceSecretPath  string `json:"service_secret_path"`
	ServiceTransitPath string `json:"service_transit_path"`
	SpaceSecretPath    string `json:"space_secret_path"`
}

type VaultCredentials VaultSchema

type VaultRawConnector struct{}

func init() {
	gautocloud.RegisterConnector(NewVaultRawConnector())
}

func (v VaultRawConnector) Id() string {
	return "hsdp:vault-raw"
}

func (v VaultRawConnector) Name() string {
	return ".*vault.*"
}

func (v VaultRawConnector) Tags() []string {
	return []string{"Vault.*", "vault.*"}
}

func (v VaultRawConnector) Load(schema interface{}) (interface{}, error) {
	fSchema, ok := schema.(VaultSchema)
	if !ok {
		return nil, fmt.Errorf("no VaultSchema detected")
	}
	if fSchema.RoleID == "" || fSchema.SecretID == "" {
		return nil, fmt.Errorf("vault approle credentials missing")
	}
	return VaultCredentials(fSchema), nil
}

func (v VaultRawConnector) Schema() interface{} {
	return VaultSchema{}
}

func NewVaultRawConnector() connectors.Connector {
	return &VaultRawConnector{}
}
