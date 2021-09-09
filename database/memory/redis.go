package memory

import (
	"github.com/go-redis/redis"
)
func RedisCnn() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	return client
}