package storage

import (
	"github.com/WayDBae/eWallet/internal/storage/currency"
	"github.com/WayDBae/eWallet/internal/storage/rdb"
	"github.com/WayDBae/eWallet/internal/storage/user"
	"github.com/WayDBae/eWallet/internal/storage/wallet"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	user.Module,
	rdb.Module,
	wallet.Module,
	currency.Module,
)
