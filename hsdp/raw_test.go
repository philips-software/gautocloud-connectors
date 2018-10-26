package hsdp_test

import (
	"github.com/cloudfoundry-community/gautocloud/connectors"
	. "github.com/loafoe/gautocloud-connectors/hsdp"

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
})
