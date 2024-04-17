package config

type Node struct {
	URL      string
	Mnemonic string
}

func NodeURL() string {
	return cfg.Node.URL
}

func NodeMnemonic() string {
	return cfg.Node.Mnemonic
}
