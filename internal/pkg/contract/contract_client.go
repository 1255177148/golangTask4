package contract

import "github.com/ethereum/go-ethereum/ethclient"

type Instance struct {
	HttpClient *ethclient.Client
	WsClient   *ethclient.Client
}

var Client *Instance
