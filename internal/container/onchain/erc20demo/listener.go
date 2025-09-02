package erc20demo

import (
	"github.com/1255177148/golangTask4/config"
	"github.com/1255177148/golangTask4/internal/app/service/onchain/erc20demo"
)

// InitERC20Listener 初始化ERC20Demo合约的事务监听
//
// 参数：
//
//	service    - erc20service依赖
//	startBlock - 合约部署时的区块号，可以从配置文件中获取
func InitERC20Listener(service *erc20demo.Erc20Service) *erc20demo.Listener {
	erc20demoListener := erc20demo.NewListener(service, config.Cfg.Contract.ERC20Contract.ERC20BlockNumber)
	erc20demoListener.Start()
	return erc20demoListener
}
