package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Redis struct {
	Client redis.Cmdable
}

func (r *Redis) SetPrefix(string) {
}

func (r *Redis) Connect() error {
	return nil
}

// Get from key
func (r *Redis) Get(key string) (string, error) {
	return r.Client.Get(r.ctx(), key).Result()
}

// Set value with key and expire time
func (r *Redis) Set(key string, val interface{}, expire int) error {
	return r.Client.Set(r.ctx(), key, val, time.Duration(expire)*time.Second).Err()
}

// Del delete key in redis
func (r *Redis) Del(key string) error {
	return r.Client.Del(r.ctx(), key).Err()
}

// HashGet from key
func (r *Redis) HashGet(hk, key string) (string, error) {
	return r.Client.HGet(r.ctx(), hk, key).Result()
}

// HashDel delete key in specify redis's hashtable
func (r *Redis) HashDel(hk, key string) error {
	return r.Client.HDel(r.ctx(), hk, key).Err()
}

// Increase
func (r *Redis) Increase(key string) error {
	return r.Client.Incr(r.ctx(), key).Err()
}

func (r *Redis) Decrease(key string) error {
	return r.Client.Decr(r.ctx(), key).Err()
}

// Set ttl
func (r *Redis) Expire(key string, dur time.Duration) error {
	return r.Client.Expire(r.ctx(), key, dur).Err()
}

// 获取上下文
func (r *Redis) ctx() context.Context {
	return context.Background()
}
