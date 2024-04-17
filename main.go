package main

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mamadeusia/WalletSrv/client/postgres"
	"github.com/mamadeusia/WalletSrv/config"
	"github.com/mamadeusia/WalletSrv/handler"
	pb "github.com/mamadeusia/WalletSrv/proto/wallet"
	blockchainapi "github.com/mamadeusia/WalletSrv/services/blockchain-api"
	"github.com/mamadeusia/WalletSrv/services/signer"
	"github.com/mamadeusia/WalletSrv/services/wallet"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"

	eRepo "github.com/mamadeusia/WalletSrv/domain/externalWallet/postgres"
	iRepo "github.com/mamadeusia/WalletSrv/domain/internalWallet/postgres"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	blockchainpb "github.com/mamadeusia/BlockchainSrv/proto"
)

var (
	service = "walletSRV"
	version = "latest"
)

func main() {

	// Load conigurations
	if err := config.Load(); err != nil {
		logger.Fatal("MAINROUTINE::LoadConfig:: has failed with error : %+v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	// Create service
	srv := micro.NewService(
		micro.Context(ctx),
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
		micro.BeforeStop(func() error {
			logger.Infof("shutting down %s service", service)
			cancel()
			return nil
		}),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	ethClient, err := ethclient.Dial(config.NodeURL())
	if err != nil {
		logger.Fatal(err)
	}

	repo, err := postgres.NewPostgres(context.Background(), config.PostgresURL())
	if err != nil {
		logger.Fatal(err)
	}

	externalWalletRepo := eRepo.NewRepository(repo)
	internalWallet := iRepo.NewRepository(repo)

	blockchainApiInstance := blockchainapi.NewService(ethClient)

	blockchainSrv := blockchainpb.NewBlockchainSrvService(config.GetBlockchainServiceName(), srv.Client())

	walletService := wallet.NewService(externalWalletRepo, internalWallet, signer.NewService(blockchainSrv, blockchainApiInstance), config.NodeMnemonic())

	// Register handler
	if err := pb.RegisterWalletSrvHandler(srv.Server(), handler.NewWalletHandler(walletService)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
