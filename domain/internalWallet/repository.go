package internalWallet

import (
	"context"

	"github.com/mamadeusia/WalletSrv/entity"
)

// each user one account for specific network and asset
type Repository interface {
	SetUserNewWallet(ctx context.Context, internalWallet *InternalWallet) error
	GetAllUserWallets(ctx context.Context, userID string) ([]InternalWallet, error)
	GetUserWallet(ctx context.Context, userID string, network entity.Network, asset entity.Asset) (*InternalWallet, error)
}
