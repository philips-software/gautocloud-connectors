# CUPS Eureka client

This connector supports a user defined Eureka service in Cloudfoundry. Defining such a service in Cloudfoundry:

```
cf cups -p '{"uri":"https://eureka-service-url"}' -t eureka eureka-service
```

The connector uses the tag to discover the service. It returns an instance of the Go Eureka client

## Example

```go
import (
    "github.com/cloudfoundry-community/gautocloud"
    "github.com/hsdp/gautocloud-connectors/cups"
    "github.com/loafoe/go-eureka-client/eureka"
)
```


```
    var client *eureka.Client

    err := gautocloud.Inject(&client)
    if err != nil {
        panic(err)
    }
    instance := eureka.NewInstanceInfo("my-service.foo.com", "my-service", "10.0.0.1", 443, 30, true) //Create a new instance to register
    err := client.RegisterInstance("my-service", instance) // Register new instance in your eureka(s)
    if err != nil {
        fmt.Printf("Error registering: %v\n", err)
    }
    applications, _ := client.GetApplications() // Retrieves all applications from eureka server(s)
```
