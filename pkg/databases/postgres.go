package databases

import (
	"fmt"

	"github.com/WayDBae/eWallet/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Postgres ...
func Postgres(params Dependencies) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Dushanbe",

		params.Config.Postgres.Host,
		params.Config.Postgres.Username,
		params.Config.Postgres.Password,
		params.Config.Postgres.DatabaseName,
		fmt.Sprint(params.Config.Postgres.Port),
		params.Config.Postgres.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // можно выбрать Silent, Error, Warn, Info
	})
	if err != nil {
		panic(err)
	}

	// Автоматическая миграция моделей
	if err := db.AutoMigrate(
		&entities.User{},
		&entities.Wallet{},
		&entities.Currency{},
	); err != nil {
		panic(err)
	}

	return db
}
