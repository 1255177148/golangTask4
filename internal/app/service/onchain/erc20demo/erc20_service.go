package erc20demo

import (
	"context"
	"fmt"
	"github.com/1255177148/golangTask4/contract/erc20demo"
	"github.com/1255177148/golangTask4/internal/pkg"
	"github.com/1255177148/golangTask4/internal/utils/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type Erc20Service struct {
	HttpInstance    *erc20demo.Erc20demo
	WsInstance      *erc20demo.Erc20demo
	ContractAddress *common.Address
}

// NewERC20Contract 创建合约实例
func NewERC20Contract(address string) (*Erc20Service, error) {
	contractAddr := common.HexToAddress(address)
	httpInstance, err := erc20demo.NewErc20demo(contractAddr, pkg.ContractClient.HttpClient)
	if err != nil {
		return nil, err
	}
	wsInstance, err := erc20demo.NewErc20demo(contractAddr, pkg.ContractClient.WsClient)
	if err != nil {
		return nil, err
	}
	return &Erc20Service{HttpInstance: httpInstance, WsInstance: wsInstance, ContractAddress: &contractAddr}, nil
}

func MintAsync(e *Erc20Service, account common.Address, amount *big.Int) <-chan MintResult {
	resultChan := make(chan MintResult, 1)
	go func() {
		defer close(resultChan)
		txHash, err := e.Mint(account, amount)
		resultChan <- MintResult{TxHash: txHash, Err: err}
	}()
	return resultChan
}

func (e *Erc20Service) Mint(account common.Address, amount *big.Int) (string, error) {
	privateKey, err := crypto.HexToECDSA("cabb9d1405205e92b2984ac19fbf28b17432d1f0af889d867a5df7e0e851cf4b") // 解析私钥
	if err != nil {
		panic(err)
	}
	txSender := contract.NewTxSender(pkg.ContractClient.HttpClient, privateKey)

	// 第一种调用，gasLimit没有buffer余量，直接为预估gasLimit
	//txHash, err := txSender.SendWithRetry(context.Background(), func(opts *bind.TransactOpts, i ...interface{}) (*types.Transaction, error) {
	//	return e.Instance.Mint(opts, account, amount)
	//})

	// 第二种调用，gasLimit有buffer余量，在复杂跨合约调用时可以考虑用这种
	contractABI, err := erc20demo.Erc20demoMetaData.GetAbi()
	if err != nil {
		return "", fmt.Errorf("GetAbi error: %v", err)
	}
	txHash, err := txSender.SendWithRetryByGasLimit(context.Background(), contractABI, e.ContractAddress, "mint", account, amount)
	if err != nil {
		return "", err
	}
	fmt.Printf("txHash: %s\n", txHash)
	return txHash, nil
}

// BalanceOf 调用solidity的balanceOf查询函数
func (e *Erc20Service) BalanceOf(account common.Address) (*big.Int, error) {
	callOpts := &bind.CallOpts{
		Pending: false, // 表示查询最新区块
		Context: context.Background(),
		From:    common.Address{}, // 是solidity中的msg.sender，如果不需要可以不传
	}
	balance, err := e.HttpInstance.BalanceOf(callOpts, account)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
