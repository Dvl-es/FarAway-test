package config

import (
	"github.com/joho/godotenv"
	"github.com/namsral/flag"
)

type Config struct {
	IsRelease bool
	GinPort   string

	RpcURL            string
	NftManagerAddress string
}

func NewConfig() *Config {
	godotenv.Load(".env")

	config := &Config{}

	//API
	flag.BoolVar(&config.IsRelease, "RELEASE_MODE", false, "run gin in release mode")
	flag.StringVar(&config.GinPort, "GIN_PORT", "6565", "")

	flag.StringVar(&config.RpcURL, "RPC_URL", "", "wss rpc endpoint")
	flag.StringVar(&config.NftManagerAddress, "NFT_MANAGER_ADDRESS", "", "Nft manager contract address")

	flag.Parse()

	return config
}
