package store

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

// Const inits
const (
	CacheDuration = 6 * time.Hour
)

// Struct wrapper around Redis client
type StorageService struct {
	redisClient *redis.Client
}

// Declarations inits
var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

// Initialize the store service and return a store pointer
func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_DB_ADDRESS"),
		Password: os.Getenv("REDIS_DB_PASSWORD"),
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

// SaveUrlMapping saves the short URL and original URL in the Redis cache
func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}

}

// RetrieveInitialUrl retrieves the original URL from the Redis cache, given the shortened URL
func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}
