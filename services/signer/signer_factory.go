package signer

import (
	blockchainpb "github.com/mamadeusia/BlockchainSrv/proto"
	"github.com/mamadeusia/WalletSrv/entity"
	blockchainapi "github.com/mamadeusia/WalletSrv/services/blockchain-api"
)

type PrivateKey string
type TXID string

type Service struct {
	blockchainService blockchainpb.BlockchainSrvService
	blockchainApi     blockchainapi.BlockchainApi
}

func NewService(b blockchainpb.BlockchainSrvService, blockchainApi blockchainapi.BlockchainApi) *Service {
	return &Service{
		blockchainService: b,
		blockchainApi:     blockchainApi,
	}
}

type Signer interface {
	SignVoteTransaction(wallet entity.Wallet, secret entity.Secret, vote entity.Vote) (TXID, error)
}
