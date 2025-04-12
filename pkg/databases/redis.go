package databases

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// NewRedisClient создает новый клиент Redis
func Redis(params Dependencies) (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", params.Config.Redis.Host, params.Config.Redis.Port), // Адрес Redis сервера
		Password: "",                                                                       // Пароль (если есть)
		DB:       0,                                                                        // Номер базы данных
	})

	// Проверяем подключение

	if err := client.Ping(context.Background()).Err(); err != nil {
		panic("s")
	}

	fmt.Println("Successfully connected to Redis!")
	return
}
