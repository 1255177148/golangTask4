package bootstrap

import (
	"github.com/1255177148/golangTask4/config"
	"github.com/1255177148/golangTask4/internal/pkg/contract"
	"github.com/1255177148/golangTask4/internal/utils/log"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

// InitContractClient 初始化合约client
func InitContractClient() {
	httpClient, err := ethclient.Dial(config.Cfg.Contract.HttpUrl)
	if err != nil {
		log.Fatal("Contract http client error", zap.Error(err))
	}
	wsClient, err := ethclient.Dial(config.Cfg.Contract.WebsocketUrl)
	if err != nil {
		log.Fatal("Contract websocket client error", zap.Error(err))
	}
	contract.Client = &contract.Instance{
		HttpClient: httpClient,
		WsClient:   wsClient,
	}
}
