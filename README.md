# gautocloud
Gautocloud provides a simple abstraction that golang based applications can use to discover information about the cloud environment on which they are running, to connect to services automatically with ease of use in mind. It provides out-of-the-box support for discovering common services on Heroku, Cloud Foundry and kubernetes cloud platforms, and it supports custom automatic connectors. For more details see the [gautocloud project](https://github.com/cloudfoundry-community/gautocloud)

# gautocloud connectors for HSDP services
This repository contains [gautocloud connectors](https://github.com/cloudfoundry-community/gautocloud) for select [HSDP](https://www.hsdp.io) Cloud foundry services. At this time the following connectors are supported:

  - [Twilio Raw](#twilio-raw)
  - [Twilio Go client](#twilio-go-client)
  - [DynamoDB client](#dynamodb-client)
  - [Redshift](#redshift)

# usage
Import the packages in your app, this will register all the supported connectors and you can proceed to detect the services you need:

```go
import (
    "github.com/cloudfoundry-community/gautocloud"
    "github.com/hsdp/gautocloud-connectors/hsdp"
)
```
## Twilio Raw
This returns the credentials so you can use them as you see fit

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
This returns a configured Twilio Go client based on [github.com/kevinburke/twilio-go](https://github.com/kevinburke/twilio-go)

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
A configured DynamDBClient based on the [AWS SDK](https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/)

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
This returns a configured (wrapped) PostgreSQLDB connection ready for use.

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

# Contact / Getting help

andy.lo-a-foe@philips.com

# Application examples

Coming soon..

# License

License is [MIT](LICENSE.md)
