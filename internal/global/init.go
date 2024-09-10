package global

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

var ctx = context.Background()

func InitLog() {
	var log = logrus.New()

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(logrus.WarnLevel)

	Log = log
}

func InitDatabase() {
	// init database

	// db, err := gorm.Open(postgres.New(postgres.Config{
	// 	DSN:                  "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
	// 	PreferSimpleProtocol: true,                                                                                                  // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	// }), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Db = db
}

func InitRedis() {
	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379", // use the correct address for your configuration
	// 	Password: "",               // no password set
	// 	DB:       0,                // use default DB
	// })

	// // Check the connection
	// pong, err := redisClient.Ping(ctx).Result()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(pong) // Outputs: PONG
	// }
	// RedisClient = redisClient
}
