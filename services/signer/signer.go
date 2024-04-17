package signer

// #cgo CFLAGS: -I../../trustwallet/include
// #cgo LDFLAGS: -L../../trustwallet/build -L../../trustwallet/build/local/lib -L../../trustwallet/build/trezor-crypto -lTrustWalletCore -lwallet_core_rs -lprotobuf -lTrezorCrypto -lstdc++ -lm
// #include <TrustWalletCore/TWHDWallet.h>
// #include <TrustWalletCore/TWPrivateKey.h>
// #include <TrustWalletCore/TWPublicKey.h>
// #include <TrustWalletCore/TWEthereumAbi.h>
// #include <TrustWalletCore/TWEthereumAbiFunction.h>
// #include <TrustWalletCore/TWData.h>
import "C"

import (
	"context"
	"encoding/binary"
	"time"

	"github.com/ethereum/go-ethereum/common"
	blockchainpb "github.com/mamadeusia/BlockchainSrv/proto"
	"github.com/mamadeusia/WalletSrv/config"
	"github.com/mamadeusia/WalletSrv/entity"
	"github.com/mamadeusia/WalletSrv/pkg/crypto"
	"github.com/mamadeusia/WalletSrv/pkg/types"
	"github.com/mamadeusia/WalletSrv/pkg/utils"
	"github.com/mamadeusia/WalletSrv/proto/tw/ethereum"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/logger"
)

func (s *Service) SignVoteTransaction(ctx context.Context, wallet entity.Wallet, secret entity.Secret, vote entity.Vote) (TXID, error) {

	ct, err := utils.ConvertNetworkToTWCoinType(wallet.Network)
	if err != nil {
		logger.Info("SERVICE::SignVoteTransaction, has failed with error : %v", err)
		return "", err
	}

	secret.Mnemunic = config.NodeMnemonic()

	userWallet, err := crypto.CreateWalletWithMnemonic(secret.Mnemunic, ct)
	if err != nil {
		logger.Info("SERVICE::SignVoteTransaction, has failed with error : %v", err)
		return "", err
	}

	//castVote(uint256 proposalId, uint8 support, bytes32[] calldata proof)
	//call blockchain service to get proofs
	proofs, err := s.blockchainService.Proofs(ctx, &blockchainpb.ProofsRequest{
		VoterAddress: userWallet.Address,
		ProposeID:    vote.ProposalId,
	}, client.WithDialTimeout(30*time.Second))
	if err != nil {
		logger.Info("SERVICE::SignVoteTransaction, has failed with error : %v", err)
		return "", err
	}
	dd := "castVote"

	f := C.TWEthereumAbiFunctionCreateWithString(C.TWStringCreateWithUTF8Bytes(C.CString(dd)))
	C.TWEthereumAbiFunctionAddParamUInt256(f, C.TWDataCreateWithHexString(C.TWStringCreateWithUTF8Bytes(C.CString(vote.ProposalId))), false)
	C.TWEthereumAbiFunctionAddParamUInt8(f, C.uchar(int(vote.Type)), false)
	for index, val := range proofs.Proofs {
		C.TWEthereumAbiFunctionAddInArrayParamBytes(f, C.int(index), types.TWDataCreateWithGoBytes(val))
	}

	chainID, err := entity.ExtractChainIDFromNetwork(wallet.Network)
	if err != nil {
		logger.Info("SERVICE::SignVoteTransaction, has failed with error : %v", err)
	}
	bChainID := make([]byte, 8)
	binary.LittleEndian.PutUint64(bChainID, uint64(chainID))

	nonce, err := s.blockchainApi.NonceOf(ctx, userWallet.Address)
	if err != nil {
		logger.Info("SERVICE::SignVoteTransaction, has failed with error : %v", err)
		return "", err
	}

	txInputData := ethereum.SigningInput{
		ChainId:               bChainID,
		Nonce:                 nonce,
		TxMode:                ethereum.TransactionMode_Enveloped,
		MaxInclusionFeePerGas: []byte{},
		MaxFeePerGas:          []byte{},
		ToAddress:             config.Get().ContractAddress,
		PrivateKey:            []byte(userWallet.PriKey),
		Transaction: &ethereum.Transaction{
			TransactionOneof: &ethereum.Transaction_ContractGeneric_{
				ContractGeneric: &ethereum.Transaction_ContractGeneric{
					Amount: []byte{},
					Data:   types.TWDataGoBytes(C.TWEthereumAbiEncode(f)),
				},
			},
		},
	}

	var output ethereum.SigningOutput

	if err := crypto.CreateSignedTx(&txInputData, ct, &output); err != nil {
		logger.Info("SERVICE::SignVoteTransaction, has failed with error : %v", err)
		return "", err
	}
	if output.Error != 0 {
		logger.Info("SERVICE::SignVoteTransaction, has failed with error  %v", output.ErrorMessage)
	}

	txHash, err := s.blockchainApi.Broadcast(ctx, common.BytesToHash(output.Data))
	if err != nil {
		logger.Info("SERVICE::SignVoteTransaction, has failed with error : %v", err)
		return "", err
	}

	return TXID(txHash.String()), nil
}
