package hsdp

import (
	"fmt"

	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/go-redis/redis/v8"
)

type RedisCredentials struct {
	Hostname     string `json:"hostname"`
	MasterName   string `json:"master_name"`
	Password     string `json:"password"`
	Port         int    `json:"port"`
	SentinelPort int    `json:"sentinel_port"`
}

type RedisSchema RedisCredentials

func init() {
	gautocloud.RegisterConnector(NewRedisConnector())
}

type RedisConnector struct {
}

func (r RedisConnector) Id() string {
	return "hsdp:redis-db"
}

func (r RedisConnector) Name() string {
	return ".*redis-db.*"
}

func (r RedisConnector) Tags() []string {
	return []string{"Redis.*", "redis-db.*"}
}

func (r RedisConnector) Load(schema interface{}) (interface{}, error) {
	fSchema, ok := schema.(RedisSchema)
	if !ok {
		return nil, fmt.Errorf("no RedisSchema detected")
	}

	rdb := redis.NewFailoverClusterClient(&redis.FailoverOptions{
		MasterName:       fSchema.MasterName,
		SentinelAddrs:    []string{fmt.Sprintf("%s:%d", fSchema.Hostname, fSchema.SentinelPort)},
		SentinelPassword: fSchema.Password,
		Password:         fSchema.Password,
		//RouteByLatency: true,
		//RouteRandomly: true,
	})
	return rdb, nil
}

func (r RedisConnector) Schema() interface{} {
	return RedisSchema{}
}

func NewRedisConnector() connectors.Connector {
	return &RedisConnector{}
}
