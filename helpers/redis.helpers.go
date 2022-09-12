package helpers

import (
	"errors"

	"github.com/go-redis/redis"
)

var redisCli *redis.Client

func InitiateRDS(rds *redis.Client) {
	redisCli = rds
}

func EnCache(key string, data []byte) error {
	res, err := redisCli.SetNX(key, data, 0).Result()
	if err != nil || !res {
		if err == nil {
			err = errors.New("encaching failed")
		}
		return err
	}
	return nil
}

func GetCachedValue(key string) (value []byte, err error) {
	value, err = redisCli.Get(key).Bytes()
	if err != nil {
		return nil, err
	}
	return value, nil
}
