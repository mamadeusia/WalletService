package blockchainapi

import (
	"context"
	"encoding/binary"

	"github.com/ethereum/go-ethereum/common"
	"go-micro.dev/v4/errors"
)

func (s *Service) NonceOf(ctx context.Context, address string) ([]byte, error) {
	nonce, err := s.ethClient.PendingNonceAt(ctx, common.HexToAddress(address))
	if err != nil {
		return nil, errors.InternalServerError("SERVICE::NonceOf", "failed to get pending nonce: %v", err)
	}

	var nonceBytes []byte
	binary.LittleEndian.PutUint64(nonceBytes, uint64(nonce))

	return nonceBytes, nil
}

func (s *Service) Broadcast(ctx context.Context, hash common.Hash) (common.Hash, error) {
	tx, _, err := s.ethClient.TransactionByHash(ctx, hash)
	if err != nil {
		return common.Hash{}, errors.InternalServerError("SERVICE::Broadcast", "failed to make transaction by hash: %v", err)
	}

	return tx.Hash(), nil
}
