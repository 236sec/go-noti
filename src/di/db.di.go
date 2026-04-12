package di

import (
	"sync"

	"goboilerplate.com/src/pkg/database"
	"goboilerplate.com/src/pkg/redis"
)

var GetDB = sync.OnceValue(func() database.IDatabase {
	return database.GetDatabase()
})

var GetRedis = sync.OnceValue(func() redis.IRedis {
	return redis.GetRedisClient()
})
