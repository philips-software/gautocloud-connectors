# DynamoDB

```go
import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/cloudfoundry-community/gautocloud"
    gdynamodb "github.com/loafoe/gautocloud-connectors/dynamodb"
)

```

```go
	db, err := gautocloud.GetFirst("hsdp:dynamodb-client")
	service, ok := db.(*gdynamodb.DynamoDBClient)
	if ok {
		fmt.Printf("Loaded DynamoDB client, table: %s\n", service.TableName)
		req := &dynamodb.DescribeTableInput{
			TableName: aws.String(service.TableName),
		}
		result, err := service.DescribeTable(req)
		if err != nil {
			fmt.Printf("%s", err)
		}
		table := result.Table
		fmt.Printf("done", table)
	} else {
		fmt.Printf("Loaded unknown: %v\n", service)
	}
```
