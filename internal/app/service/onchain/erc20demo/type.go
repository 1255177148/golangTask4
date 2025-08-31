package erc20demo

// MintResult 定义一个结果结构体，方便返回 txHash 或 error
type MintResult struct {
	TxHash string
	Err    error
}
