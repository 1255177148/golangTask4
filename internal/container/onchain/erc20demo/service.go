package erc20demo

import (
	"github.com/1255177148/golangTask4/config"
	"github.com/1255177148/golangTask4/internal/app/service/onchain/erc20demo"
	"github.com/1255177148/golangTask4/internal/utils/log"
	"go.uber.org/zap"
)

func InitERC20DemoService() *erc20demo.Erc20Service {
	serviceInstance, err := erc20demo.NewERC20Contract(config.Cfg.Contract.ERC20Contract.Address)
	if err != nil {
		log.Fatal("init erc20 service failed", zap.Error(err))
	}
	return serviceInstance
}
