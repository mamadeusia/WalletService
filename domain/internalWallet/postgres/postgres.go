package postgres

import (
	"context"

	"github.com/mamadeusia/WalletSrv/client/postgres"
	"github.com/mamadeusia/WalletSrv/domain/internalWallet"
	"github.com/mamadeusia/WalletSrv/entity"
	"go-micro.dev/v4/logger"
)

type PostgresRepository struct {
	Postgres *postgres.PostgresRepository
}

func NewRepository(postgres *postgres.PostgresRepository) *PostgresRepository {
	return &PostgresRepository{
		Postgres: postgres,
	}
}

// todo:: some safe methods for storing unit32
func (p *PostgresRepository) SetUserNewWallet(ctx context.Context, internalWallet *internalWallet.InternalWallet) error {
	tx, err := p.Postgres.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	qtx := p.Postgres.Queries.WithTx(tx)

	if err := qtx.CreateWallet(ctx, postgres.CreateWalletParams{
		ID:      internalWallet.Wallet.ID,
		UserID:  internalWallet.UserId,
		Address: internalWallet.Wallet.Address,
		Network: string(internalWallet.Wallet.Network),
		Asset:   string(internalWallet.Wallet.Asset),
	}); err != nil {
		logger.Info("DOMAIN::SetUserNewWallet, has failed with error : %v", err)
		return err
	}
	if err := qtx.CreateInternalWallet(ctx, postgres.CreateInternalWalletParams{
		WalletID:          internalWallet.Wallet.ID,
		DerivationVersion: string(internalWallet.Secret.DerivationVersion),
		Index:             int64(internalWallet.Secret.Index),
	}); err != nil {
		logger.Info("DOMAIN::SetUserNewWallet, has failed with error : %v", err)
	}
	return tx.Commit(ctx)
}

// todo:: find where we need this then implement
func (p *PostgresRepository) GetAllUserWallets(ctx context.Context, userID string) ([]internalWallet.InternalWallet, error) {
	return nil, nil
}

func (p *PostgresRepository) GetUserWallet(ctx context.Context, userID string, network entity.Network, asset entity.Asset) (*internalWallet.InternalWallet, error) {
	result, err := p.Postgres.Queries.GetWallet(ctx, postgres.GetWalletParams{
		UserID:  userID,
		Network: string(network),
		Asset:   string(asset),
	})
	if err != nil {
		logger.Info("DOMAIN::GetUserWallet, has failed with error : %v", err)
		return nil, err
	}
	return &internalWallet.InternalWallet{
		Wallet: &entity.Wallet{ID: result.WalletID, Address: result.Address.String, Network: entity.Network(result.Network.String), Asset: entity.Asset(result.Asset.String)},
		Secret: &entity.Secret{
			DerivationVersion: entity.DerivationVersion(result.DerivationVersion),
			Index:             uint32(result.Index),
		},
		UserId: result.UserID.String,
	}, nil
}
