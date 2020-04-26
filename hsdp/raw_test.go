package hsdp_test

import (
	. "gautocloud-connectors/hsdp"
	"github.com/cloudfoundry-community/gautocloud/connectors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Raw", func() {
	var connector connectors.Connector
	Context("Twilio", func() {
		BeforeEach(func() {
			connector = NewTwilioRawConnector()
		})
		It("Should return a TwilioSubAccount struct when passing a TwilioSchema", func() {
			data, err := connector.Load(TwilioSchema{
				TwilioAuthToken: "StrongPassw0rd",
				TwilioSID:       "AKFooBar",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).Should(BeEquivalentTo(
				TwilioSubAccount{
					SID:       "AKFooBar",
					AuthToken: "StrongPassw0rd",
				},
			))
		})
	})
	Context("DynamoDB", func() {
		BeforeEach(func() {
			connector = NewDynamoDBRawConnector()
		})
		It("Should return a DynamoDBSchema like struct when passing a DynamoDBSchema", func() {
			data, err := connector.Load(DynamoDBSchema{
				AWSKey:    "some-key",
				AWSRegion: "us-east-1",
				AWSSecret: "StrongPassw0rd",
				TableName: "table-name",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).Should(BeEquivalentTo(
				DynamoDBSchema{
					AWSKey:    "some-key",
					AWSRegion: "us-east-1",
					AWSSecret: "StrongPassw0rd",
					TableName: "table-name",
				},
			))

		})
	})
	Context("VaultRaw", func() {
		BeforeEach(func() {
			connector = NewVaultRawConnector()
		})
		It("Should return a VaultCredential like struct when passing a VaultSchema", func() {
			data, err := connector.Load(VaultSchema{
				SecretID:           "some-key",
				RoleID:             "us-east-1",
				Endpoint:           "http://vault.local",
				ServiceSecretPath:  "/service/secret",
				SpaceSecretPath:    "/space/secret",
				ServiceTransitPath: "/transit/path",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).Should(BeEquivalentTo(
				VaultCredentials{
					SecretID:           "some-key",
					RoleID:             "us-east-1",
					Endpoint:           "http://vault.local",
					ServiceSecretPath:  "/service/secret",
					SpaceSecretPath:    "/space/secret",
					ServiceTransitPath: "/transit/path",
				},
			))

		})
	})
	Context("Redshift", func() {
		BeforeEach(func() {
			connector = NewRedshiftRawConnector()
		})
		It("Should return a RedshiftCredentials struct when passing a RedshiftSchema", func() {
			data, err := connector.Load(RedshiftSchema{
				Password:     "StrongPassw0rd",
				Username:     "AKFooBar",
				DatabaseName: "hsdpredhsift",
				Hostname:     "foo.bar.com",
				Port:         5349,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).Should(BeEquivalentTo(
				RedshiftCredentials{
					Password:     "StrongPassw0rd",
					Username:     "AKFooBar",
					DatabaseName: "hsdpredhsift",
					Hostname:     "foo.bar.com",
					Port:         5349,
				},
			))
		})
	})
})
