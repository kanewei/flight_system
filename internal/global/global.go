package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Log         *logrus.Logger
	Db          *gorm.DB
	RedisClient *redis.Client
)

func Shutdown() {

}
