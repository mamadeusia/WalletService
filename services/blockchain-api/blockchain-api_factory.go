package blockchainapi

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
)

// TODO have some implementation from blockatlas or blockbook or node ?
type Notification struct {
	Address string
	Network string
	Asset   string
	// ABI ?? another transactions or ..

}

type Notifier interface {
	Notify(ctx context.Context, address, network, asset string) error
	Listen(ctx context.Context) <-chan Notification
}

type Service struct {
	ethClient *ethclient.Client
}

func NewService(ethClient *ethclient.Client) *Service {
	return &Service{
		ethClient: ethClient,
	}
}

type BlockchainApi interface {
	Broadcast(ctx context.Context, hash common.Hash) (common.Hash, error)
	NonceOf(ctx context.Context, address string) ([]byte, error)
}
