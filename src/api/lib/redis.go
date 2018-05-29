package lib

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

type RedisLib struct {
}

func (r RedisLib) SetToken(uid string, expire int) interface{} {

	a := time.Duration(expire) * time.Second

	redis, err := cache.NewCache("redis", `{"key":"user","conn":":6379","dbNum":"0"}`)
	if err == nil {
		// 获取当前Unix时间戳
		t := time.Now()
		timeUnix := t.Unix()
		// 由于获取的Unix 时间戳是int64位，需要转换为string类型 以供设置redis 缓存
		timeUnixForString := strconv.FormatInt(timeUnix, 10)

		// 将时间戳转换为md5
		md5 := md5.Sum([]byte(timeUnixForString))
		token := fmt.Sprintf("%x", md5)

		// 设置缓存

		redis.Put(uid, token, a)
		return token
	}
	return false
}
