package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

const (
	defaultMaxRetry  = 3
	defaultGasBuffer = 1.2 // Gas buffer系数
)

// TxSender 托管模式交易发送器
type TxSender struct {
	Client     *ethclient.Client
	PrivateKey *ecdsa.PrivateKey
	MaxRetry   int
}

// NewTxSender 创建交易发送器
func NewTxSender(client *ethclient.Client, privateKey *ecdsa.PrivateKey) *TxSender {
	return &TxSender{
		Client:     client,
		PrivateKey: privateKey,
		MaxRetry:   defaultMaxRetry,
	}
}

// SendWithRetry 托管模式下发送交易,使用回调函数的形式来调用abi里面的合约函数
// 支持动态方法名、可变参数、Gas模拟、动态nonce和自动重试
//
// 入参说明：
//
//	ctx          - 上下文，可用于控制请求超时或取消
//	privateKey   - 后端托管账户私钥（*ecdsa.PrivateKey）
//	contractCallFunc   - 合约方法调用函数，例如 "setItem()" 或 "transfer()"
//	args         - 可变参数（...interface{}），对应方法入参
//
// 返回值：
//
//	string - 交易 Hash
//	error  - 如果发送失败，返回具体错误
func (s *TxSender) SendWithRetry(
	ctx context.Context,
	contractCallFunc func(*bind.TransactOpts, ...interface{}) (*types.Transaction, error)) (string, error) {

	if s.PrivateKey == nil {
		return "", fmt.Errorf("托管模式必须提供 PrivateKey")
	}
	// ---------------- 自动获取链ID ----------------
	chainID, err := s.Client.ChainID(ctx)
	if err != nil {
		return "", fmt.Errorf("获取链ID失败: %w", err)
	}
	// 创建托管账户的 TransactOpts 对象，绑定私钥和 chainID
	auth, err := bind.NewKeyedTransactorWithChainID(s.PrivateKey, chainID)
	if err != nil {
		return "", fmt.Errorf("创建Transactor失败: %w", err)
	}
	// 最大重试机制
	maxRetry := 3
	for attempt := 0; attempt < maxRetry; attempt++ {
		// ---------------- 获取最新 pending nonce ----------------
		nonce, err := s.Client.PendingNonceAt(ctx, auth.From)
		if err != nil {
			return "", fmt.Errorf("获取 pending nonce 失败: %w", err)
		}
		auth.Nonce = big.NewInt(int64(nonce))

		// ---------------- 模拟估算 Gas ----------------

		// 初始 GasLimit 为0，由 abigen 内部估算
		auth.GasLimit = 0
		// 设置 GasTipCap & GasFeeCap
		tipCap, err := s.Client.SuggestGasTipCap(ctx)
		if err != nil {
			return "", fmt.Errorf("获取建议小费单价失败: %w", err)
		}
		header, err := s.Client.HeaderByNumber(ctx, nil)
		if err != nil {
			return "", fmt.Errorf("获取最新区块头失败: %w", err)
		}
		baseFee := header.BaseFee
		if baseFee == nil {
			baseFee = big.NewInt(1e9)
		}
		feeCap := new(big.Int).Add(new(big.Int).Mul(baseFee, big.NewInt(2)), tipCap)
		auth.GasTipCap = tipCap
		auth.GasFeeCap = feeCap

		// ---------------- 发送交易 ----------------
		tx, err := contractCallFunc(auth)
		if err != nil {
			log.Printf("发送交易失败 (attempt %d/%d): %v", attempt+1, maxRetry, err)
			time.Sleep(2 * time.Second)
		} else {
			return tx.Hash().Hex(), nil
		}
	}
	return "", fmt.Errorf("交易发送失败，重试 %d 次仍然失败", maxRetry)
}

// SendWithRetryByGasLimit 托管模式下发送交易,使用gasLimit手动预估，并用函数名的形式调用
// 支持动态方法名、可变参数、Gas模拟、动态nonce和自动重试
//
// 入参说明：
//
//		ctx          - 上下文，可用于控制请求超时或取消
//		contractABI  - 合约ABI
//		contractAddr - 合约地址
//		method       - 要调用的合约函数名
//	 params       - 可变参数，合约函数的入参
//
// 返回值：
//
//	string - 交易 Hash
//	error  - 如果发送失败，返回具体错误
func (s *TxSender) SendWithRetryByGasLimit(
	ctx context.Context,
	contractABI *abi.ABI,
	contractAddr *common.Address,
	method string,
	params ...interface{}) (string, error) {

	if s.PrivateKey == nil {
		return "", fmt.Errorf("托管模式必须提供 PrivateKey")
	}
	// ---------------- 自动获取链ID ----------------
	chainID, err := s.Client.ChainID(ctx)
	if err != nil {
		return "", fmt.Errorf("获取链ID失败: %w", err)
	}
	// 创建托管账户的 TransactOpts 对象，绑定私钥和 chainID
	auth, err := bind.NewKeyedTransactorWithChainID(s.PrivateKey, chainID)
	if err != nil {
		return "", fmt.Errorf("创建Transactor失败: %w", err)
	}
	// 最大重试机制
	maxRetry := 3
	for attempt := 0; attempt < maxRetry; attempt++ {
		// ---------------- 获取最新 pending nonce ----------------
		nonce, err := s.Client.PendingNonceAt(ctx, auth.From)
		if err != nil {
			return "", fmt.Errorf("获取 pending nonce 失败: %w", err)
		}
		auth.Nonce = big.NewInt(int64(nonce))

		// ---------------- 模拟估算 Gas ----------------
		data, err := contractABI.Pack(method, params...)
		if err != nil {
			return "", fmt.Errorf("生成交易数据失败: %w", err)
		}
		fmt.Println("data:", string(data))
		callMsg := ethereum.CallMsg{
			From:  auth.From,
			To:    contractAddr,
			Value: big.NewInt(0),
			Data:  data,
		}
		gasEstimate, err := s.Client.EstimateGas(ctx, callMsg)
		if err != nil {
			log.Printf("警告: 估算 Gas 失败: %v", err)
			gasEstimate = 0
		}
		auth.GasLimit = uint64(float64(gasEstimate) * defaultGasBuffer)
		if auth.GasLimit == 0 {
			auth.GasLimit = 300_000 // 默认值防止为0
		}
		fmt.Println("gasLimit:", auth.GasLimit)
		// 设置 GasTipCap & GasFeeCap
		tipCap, err := s.Client.SuggestGasTipCap(ctx)
		if err != nil {
			return "", fmt.Errorf("获取建议小费单价失败: %w", err)
		}
		header, err := s.Client.HeaderByNumber(ctx, nil)
		if err != nil {
			return "", fmt.Errorf("获取最新区块头失败: %w", err)
		}
		baseFee := header.BaseFee
		if baseFee == nil {
			baseFee = big.NewInt(1e9)
		}
		feeCap := new(big.Int).Add(new(big.Int).Mul(baseFee, big.NewInt(2)), tipCap)
		auth.GasTipCap = tipCap
		auth.GasFeeCap = feeCap

		// ---------------- 发送交易 ----------------
		tx := types.NewTx(&types.DynamicFeeTx{
			ChainID:   chainID,
			Nonce:     nonce,
			GasTipCap: tipCap,
			GasFeeCap: feeCap,
			Gas:       auth.GasLimit,
			To:        contractAddr,
			Value:     big.NewInt(0),
			Data:      data,
		})

		signedTx, err := auth.Signer(auth.From, tx)
		if err != nil {
			return "", fmt.Errorf("签名交易失败: %w", err)
		}

		err = s.Client.SendTransaction(ctx, signedTx)
		if err != nil {
			log.Printf("发送交易失败 (attempt %d/%d): %v", attempt+1, maxRetry, err)
			time.Sleep(2 * time.Second)
		}
		txHash := signedTx.Hash().Hex()
		// ---------------- 校验交易是否已广播 ----------------
		for i := 0; i < 5; i++ { // 最多尝试5次，每次间隔1秒
			_, isPending, err := s.Client.TransactionByHash(ctx, signedTx.Hash())
			if err == nil && isPending {
				break // 交易在 mempool
			}
			time.Sleep(1 * time.Second)
		}

		// ---------------- 等待交易上链 ----------------
		receipt, err := bind.WaitMined(ctx, s.Client, signedTx)
		if err != nil {
			log.Printf("等待交易上链失败 (attempt %d/%d): %v", attempt+1, maxRetry, err)
			time.Sleep(2 * time.Second)
			continue
		}
		if receipt.Status == 1 {
			return txHash, nil
		} else {
			log.Printf("交易失败上链 (attempt %d/%d)，状态码: %d", attempt+1, maxRetry, receipt.Status)
		}
	}
	return "", fmt.Errorf("交易发送失败，重试 %d 次仍然失败", maxRetry)
}

// UnsignedTx 构造未签名交易，支持 EIP-1559
//
// 入参说明：
//
//	from         - 交易的发起者地址
//	contractABI  - 合约对象的 ABI
//	contractAddr - 合约地址
//	funcName     - 函数名称
//	params       - 可变参数（...interface{}），对应方法入参
//
// 返回值：
//
//	map[string]interface{} - 交易数据
//	error                  - 如果构建交易数据失败，返回具体错误
func (s *TxSender) UnsignedTx(
	from common.Address,
	contractABI *abi.ABI,
	contractAddr common.Address,
	funcName string,
	params ...interface{},
) (map[string]interface{}, error) {

	// 1️⃣ 使用 ABI.Pack 构造交易 data
	data, err := contractABI.Pack(funcName, params...)
	if err != nil {
		return nil, fmt.Errorf("ABI.Pack failed: %w", err)
	}

	// 2️⃣ 预估 gasLimit
	msg := ethereum.CallMsg{
		From: from,
		To:   &contractAddr,
		Data: data,
	}
	gasLimit, err := s.Client.EstimateGas(context.Background(), msg)
	if err != nil {
		return nil, fmt.Errorf("EstimateGas failed: %w", err)
	}
	// 可加 buffer
	gasLimit = uint64(float64(gasLimit) * 1.2)

	// 3️⃣ EIP-1559
	tipCap, err := s.Client.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, fmt.Errorf("SuggestGasTipCap failed: %w", err)
	}

	header, err := s.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("HeaderByNumber failed: %w", err)
	}

	baseFee := header.BaseFee
	if baseFee == nil {
		baseFee = big.NewInt(1e9) // 默认 1 Gwei
	}

	maxFee := new(big.Int).Add(new(big.Int).Mul(baseFee, big.NewInt(2)), tipCap)

	// 4️⃣ 返回前端 JSON
	txData := map[string]interface{}{
		"from":                 from.Hex(),
		"to":                   contractAddr.Hex(),
		"value":                "0",
		"data":                 fmt.Sprintf("0x%x", data),
		"gasLimit":             gasLimit,
		"maxPriorityFeePerGas": tipCap.String(),
		"maxFeePerGas":         maxFee.String(),
	}

	return txData, nil
}
