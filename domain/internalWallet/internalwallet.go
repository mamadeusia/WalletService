package internalWallet

import "github.com/mamadeusia/WalletSrv/entity"

type InternalWallet struct {
	Wallet *entity.Wallet
	Secret *entity.Secret
	UserId string
}
