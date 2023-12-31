package config

// // NOTE: ONLY PUT STRUCT DEFINITIONS IN THIS FILE
// //
// // After making edits here, run 'make cfgdoc-gen' (or 'make gen')

// Common is common config between full node and miner
type Common struct {
	API      API
	RouteCfg RouteCfg
}

// API contains configs for API endpoint
type API struct {
	ListenAddress       string
	RemoteListenAddress string
	Timeout             Duration
}

// TransactionCfg transaction config
type TransactionCfg struct {
	Common
	// database address
	DatabaseAddress string
}

// MallCfg base config
type MallCfg struct {
	Common
	// used when 'ListenAddress' is unspecified. must be a valid duration recognized by golang's time.ParseDuration function
	Timeout string

	DryRun                bool
	AliyunAccessKeyID     string
	AliyunAccessKeySecret string

	DatabaseAddress string

	TitanContractorAddr string
	LotusWsAddr         string
	LotusHTTPSAddr      string
	PrivateKeyStr       string
	PaymentAddresses    []string

	TrxHTTPSAddr      string
	TrxContractorAddr string
	RechargeAddresses []string

	TrxHeight int64

	Email EmailConfig
}

type RouteCfg struct {
	EtcdAddress string
	Mode        string
	ApiListen   string
	SecretKey   string
}

type EmailConfig struct {
	Name     string
	SMTPHost string
	SMTPPort string
	Password string
}
