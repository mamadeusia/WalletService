package wallet

import (
	"context"

	"github.com/mamadeusia/WalletSrv/domain/externalWallet"
	"github.com/mamadeusia/WalletSrv/domain/internalWallet"
	"github.com/mamadeusia/WalletSrv/entity"

	"github.com/mamadeusia/WalletSrv/pkg/crypto"
	"github.com/mamadeusia/WalletSrv/pkg/utils"
	"go-micro.dev/v4/logger"
)

func (s *Service) CreateInternalWallet(ctx context.Context, internalWallet internalWallet.InternalWallet) (*entity.Wallet, error) {
	ct, err := utils.ConvertNetworkToTWCoinType(internalWallet.Wallet.Network)
	if err != nil {
		logger.Info("SERVICE::CreateInternalWallet, has failed with error : %v", err)
		return nil, err
	}

	//todo:: fetching mnemunic from vault:(((((())))))
	wallet, err := crypto.CreateWalletWithMnemonic(s.mnemonic, ct)
	if err != nil {
		logger.Info("SERVICE::CreateInternalWallet, has failed with error : %v", err)
		return nil, err
	}
	internalWallet.Wallet.Address = wallet.Address
	s.inRepo.SetUserNewWallet(ctx, &internalWallet)
	return internalWallet.Wallet, nil
}

func (s *Service) AddExternalWallet(ctx context.Context, externalWallet externalWallet.ExternalWallet) error {
	if err := s.exRepo.SetUserNewWallet(ctx, externalWallet); err != nil {
		logger.Info("SERVICE::AddExternalWallet, has failed with error : %v", err)
		return err
	}
	return nil
}

func (s *Service) InternalVoteOnProposal(ctx context.Context, vote entity.Vote) (TXID, error) {
	//todo:: after we add other networks and asset we should fix this shiiit
	wallet, err := s.inRepo.GetUserWallet(ctx, vote.UserId, entity.BSC, entity.USDC)
	if err != nil {
		logger.Info("SERVICE::InternalVoteOnProposal, has faild with error : %v", err)
		return "", err
	}
	txID, err := s.signer.SignVoteTransaction(ctx, *wallet.Wallet, *wallet.Secret, vote)
	if err != nil {
		logger.Info("SERVICE::InternalVoteOnProposal, has failed with error , %v", err)
		return "", err
	}
	return TXID(txID), nil
}

func (s *Service) ExternalVoteOnProposal(ctx context.Context, vote entity.Vote) (string, error) {
	//todo:: first dapp should be implemented
	return "", nil
}
