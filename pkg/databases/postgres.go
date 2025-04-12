package databases

import (
	"fmt"
	"log"

	"github.com/WayDBae/eWallet/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Postgres ...
func Postgres(params Dependencies) (con *gorm.DB) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Dushanbe",

		params.Config.Postgres.Host,
		params.Config.Postgres.Username,
		params.Config.Postgres.Password,
		params.Config.Postgres.DatabaseName,
		params.Config.Postgres.SSLMode,
		fmt.Sprint(params.Config.Postgres.Port),
	)

	con, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default})
	if err != nil {
		log.Println(err)
		// Надо?
		return
	}

	err = con.AutoMigrate(&entities.User{})
	if err != nil {
		panic(err)
	}
	return
}
