package databases

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

// NewRedisClient создает новый клиент Redis
func Redis(params Dependencies) (client *redis.Client) {
	opt, err := redis.ParseURL(params.Config.Redis.URL)
	if err != nil {
		panic("err while trying to parse url")
	}

	client = redis.NewClient(opt)

	// Проверяем подключение

	// if err := client.Ping(context.Background()).Err(); err != nil {
	// 	params.Logger.Warn().Err(err).Msg("An error occurred while trying to ping redis host")
	// }

	fmt.Println("Successfully connected to Redis!")
	return
}
