package cups_test

import (
	"github.com/cloudfoundry-community/gautocloud/connectors"
	. "gautocloud-connectors/cups"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cups", func() {
	var connector connectors.Connector
	Context("Eureka", func() {
		BeforeEach(func() {
			connector = NewEurekaConnector()
		})
		It("Should return a Go Eureka client struct when passing a EurekaSchema", func() {
			data, err := connector.Load(EurekaSchema{
				URI: "https://eureka-service",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).ShouldNot(BeNil())
		})
	})
})
