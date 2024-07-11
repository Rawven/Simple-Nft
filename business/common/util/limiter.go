package util

import (
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/redis/go-redis/v9"
	"io"
)

var limitScript *redis.Script

func InitLimiter() {
	file, f, err := fileutil.ReadFile("../script/redis_limiter.lua")
	if err != nil {
		panic(err)
	}
	defer f()
	dat, err := io.ReadAll(file)
	if err != nil {
		return
	}
	limitScript = redis.NewScript(string(dat))
}

func GetLimiter() *redis.Script {
	return limitScript
}
