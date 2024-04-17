package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Wallet struct {
	ID      uuid.UUID
	Address string
	Network Network
	Asset   Asset
}

func ExtractChainIDFromNetwork(network Network) (int64, error) {
	switch network {
	case BSC:
		return 8001, nil
	default:
		return 0, errors.New("network is not supported")
	}
}

//TODO :: error package
