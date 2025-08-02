// 异步任务worker pool

package taskpool

import (
	"context"
	"github.com/1255177148/golangTask4/internal/utils/log"
	"sync"
	"time"
)

// Task 定义任务函数类型,接收 context 以支持超时/取消
type Task func(ctx context.Context) error

// Job 代表一个任务，包括超时时间和最大重试次数
type Job struct {
	Task       Task          // 任务函数
	Timeout    time.Duration // 超时时间
	MaxRetries int           // 最大重试次数
}

type WorkerPool struct {
	TaskQueue chan Job       // 任务队列
	wg        sync.WaitGroup //等待所有任务完成
	quit      chan struct{}  // 用于优雅关闭worker pool
}

func NewWorkerPool(workerCount int, queueSize int) *WorkerPool {
	wp := &WorkerPool{
		TaskQueue: make(chan Job, queueSize),
		quit:      make(chan struct{}),
	}
	// 启动workerCount个worker
	for i := 0; i < int(workerCount); i++ {
		go wp.worker(i)
	}
	return wp
}

// worker 每个worker不断从TaskQueue取任务执行
func (wp *WorkerPool) worker(id int) {
	for {
		select {
		case <-wp.quit:
			// 收到关闭信号，退出worker
			log.DebugF("❌ Worker-%d 停止运行\n", id)
			return
		case job := <-wp.TaskQueue:
			// 收到任务，就执行任务
			wp.wg.Add(1)
			wp.executeJob(id, job)
			wp.wg.Done()
		}
	}
}

// executeJob 执行任务，支持超时和重试
func (wp *WorkerPool) executeJob(workerID int, job Job) {
	var err error
	// 重试指定次数
	for attempt := 1; attempt <= job.MaxRetries; attempt++ {
		// 创建带超时的context
		ctx, cancel := context.WithTimeout(context.Background(), job.Timeout)
		log.DebugF("🚀 Worker-%d 开始任务（第 %d 次尝试）\n", workerID, attempt)
		err = job.Task(ctx) // 执行任务
		cancel()            // 手动关闭context
		if err == nil {
			log.DebugF("✅ Worker-%d 任务完成\n", workerID)
			return
		}
		log.DebugF("❌ Worker-%d 任务失败:%v\n", workerID, err)
	}
	log.DebugF("❌ Worker-%d 任务最终失败\n", workerID)
}

// Submit 提交任务，若队列已满则返回false
func (wp *WorkerPool) Submit(job Job) bool {
	select {
	case wp.TaskQueue <- job:
		return true
	default:
		return false
	}
}

// Shutdown 优雅关闭，等待所有任务执行完成后关闭
func (wp *WorkerPool) Shutdown() {
	close(wp.quit)
	wp.wg.Wait()
	log.Debug("✅ 所有任务已完成, Worker Pool 已关闭")
}
