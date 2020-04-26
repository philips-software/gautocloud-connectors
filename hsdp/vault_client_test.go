package hsdp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStripV1(t *testing.T) {
	client := VaultClient{
		VaultCredentials: VaultCredentials{
			OrgSecretPath:      "/v1/cf/2d986b98-b3b7-47bf-aeef-ee95da274cf8/secret",
			SpaceSecretPath:    "/v1/cf/2d986b98-b3b7-47bf-aeef-ee95da274cf8/secret",
			ServiceTransitPath: "/v1/cf/2d986b98-b3b7-47bf-aeef-ee95da274cf8/secret",
			ServiceSecretPath:  "/v1/cf/2d986b98-b3b7-47bf-aeef-ee95da274cf8/secret",
		},
	}
	client.stripV1()
	assert.Equal(t, client.VaultCredentials.ServiceSecretPath, "cf/2d986b98-b3b7-47bf-aeef-ee95da274cf8/secret")
	assert.Equal(t, client.VaultCredentials.SpaceSecretPath, "cf/2d986b98-b3b7-47bf-aeef-ee95da274cf8/secret")
	assert.Equal(t, client.VaultCredentials.ServiceTransitPath, "cf/2d986b98-b3b7-47bf-aeef-ee95da274cf8/secret")
	assert.Equal(t, client.VaultCredentials.ServiceSecretPath, "cf/2d986b98-b3b7-47bf-aeef-ee95da274cf8/secret")
	client.stripV1()
	assert.Equal(t, client.VaultCredentials.ServiceSecretPath, "cf/2d986b98-b3b7-47bf-aeef-ee95da274cf8/secret")
	assert.Equal(t, client.VaultCredentials.SpaceSecretPath, "cf/2d986b98-b3b7-47bf-aeef-ee95da274cf8/secret")
	assert.Equal(t, client.VaultCredentials.ServiceTransitPath, "cf/2d986b98-b3b7-47bf-aeef-ee95da274cf8/secret")
	assert.Equal(t, client.VaultCredentials.ServiceSecretPath, "cf/2d986b98-b3b7-47bf-aeef-ee95da274cf8/secret")
}
