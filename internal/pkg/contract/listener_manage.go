package contract

import "sync"

type ListenerManager struct {
	mu        sync.Mutex
	listeners []Listener
}

// Add 添加 listener
func (m *ListenerManager) Add(l Listener) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.listeners = append(m.listeners, l)
}

// StartAll 启动全部 listener
func (m *ListenerManager) StartAll() {
	m.mu.Lock()
	listeners := append([]Listener(nil), m.listeners...) // 拷贝 slice 避免锁住整个循环
	m.mu.Unlock()
	for _, l := range listeners {
		l.StartListen()
	}
}

// StopAll 停止全部 listener
func (m *ListenerManager) StopAll() {
	m.mu.Lock()
	listeners := append([]Listener(nil), m.listeners...)
	m.mu.Unlock()
	for _, l := range listeners {
		l.StopListen()
	}
}
