package handler

import (
	"context"
	"errors"
	"math/rand"

	"github.com/google/uuid"
	"github.com/mamadeusia/WalletSrv/domain/externalWallet"
	"github.com/mamadeusia/WalletSrv/domain/internalWallet"
	"github.com/mamadeusia/WalletSrv/entity"
	pb "github.com/mamadeusia/WalletSrv/proto/wallet"
	"github.com/mamadeusia/WalletSrv/services/wallet"
	"go-micro.dev/v4/logger"
)

type WalletHandler struct {
	service *wallet.Service
}

func NewWalletHandler(ws *wallet.Service) *WalletHandler {
	return &WalletHandler{
		service: ws,
	}
}

func (e *WalletHandler) CreateNewInternalWallet(ctx context.Context, req *pb.CreateNewInternalWalletRequest, rsp *pb.CreateNewInternalWalletResponse) error {
	//todo: mobin jan
	if req.GetAsset().Enum() == nil || req.GetNetwork().Enum() == nil || req.UserId == "" {
		logger.Errorf("HANDLER::CreateNewInternalWallet, has failed with error : invalid input type : %v", req)
		return errors.New("invalid input type")
	}

	//todo: here we have to check for networks but now we have just bsc and only usdc so i put it constant
	wallet, err := e.service.CreateInternalWallet(ctx, internalWallet.InternalWallet{
		Wallet: &entity.Wallet{
			ID:      uuid.New(),
			Network: entity.BSC,
			Asset:   entity.USDC,
		},
		Secret: &entity.Secret{
			DerivationVersion: entity.BSC_V1,
			Index:             rand.Uint32(),
		},
		UserId: req.UserId,
	})

	if err != nil {
		logger.Errorf("HANDLER::CreateNewInternalWallet, has failed with error : %v", err)
		return err
	}
	rsp.Address = wallet.Address
	return nil
}

func (e *WalletHandler) AddExternalWallet(ctx context.Context, req *pb.AddExternalWalletRequest, rsp *pb.AddExternalWalletResponse) error {
	//todo:: mobin jan
	if req.Address == "" || req.Asset.Enum() == nil || req.Network.Enum() == nil || req.UserId == "" {
		logger.Errorf("HANDLER::AddExternalWallet, has failed with error : invalid input type, %v", req)
		return errors.New("invlaid input type")
	}
	//todo:: we have to check the asset and network but right now we just have usdc on bsc
	if err := e.service.AddExternalWallet(ctx, externalWallet.ExternalWallet{
		Wallet: &entity.Wallet{
			ID:      uuid.New(),
			Address: req.Address,
			Network: entity.BSC,
			Asset:   entity.USDC,
		},
		UserId: req.UserId,
	}); err != nil {
		logger.Errorf("HANDLER::AddExternalWallet, has failed with error : %v", err)
		return err
	}
	return nil
}

func (e *WalletHandler) InternalVoteOnProposal(ctx context.Context, req *pb.InternalVoteOnProposalRequest, rsp *pb.InternalVoteOnProposalResponse) error {
	//todo: mobin jan
	if req.ProposalId == "" || req.Type.Enum() == nil || req.UserId == "" {
		logger.Errorf("HANDLER::InternalVoteOnProposal, has failed with error: invalid input type , %v", req)
		return nil
	}

	txId, err := e.service.InternalVoteOnProposal(ctx, entity.Vote{
		UserId:     req.UserId,
		Type:       entity.VoteType(req.Type),
		ProposalId: req.ProposalId,
	})

	if err != nil {
		logger.Errorf("HANDLER::InternalVoteOnProposal, has failed with error, invalid input type, %v", req)
		return err
	}
	rsp.TxId = string(txId)
	return nil
}
func (e *WalletHandler) ExternalVoteOnProposal(ctx context.Context, req *pb.ExternalVoteOnProposalRequest, rsp *pb.ExternalVoteOnProposalResponse) error {
	//todo : mobin jan
	if req.ProposalId == "" || req.Type.Enum() == nil || req.UserId == "" {
		logger.Errorf("HANDLER::ExternalVoteOnProposal, has failed with error: invalid input type , %v", req)
		return nil
	}

	deepLink, err := e.service.ExternalVoteOnProposal(ctx, entity.Vote{
		UserId:     req.UserId,
		Type:       entity.VoteType(req.Type),
		ProposalId: req.ProposalId,
	})

	if err != nil {
		logger.Errorf("HANDLER::ExternalVoteOnProposal, has failed with error : %v", err)
		return err
	}
	rsp.DeeplinkUrl = deepLink
	return nil
}

// CreateRecieptQrcode implements WalletSrv.WalletSrvHandler
func (*WalletHandler) CreateRecieptQrcode(context.Context, *pb.CreateRecieptQrcodeRequest, *pb.CreateRecieptQrcodeResponse) error {
	panic("unimplemented")
}

// Transfer implements WalletSrv.WalletSrvHandler
func (*WalletHandler) Transfer(context.Context, *pb.TransferRequest, *pb.TransferResponse) error {
	panic("unimplemented")
}
