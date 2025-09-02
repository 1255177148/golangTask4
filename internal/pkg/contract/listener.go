package contract

// Listener 是所有合约监听器需要实现的接口
type Listener interface {
	StartListen()
	StopListen()
}
