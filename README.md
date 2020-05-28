# gautocloud
Gautocloud provides a simple abstraction that golang based applications can use to discover information about the cloud environment on which they are running, to connect to services automatically with ease of use in mind. It provides out-of-the-box support for discovering common services on Heroku, Cloud Foundry and kubernetes cloud platforms, and it supports custom automatic connectors. For more details see the [gautocloud project](https://github.com/cloudfoundry-community/gautocloud)

# gautocloud connectors for HSDP services
This repository contains [gautocloud connectors](https://github.com/cloudfoundry-community/gautocloud) for select [HSDP](https://www.hsdp.io) Cloud foundry services. At this time the following connectors are supported:

  - [Twilio Raw](#twilio-raw)
  - [Twilio Go client](#twilio-go-client)
  - [DynamoDB client](#dynamodb-client)
  - [Redshift](#redshift)
  - [Vault client](#vault-client)
  - [S3 client](#S3-client)

# usage
Import the packages in your app, this will register all the supported connectors and you can proceed to detect the services you need:

```go
import (
    "github.com/cloudfoundry-community/gautocloud"
    "github.com/philips-software/gautocloud-connectors/hsdp"
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
    "github.com/cloudfoundry-community/gautocloud"
    _ "github.com/philips-software/gautocloud-connectors/hsdp"
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
    "github.com/cloudfoundry-community/gautocloud"
    _ "github.com/philips-software/gautocloud-connectors/hsdp"
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
    _ "github.com/philips-software/gautocloud-connectors/hsdp"
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

## Vault client
A VaultClient instance is returns which is composed of the official Hashicorp Go Vault client and a VaultCredentials struct containing all the fields found in the service broker

```go
import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/philips-software/gautocloud-connectors/hsdp"
)

```

```go
	var vaultClient *hsdp.VaultClient

	err := gautocloud.Inject(&vaultClient)

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	secret, err := vaultClient.WriteSpaceSecret("mysecret",
	 	map[string]interface{}{
			"foo": "bar",
			"alice": "bob",
			"bob": "trudy",
		})
		
```

## S3 client
An S3Client instance is returns which is composed of an AWS S3 Credentials Session and S3Credentials struct containing all the fields found in the service broker credentials

```go
package main

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/philips-software/gautocloud-connectors/hsdp"
)

func main() {
	var svc *hsdp.S3Client

	err := gautocloud.Inject(&svc)

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("cf-s3-mys3-buck-etaa-here-b3e81da982fa"),
		Key:    aws.String("/public/somefile.zip"),
	})
	str, err := req.Presign(15 * time.Minute)

	fmt.Println("The URL is:", str, " err:", err)
}
```


# Contact / Getting help

andy.lo-a-foe@philips.com

# Application examples

Coming soon..

# License

License is [MIT](LICENSE.md)
