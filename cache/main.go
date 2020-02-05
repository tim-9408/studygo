package cache 

import (
	"os"
	"strconv"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client 

func Redis()  {
	strconv.ParseUnit(os.Getenv("REDIS_DB"),10,64)
	va
}