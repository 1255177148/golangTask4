package pkg

import "github.com/ethereum/go-ethereum/ethclient"

type ContractInstance struct {
	HttpClient *ethclient.Client
	WsClient   *ethclient.Client
}

var ContractClient *ContractInstance
