package utils

import (
	"errors"

	"github.com/mamadeusia/WalletSrv/pkg/crypto"

	"github.com/mamadeusia/WalletSrv/entity"
)

func ConvertNetworkToTWCoinType(network entity.Network) (crypto.CoinType, error) {
	//TODO:omid: ask about this
	switch network {
	case entity.BSC:
		return crypto.CoinTypeEthereum, nil
	default:
		return 0, errors.New("unkown coin type")
	}
}
