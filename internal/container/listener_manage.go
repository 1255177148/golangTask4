package container

import (
	"github.com/1255177148/golangTask4/internal/pkg/contract"
)

var ListenerManager = &contract.ListenerManager{}

// InitListenerManage 初始化Listener管理器
func InitListenerManage() {
	ListenerManager.Add(Instance.ERC20Listener)

	// 启动全部
	ListenerManager.StartAll()
}

// StopAllListeners 关闭所有的合约Listener
func StopAllListeners() {
	ListenerManager.StopAll()
}
