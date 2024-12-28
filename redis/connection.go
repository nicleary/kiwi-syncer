package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

var redisClient *redis.Client

func InitializeRedisClient() error {
	// Get host and port together
	redisHost := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))

	// Get redis DB number from environment variable
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	if err != nil {
		return err
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDB,
	})

	return nil
}

func GetRedisClient() *redis.Client {
	return redisClient
}
