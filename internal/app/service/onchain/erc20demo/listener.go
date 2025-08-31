package erc20demo

import (
	"context"
	"fmt"
	"github.com/1255177148/golangTask4/contract/erc20demo"
	"github.com/1255177148/golangTask4/internal/pkg"
	"github.com/1255177148/golangTask4/internal/utils/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"sync"
	"time"
)

type Listener struct {
	client          *ethclient.Client
	service         *Erc20Service
	ctx             context.Context
	cancel          context.CancelFunc
	lastBlock       uint64 // 实时事件最大区块
	lastSyncedBlock uint64 // 历史事件同步进度
	mu              sync.Mutex
	processed       map[string]struct{}
}

// NewListener 初始化监听实例
func NewListener(service *Erc20Service, startBlock uint64) *Listener {
	ctx, cancel := context.WithCancel(context.Background())
	return &Listener{
		client:          pkg.ContractClient.WsClient,
		service:         service,
		ctx:             ctx,
		cancel:          cancel,
		lastBlock:       startBlock,
		lastSyncedBlock: startBlock,
		processed:       make(map[string]struct{}),
	}
}

// Start 启动监听
func (l *Listener) Start() {
	go l.syncLoop()  // 历史事件同步
	go l.watchLoop() // 实时事件订阅
}

// Stop 停止监听
func (l *Listener) Stop() {
	l.cancel()
}

// syncLoop 历史事件同步循环
func (l *Listener) syncLoop() {
	ticker := time.NewTicker(time.Minute * 2) //每2分钟触发一次
	defer ticker.Stop()

	for {
		select {
		case <-l.ctx.Done():
			return
		case <-ticker.C:
			l.syncHistory()
		}
	}
}

func (l *Listener) syncHistory() {
	header, err := l.client.HeaderByNumber(l.ctx, nil)
	if err != nil {
		log.Error("获取最新区块失败:", zap.Error(err))
		return
	}
	latest := header.Number.Uint64()

	if l.lastSyncedBlock >= latest {
		return
	}

	batchSize := uint64(5000) // 每批拉取 5000 个区块，可根据节点性能调整
	start := l.lastSyncedBlock + 1
	for start <= latest {
		end := start + batchSize - 1
		if end > latest {
			end = latest
		}

		// ---------------- Transfer 历史事件 ----------------
		transferIter, err := l.service.WsInstance.FilterTransfer(&bind.FilterOpts{Start: start, End: &end, Context: l.ctx})
		if err != nil {
			log.Error("FilterTransfer失败: ", zap.Error(err))
		} else {
			for transferIter.Next() {
				ev := transferIter.Event
				l.handleEvent(ev.Raw.BlockNumber, ev.Raw.Index,
					fmt.Sprintf("[历史 Transfer] %s -> %s, value=%s", ev.From.Hex(), ev.To.Hex(), ev.Amount.String()))
			}
		}

		// ---------------- Mint 历史事件 ----------------
		mintIter, err := l.service.WsInstance.FilterMint(&bind.FilterOpts{Start: start, End: &end, Context: l.ctx})
		if err != nil {
			log.Error("FilterMint失败: ", zap.Error(err))
		} else {
			for mintIter.Next() {
				ev := mintIter.Event
				l.handleEvent(ev.Raw.BlockNumber, ev.Raw.Index,
					fmt.Sprintf("[历史 Mint] -> %s, value=%s", ev.To.Hex(), ev.Amount.String()))
			}
		}
		// 更新历史同步进度
		l.mu.Lock()
		l.lastSyncedBlock = end
		l.mu.Unlock()
		start = end + 1
	}
}

// watchLoop 实时事件订阅
func (l *Listener) watchLoop() {
	transferCh := make(chan *erc20demo.Erc20demoTransfer)
	mintCh := make(chan *erc20demo.Erc20demoMint)

	transferSub, err := l.service.WsInstance.WatchTransfer(&bind.WatchOpts{Context: l.ctx}, transferCh)
	if err != nil {
		log.Fatal("实时订阅 Transfer 失败: ", zap.Error(err))
	}
	mintSub, err := l.service.WsInstance.WatchMint(&bind.WatchOpts{Context: l.ctx}, mintCh)
	if err != nil {
		log.Fatal("实时订阅 Mint 失败: ", zap.Error(err))
	}

	for {
		select {
		case <-l.ctx.Done():
			transferSub.Unsubscribe() // 关闭订阅通道
			mintSub.Unsubscribe()     // 关闭订阅通道
			return
		case err := <-transferSub.Err():
			log.Error("Transfer订阅错误:", zap.Error(err))
			time.Sleep(time.Second * 5)
			log.Debug("开始重新连接")
			go l.watchLoop() // 异常重连
			return
		case err := <-mintSub.Err():
			log.Error("Mint订阅错误:", zap.Error(err))
			time.Sleep(time.Second * 5)
			log.Debug("开始重新连接")
			go l.watchLoop() // 异常重连
			return
		case ev := <-transferCh:
			l.handleEvent(ev.Raw.BlockNumber, ev.Raw.Index,
				fmt.Sprintf("[历史 Mint] -> %s, value=%s", ev.To.Hex(), ev.Amount.String()))
		case ev := <-mintCh:
			l.handleEvent(ev.Raw.BlockNumber, ev.Raw.Index,
				fmt.Sprintf("[历史 Mint] -> %s, value=%s", ev.To.Hex(), ev.Amount.String()))
		}
	}
}

// handleEvent 处理事件并去重
func (l *Listener) handleEvent(blockNumber uint64, logIndex uint, msg string) {
	id := fmt.Sprintf("%d-%d", blockNumber, logIndex)
	l.mu.Lock()
	defer l.mu.Unlock()
	if _, ok := l.processed[id]; ok {
		return // 已处理过
	}
	l.processed[id] = struct{}{}

	// 更新实时最大区块，仅用于监控
	if blockNumber > l.lastBlock {
		l.lastBlock = blockNumber
	}

	fmt.Println(msg)
	// TODO: 入库或触发业务逻辑
}
