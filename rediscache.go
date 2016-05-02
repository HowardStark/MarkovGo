package main

import (
	"errors"
	"math/rand"

	"gopkg.in/redis.v3"
)

// RedisCache is a caching module
// that integrates with a redis
// database for its backend.
type RedisCache struct {
	address    string
	password   string
	db         int64
	connection *redis.Client
	status     RedisStatus
}

// RedisStatus is the enum type to represent the different
// states/statuses the redis client could have
type RedisStatus int

const (
	// RedisConnected indicates that the redis client
	// is connected to the server
	RedisConnected RedisStatus = iota
	// RedisDisconnected indicates that the redis client
	// cannot reach the server or the connection has
	// been interrupted.
	RedisDisconnected
	// RedisDefaultPort is the default port the redis
	// server listens on.
	RedisDefaultPort string = "6379"
)

// NewRedisCache builds a new redis cache from an
// address, a password, and a database.
func NewRedisCache(addr, pass string, db int) (*RedisCache, error) {
	redisCache := &RedisCache{
		address:    addr,
		password:   pass,
		db:         int64(db),
		connection: nil,
		status:     RedisDisconnected,
	}
	err := redisCache.Connect()
	if err != nil {
		return redisCache, err
	}
	return redisCache, nil
}

// Connect initializes the RedisCache
// using the credentials in the current
// instance.
func (redisCache *RedisCache) Connect() error {
	redisCache.connection = redis.NewClient(&redis.Options{
		Addr:     redisCache.address,
		Password: redisCache.password,
		DB:       redisCache.db,
	})
	_, err := redisCache.connection.Ping().Result()
	if err != nil {
		redisCache.status = RedisDisconnected
		return err
	}
	return nil
}

// AddPair adds a new potential value to the passed key
func (redisCache *RedisCache) AddPair(key string, value string) error {
	if redisCache.status != RedisDisconnected {
		err := redisCache.connection.LPush(key, value).Err()
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("rediscache: redis client isn't connected")
}

// GetRandom returns a random value from the given key
func (redisCache *RedisCache) GetRandom(key string) (string, error) {
	if redisCache.status != RedisDisconnected {
		len, lenErr := redisCache.connection.LLen(key).Result()
		if lenErr != nil {
			return "", lenErr
		}
		if len <= 0 {
			return "", errors.New("rediscache: length was less than or equal to 0")
		}
		index := rand.Int63n(len - 1)
		value, valueErr := redisCache.connection.LIndex(key, index).Result()
		if valueErr != nil {
			return "", valueErr
		}
		return value, nil
	}
	return "", errors.New("rediscache: redis client isn't connected")
}
