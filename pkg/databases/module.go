package databases

import (
	"database/sql"

	"github.com/WayDBae/eWallet/pkg/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

// OracleModule ...
// var OracleModule = fx.Provide(NewOracleConn)

var Module = fx.Options(PostgresModule, RedisModule)

// PostgresModule ...
var PostgresModule = fx.Provide(NewPostgresConn)

// RedisModule ...
var RedisModule = fx.Provide(NewRedisConn)

// Dependencies ...
type Dependencies struct {
	fx.In

	Logger zerolog.Logger
	Config *config.Config
}

// NewOracleConn ...
func NewOracleConn(params Dependencies) *sql.DB {
	return Oracle(params)
}

// NewPostgresConn ...
func NewPostgresConn(params Dependencies) *gorm.DB {
	return Postgres(params)
}

func NewRedisConn(params Dependencies) (client *redis.Client) {
	return Redis(params)
}
