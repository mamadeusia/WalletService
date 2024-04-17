package entity

// we need to read data from vault to be more secure
type Secret struct {
	Mnemunic          string
	DerivationVersion DerivationVersion
	Index             uint32
}

// we need to have vault as secret management system for feautre versions
// but for implementations for demo we just try to write code in a way that be most compatible with
// vault implementation

// so we use the seed hash in the database and every time we need to do something with secrets like
// sign transactions and etc we should get the appropriate data from vault and regenrate the private key
// from that .
