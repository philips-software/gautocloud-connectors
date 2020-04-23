package hsdp_test

import (
	"github.com/cloudfoundry-community/gautocloud/connectors"
	. "gautocloud-connectors/hsdp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	var connector connectors.Connector
	Context("Twilio", func() {
		BeforeEach(func() {
			connector = NewTwilioClientConnector()
		})
		It("Should return a Twilio Go client when passing a TwilioSchema", func() {
			data, err := connector.Load(TwilioSchema{
				TwilioAuthToken: "StrongPassw0rd",
				TwilioSID:       "AKFooBar",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).NotTo(BeNil())
		})
	})
	Context("DynamoDB", func() {
		BeforeEach(func() {
			connector = NewDynamoDBClientConnector()
		})
		It("Should return a DynamoDB Go client when passing a DynamoDBSchema", func() {
			data, err := connector.Load(DynamoDBSchema{
				AWSKey:    "some-key",
				AWSRegion: "us-east-1",
				AWSSecret: "StrongPassw0rd",
				TableName: "table-name",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).NotTo(BeNil())
		})
	})

})
