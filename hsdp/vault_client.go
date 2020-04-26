package hsdp

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	vault "github.com/hashicorp/vault/api"
	"strings"
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

func (v *VaultClient) WriteOrgSecret(path string, data map[string]interface{}) (*vault.Secret, error) {
	return v.Client.Logical().Write(v.OrgSecretPath+"/"+path, data)
}

func (v *VaultClient) ReadOrgSecret(path string) (*vault.Secret, error) {
	return v.Client.Logical().Read(v.OrgSecretPath + "/" + path)
}

func (v *VaultClient) WriteSpaceSecret(path string, data map[string]interface{}) (*vault.Secret, error) {
	return v.Client.Logical().Write(v.SpaceSecretPath+"/"+path, data)
}

func (v *VaultClient) ReadSpaceSecret(path string) (*vault.Secret, error) {
	return v.Client.Logical().Read(v.SpaceSecretPath + "/" + path)
}

func (v *VaultClient) WriteServiceSecret(path string, data map[string]interface{}) (*vault.Secret, error) {
	return v.Client.Logical().Write(v.ServiceSecretPath+"/"+path, data)
}

func (v *VaultClient) ReadServiceSecret(path string) (*vault.Secret, error) {
	return v.Client.Logical().Read(v.ServiceSecretPath + "/" + path)
}
func (v *VaultClient) WriteServiceTransit(path string, data map[string]interface{}) (*vault.Secret, error) {
	return v.Client.Logical().Write(v.ServiceTransitPath+"/"+path, data)
}

func (v *VaultClient) ReadServiceTransit(path string) (*vault.Secret, error) {
	return v.Client.Logical().Read(v.ServiceTransitPath + "/" + path)
}

func (v *VaultClient) stripV1() {
	for _, s := range []*string{
		&v.OrgSecretPath,
		&v.ServiceTransitPath,
		&v.SpaceSecretPath,
		&v.ServiceSecretPath,
	} {
		if parts := strings.Split(*s, "/"); len(parts) > 2 && parts[1] == "v1" {
			*s = strings.Join(parts[2:], "/")
		}
	}
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
	client.SetAddress(fSchema.Endpoint)
	secret, err := client.Logical().Write("auth/approle/login", map[string]interface{}{
		"role_id":   fSchema.RoleID,
		"secret_id": fSchema.SecretID,
	})
	if err != nil {
		return nil, err
	}
	client.SetToken(secret.Auth.ClientToken)
	client.Auth().Token().RenewSelf(1800)

	vaultClient := &VaultClient{
		Client:           client,
		VaultCredentials: fSchema,
	}
	vaultClient.stripV1()
	return vaultClient, nil
}

func (v VaultClientConnector) Schema() interface{} {
	return v.vaultRawConnector.Schema()
}

func NewVaultClientConnector() connectors.Connector {
	return &VaultClientConnector{
		vaultRawConnector: NewVaultRawConnector(),
	}
}
