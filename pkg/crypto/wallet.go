package crypto

// #cgo CFLAGS: -I../../trustwallet/include
// #cgo LDFLAGS: -L../../trustwallet/build -L../../trustwallet/build/local/lib -L../../trustwallet/build/trezor-crypto -lTrustWalletCore -lwallet_core_rs -lprotobuf -lTrezorCrypto -lstdc++ -lm
// #include <TrustWalletCore/TWHDWallet.h>
// #include <TrustWalletCore/TWPrivateKey.h>
// #include <TrustWalletCore/TWPublicKey.h>
import "C"

import (
	"errors"
	"fmt"

	"github.com/mamadeusia/WalletSrv/pkg/types"
)

type Wallet struct {
	Address string
	PriKey  string
	PubKey  string
	CoinType
}

func CreateWalletWithMnemonic(mn string, ct CoinType) (*Wallet, error) {
	if !IsMnemonicValid(mn) {
		return nil, errors.New("mnemonic is not valid")
	}

	str := types.TWStringCreateWithGoString(mn)
	empty := types.TWStringCreateWithGoString("")
	defer C.TWStringDelete(str)
	defer C.TWStringDelete(empty)

	tw := C.TWHDWalletCreateWithMnemonic(str, empty)

	defer C.TWHDWalletDelete(tw)

	priKey := C.TWHDWalletGetKeyForCoin(tw, C.enum_TWCoinType(ct))
	defer C.TWPrivateKeyDelete(priKey)
	priKeyData := C.TWPrivateKeyData(priKey)
	defer C.TWDataDelete(priKeyData)

	pubKey := C.TWPrivateKeyGetPublicKeySecp256k1(priKey, true)
	defer C.TWPublicKeyDelete(pubKey)
	pubKeyData := C.TWPublicKeyData(pubKey)
	defer C.TWDataDelete(pubKeyData)

	address := C.TWHDWalletGetAddressForCoin(tw, C.enum_TWCoinType(ct))
	defer C.TWStringDelete(address)

	return &Wallet{
		Address:  types.TWStringGoString(address),
		PriKey:   types.TWDataHexString(priKeyData),
		PubKey:   types.TWDataHexString(pubKeyData),
		CoinType: ct,
	}, nil
}

func FetchPrivateKeyFromSeed(entropy []byte, passphrase string, ct CoinType) string {
	pass := types.TWStringCreateWithGoString(passphrase)
	defer C.TWStringDelete(pass)
	fmt.Println(pass)
	fmt.Println("*****************")
	data := types.TWDataCreateWithGoBytes(entropy)
	defer C.TWDataDelete(data)

	fmt.Println("DDDDDDDD")
	fmt.Println(types.TWDataHexString(data))

	fmt.Println(data)
	fmt.Println("*****************")
	tw := C.TWHDWalletCreateWithEntropy(data, pass)
	defer C.TWHDWalletDelete(tw)
	fmt.Println(tw)
	fmt.Println("*****************")

	fmt.Println(C.enum_TWCoinType(ct))
	fmt.Println("*****************")
	priKey := C.TWHDWalletGetKeyForCoin(tw, C.enum_TWCoinType(ct))
	defer C.TWPrivateKeyDelete(priKey)
	priKeyData := C.TWPrivateKeyData(priKey)
	defer C.TWDataDelete(priKeyData)
	fmt.Println(types.TWDataHexString(priKeyData))
	return types.TWDataHexString(priKeyData)
}
