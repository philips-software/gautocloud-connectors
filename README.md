# gautocloud
Gautocloud provides a simple abstraction that golang based applications can use to discover information about the cloud environment on which they are running, to connect to services automatically with ease of use in mind. It provides out-of-the-box support for discovering common services on Heroku, Cloud Foundry and kubernetes cloud platforms, and it supports custom automatic connectors. For more details see the [gautocloud project](https://github.com/cloudfoundry-community/gautocloud)

# gautocloud connector for HSDP CF services
This repository contains [gautocloud connectors](https://github.com/cloudfoundry-community/gautocloud) for select [HSDP](https://www.hsdp.io) Cloudfoundry services. At this time the following connectors are supported:

- [gautocloud](#gautocloud)
- [gautocloud connector for HSDP CF services](#gautocloud-connector-for-hsdp-cf-services)
- [usage](#usage)
  - [Twilio Raw](#twilio-raw)
  - [Twilio Go client](#twilio-go-client)
  - [DynamoDB client](#dynamodb-client)
  - [Redshift](#redshift)
- [Author](#author)
- [License](#license)

# usage

```go
import (
    "github.com/cloudfoundry-community/gautocloud"
    "github.com/hsdp/gautocloud-connectors/hsdp"
)
```
## Twilio Raw

```go
svc, err := gautocloud.GetFirst("hsdp:twilio-raw")
if err == nil {
    account, ok := svc.(hsdp.TwilioSubAccount)
    if ok {
        fmt.Printf("Loaded Twilio subaccount with SID: %s\n", account.SID)
    }
}
```

## Twilio Go client

```go
import (
    "github.com/kevinburke/twilio-go"
)
```

```go
svc, err := gautocloud.GetFirst("hsdp:twilio-client")
if err == nil {
    client, ok := svc.(*twilio.Client)
    if ok {
        // Iterate over calls
        iterator := client.Calls.GetPageIterator(url.Values{})
        for {
            page, err := iterator.Next(context.TODO())
            if err == twilio.NoMoreResults {
                break
            }
            fmt.Println("start", page.Start)
        }
    }
}
```

## DynamoDB client

```go
import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)
```

```go
db, err := gautocloud.GetFirst("hsdp:dynamodb-client")
client, ok := db.(*hsdp.DynamoDBClient)
if ok {
        fmt.Printf("Loaded DynamoDB client, table: %s\n", client.TableName)
        req := &dynamodb.DescribeTableInput{
                TableName: aws.String(client.TableName),
        }
        // Fetches and display details of the DynamoDB table
        result, err := client.DescribeTable(req)
        if err != nil {
                fmt.Printf("%s", err)
        }
        table := result.Table
        fmt.Printf("done: %v\n", table)
}
```

## Redshift

```go

import (
    "github.com/cloudfoundry-community/gautocloud"
    "github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
    _ "github.com/hsdp/gautocloud-connectors/hsdp"
)
```

```go
	var rs *dbtype.PostgresqlDB

	err := gautocloud.InjectFromId("hsdp:redshift", &rs)

	if err != nil {
		fmt.Printf("failed to find attached database: %v\n", err)
		return
	}

	rows, err := rs.Query("SELECT now()")

	if err != nil {
		fmt.Printf("query error: %v\n", err)
		return
    }
    
	for rows.Next() {
		var col string
		rows.Scan(&col)
		fmt.Printf("%v\n", col)
	}
```

# Author

See [AUTHORS.TXT](AUTHORS.txt)

# License

License is MIT
