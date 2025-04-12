package storage

import (
	"github.com/WayDBae/eWallet/internal/storage/rdb"
	"github.com/WayDBae/eWallet/internal/storage/user"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	user.Module,
	rdb.Module,
)
