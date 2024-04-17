package externalWallet

import (
	"context"

	"github.com/mamadeusia/WalletSrv/entity"
)

type Repository interface {
	SetUserNewWallet(ctx context.Context, w ExternalWallet) error
	GetAllUserWallets(ctx context.Context, userID string) ([]ExternalWallet, error)
	GetUserWallet(ctx context.Context, userID string, network entity.Network, asset entity.Asset) (*ExternalWallet, error)
}
