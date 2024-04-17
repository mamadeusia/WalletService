package crypto

// #cgo CFLAGS: -I../../trustwallet/include
// #cgo LDFLAGS: -L../../trustwallet/build -L../../trustwallet/build/local/lib -L../../trustwallet/build/trezor-crypto -lTrustWalletCore -lwallet_core_rs -lprotobuf -lTrezorCrypto -lstdc++ -lm
// #include <TrustWalletCore/TWCoinType.h>
// #include <TrustWalletCore/TWCoinTypeConfiguration.h>
import "C"

import "github.com/mamadeusia/WalletSrv/pkg/types"

type CoinType uint32

const (
	CoinTypeBitcoin  CoinType = C.TWCoinTypeBitcoin
	CoinTypeBinance  CoinType = C.TWCoinTypeBinance
	CoinTypeEthereum CoinType = C.TWCoinTypeEthereum
	CoinTypeTron     CoinType = C.TWCoinTypeTron
)

func (c CoinType) GetName() string {
	name := C.TWCoinTypeConfigurationGetName(C.enum_TWCoinType(c))
	defer C.TWStringDelete(name)
	return types.TWStringGoString(name)
}

func (c CoinType) Decimals() int {
	return int(C.TWCoinTypeConfigurationGetDecimals(C.enum_TWCoinType(c)))
}
