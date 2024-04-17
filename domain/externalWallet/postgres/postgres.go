package postgres

import (
	"context"

	"github.com/mamadeusia/WalletSrv/client/postgres"
	"github.com/mamadeusia/WalletSrv/domain/externalWallet"
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

func (p *PostgresRepository) SetUserNewWallet(ctx context.Context, w externalWallet.ExternalWallet) error {
	if err := p.Postgres.Queries.CreateWallet(ctx, postgres.CreateWalletParams{
		ID:      w.Wallet.ID,
		UserID:  w.UserId,
		Address: w.Wallet.Address,
		Network: string(w.Wallet.Network),
		Asset:   string(w.Wallet.Asset),
	}); err != nil {
		logger.Info("DOMAIN::SetUserNewWallet, has failed with error : %v", err)
		return err
	}
	return nil
}

// todo:: find where we need this then implement
func (p *PostgresRepository) GetAllUserWallets(ctx context.Context, userID string) ([]externalWallet.ExternalWallet, error) {
	wallets, err := p.Postgres.Queries.GetAllWallet(ctx, postgres.GetAllWalletParams{
		UserID: userID,
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		logger.Info("DOMAIN::GetAllUserWallets, has failed with error : %v", err)
		return nil, err
	}
	var ewallets []externalWallet.ExternalWallet

	for _, wallet := range wallets {
		ewallets = append(ewallets, externalWallet.ExternalWallet{
			Wallet: &entity.Wallet{
				ID:      wallet.ID,
				Address: wallet.Address,
				Network: entity.Network(wallet.Network),
				Asset:   entity.Asset(wallet.Asset),
			},
			UserId: userID,
		})
	}
	return ewallets, nil
}
func (p *PostgresRepository) GetUserWallet(ctx context.Context, userID string, network entity.Network, asset entity.Asset) (*externalWallet.ExternalWallet, error) {
	result, err := p.Postgres.Queries.GetWallet(ctx, postgres.GetWalletParams{
		UserID:  userID,
		Network: string(network),
		Asset:   string(asset),
	})
	if err != nil {
		logger.Info("DOMAIN::GetUserWallet, has failed with error : %v", err)
		return nil, err
	}

	return &externalWallet.ExternalWallet{
		Wallet: &entity.Wallet{
			ID:      result.ID.UUID,
			Address: result.Address.String,
			Network: entity.Network(result.Network.String),
			Asset:   entity.Asset(result.Asset.String),
		},
		UserId: result.UserID.String,
	}, nil
}
