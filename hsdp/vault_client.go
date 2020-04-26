package hsdp

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	vault "github.com/hashicorp/vault/api"
)

func init() {
	gautocloud.RegisterConnector(NewVaultClientConnector())
}

type VaultClientConnector struct {
	vaultRawConnector connectors.Connector
}

type VaultClient struct {
	*vault.Client
	VaultCredentials
}

func (v VaultClientConnector) Id() string {
	return "hsdp:vault-client"
}

func (v VaultClientConnector) Name() string {
	return v.vaultRawConnector.Name()
}

func (v VaultClientConnector) Tags() []string {
	return v.vaultRawConnector.Tags()
}

func (v VaultClientConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := v.vaultRawConnector.Load(schema)
	if err != nil {
		return nil, err
	}
	fSchema := schema.(VaultCredentials)

	client, err := vault.NewClient(vault.DefaultConfig())
	if err != nil {
		return nil, err
	}
	secret, err := client.Logical().Write("auth/approle/login", map[string]interface{}{
		"role_id":   fSchema.RoleID,
		"secret_id": fSchema.SecretID,
	})
	if err != nil {
		return nil, err
	}
	client.SetToken(secret.Auth.ClientToken)
	client.Auth().Token().RenewSelf(1800)

	return &VaultClient{
		Client:           client,
		VaultCredentials: fSchema,
	}, nil
}

func (v VaultClientConnector) Schema() interface{} {
	return v.vaultRawConnector.Schema()
}

func NewVaultClientConnector() connectors.Connector {
	return &VaultClientConnector{
		vaultRawConnector: NewVaultRawConnector(),
	}
}
