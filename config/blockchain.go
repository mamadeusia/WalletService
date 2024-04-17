package config

type Blockchain struct {
	Name string
}

func GetBlockchainServiceName() string {
	return cfg.Blockchain.Name
}
