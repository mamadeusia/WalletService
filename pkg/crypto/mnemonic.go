package crypto

// #cgo CFLAGS: -I../../trustwallet/include
// #cgo LDFLAGS: -L../../trustwallet/build -L../../trustwallet/build/local/lib -L../../trustwallet/build/trezor-crypto -lTrustWalletCore -lwallet_core_rs -lprotobuf -lTrezorCrypto -lstdc++ -lm
// #include <TrustWalletCore/TWMnemonic.h>
import "C"

import "github.com/mamadeusia/WalletSrv/pkg/types"

func IsMnemonicValid(mn string) bool {
	str := types.TWStringCreateWithGoString(mn)
	defer C.TWStringDelete(str)
	return bool(C.TWMnemonicIsValid(str))
}
