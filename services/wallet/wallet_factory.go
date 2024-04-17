package wallet

import (
	"context"

	"github.com/mamadeusia/WalletSrv/domain/externalWallet"
	epostgres "github.com/mamadeusia/WalletSrv/domain/externalWallet/postgres"
	"github.com/mamadeusia/WalletSrv/domain/internalWallet"
	ipostgres "github.com/mamadeusia/WalletSrv/domain/internalWallet/postgres"
	"github.com/mamadeusia/WalletSrv/entity"
	"github.com/mamadeusia/WalletSrv/services/signer"
)

type TXID string

type Service struct {
	exRepo   *epostgres.PostgresRepository
	inRepo   *ipostgres.PostgresRepository
	signer   *signer.Service
	mnemonic string
}

func NewService(e *epostgres.PostgresRepository, i *ipostgres.PostgresRepository, s *signer.Service, m string) *Service {
	return &Service{
		exRepo:   e,
		inRepo:   i,
		signer:   s,
		mnemonic: m,
	}
}

type Wallet interface {
	CreateInternalWallet(ctx context.Context, internalWallet internalWallet.InternalWallet) (*entity.Wallet, error)
	AddExternalWallet(ctx context.Context, externalWallet externalWallet.ExternalWallet) error
	InternalVoteOnProposal(ctx context.Context, vote entity.Vote) (TXID, error)
	ExternalVoteOnProposal(ctx context.Context, vote entity.Vote) (string, error)
}
