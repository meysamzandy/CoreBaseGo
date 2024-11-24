package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"time"
)

// Create a global context to be used for Redis operations
var ctx = context.Background()

// RedisClient is a singleton instance of the Redis client, accessible throughout your Go project.
var RedisClient *redis.Client

// ConnectToRedis initializes a new Redis client and connects to the Redis server
// It takes an integer parameter `db` which specifies the Redis database to use
func ConnectToRedis() {
	// Load Redis configuration from environment variables or a configuration file (using viper for demonstration)
	redisHost := viper.GetString("REDIS_HOST")
	redisPort := viper.GetString("REDIS_PORT")
	redisUser := viper.GetString("REDIS_USER")
	redisPassword := viper.GetString("REDIS_PASS")
	redisDB := viper.GetInt("REDIS_DB") // Assuming you're using an int for DB

	if redisHost == "" || redisPort == "" {
		panic("Missing required Redis configuration: REDIS_HOST or REDIS_PORT")
	}

	// Create a Redis client
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Username: redisUser,
		Password: redisPassword,
		DB:       redisDB,
	})

	// Check connection to Redis server
	ctx := context.Background()
	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		panic(fmt.Errorf("could not connect to Redis: %w", err))
	}

}

// GetRedisClient returns the singleton Redis client instance
func GetRedisClient() *redis.Client {
	return RedisClient
}

// SetRedisStringData stores any data type in Redis under a specific key
// `value` is marshaled to JSON format before being stored
func SetRedisStringData(key string, value interface{}) error {
	rdb := GetRedisClient()
	// Marshal the value to JSON format
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("could not marshal data: %v", err)
	}

	// Set the marshaled data in Redis under the specified key with no expiration
	err = rdb.Set(ctx, key, data, 0).Err()
	if err != nil {
		return fmt.Errorf("could not set data: %v", err)
	}
	return nil
}

// GetRedisStringData retrieves data from Redis with type conversion
// The retrieved data is unmarshalled from JSON format into the `target` interface
func GetRedisStringData(key string, target interface{}) (error, error) {
	rdb := GetRedisClient()
	// Get the data stored in Redis under the specified key
	val, err := rdb.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		// Key does not exist in Redis
		return fmt.Errorf("key does not exist"), nil
	} else if err != nil {
		// An error occurred while trying to get the data from Redis
		return fmt.Errorf("could not get data: %v", err), nil
	}

	// Unmarshal the retrieved JSON data into the target interface
	err = json.Unmarshal([]byte(val), target)
	if err != nil {
		return fmt.Errorf("could not unmarshal data: %v", err), nil
	}
	return nil, nil
}

// SetRedisHashData stores any data type in Redis under a specific key using hash
// `value` is marshaled to JSON format before being stored
// `expiration` specifies the expiration duration for the Redis key
// If `expiration` is 0, no expiration is set (unlimited)
func SetRedisHashData(prefix string, key string, value interface{}, expiration time.Duration) error {
	rdb := GetRedisClient()
	key = prefix + ":" + key
	// Marshal the value to JSON format
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("could not marshal data: %v", err)
	}

	// Set the marshaled data in Redis hash under the specified key
	err = rdb.HSet(ctx, key, "data", data).Err()
	if err != nil {
		return fmt.Errorf("could not set data: %v", err)
	}

	// Set expiration for the key if expiration is greater than 0
	if expiration > 0 {
		err = rdb.Expire(ctx, key, expiration).Err()
		if err != nil {
			return fmt.Errorf("could not set expiration: %v", err)
		}
	}

	return nil
}

// GetRedisHashData retrieves data from Redis hash with type conversion
// The retrieved data is unmarshalled from JSON format into the `target` interface
func GetRedisHashData(prefix string, key string) (string, error) {
	rdb := GetRedisClient()
	key = prefix + ":" + key
	// Get the data stored in Redis hash under the specified key
	val, err := rdb.HGet(ctx, key, "data").Result()
	if errors.Is(err, redis.Nil) {
		// Key does not exist in Redis (handle as per requirement)
		return "", nil // Or return a specific error indicating missing key
	} else if err != nil {
		// An error occurred while trying to get the data from Redis
		return "", fmt.Errorf("could not get data: %v", err)
	}

	// Unmarshal the retrieved JSON data into a string (assuming code is a string)
	var data string
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		return "", fmt.Errorf("could not unmarshal data: %v", err)
	}
	return data, nil
}

// DeleteRedisHashKey remove hash keys from redis
func DeleteRedisHashKey(prefix string, key string) error {
	rdb := GetRedisClient()
	key = prefix + ":" + key
	return rdb.Del(ctx, key).Err()
}

// ExistsHashKey checks if a hash key exists in the Redis database.
func ExistsHashKey(prefix string, key string) (bool, error) {
	key = prefix + ":" + key
	rdb := GetRedisClient() // Get the singleton Redis client
	ctx := context.Background()
	val, err := rdb.HExists(ctx, key, "data").Result()
	if err != nil {
		return false, fmt.Errorf("error checking hash key %s: %w", key, err)
	}

	return val, nil
}

// ExistsHashKeyWithTTL checks if a hash key exists in the Redis database and returns its remaining TTL.
func ExistsHashKeyWithTTL(prefix string, key string) (bool, time.Duration, error) {
	key = prefix + ":" + key
	client := GetRedisClient() // Get the singleton Redis client

	ctx := context.Background()

	// Check for existence using HExists
	val, err := client.HExists(ctx, key, "data").Result()
	if err != nil {
		return false, 0, fmt.Errorf("error checking hash key %s: %w", key, err)
	}

	if !val {
		// Key doesn't exist
		return false, 0, nil
	}

	// Key exists, get TTL
	ttl, err := client.TTL(ctx, key).Result()
	if err != nil {
		// Handle potential errors retrieving TTL (e.g., key might not have an expiration set)
		return true, 0, fmt.Errorf("error getting TTL for key %s: %w", key, err)
	}

	return true, ttl, nil
}
