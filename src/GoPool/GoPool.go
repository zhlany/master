package GoPool

import (
	"fmt"
	"sync"
)

// Result 定义一个通用的结果结构体
type Result interface {
	// SetResult 设置结果
	SetResult(interface{})
}

// GoroutinePool 定义一个协程池结构体
type GoroutinePool struct {
	maxWorkers int         // 最大工作协程数
	workers    []*Worker   // 工作协程数组
	taskChan   chan func() // 任务通道
	resultChan chan Result // 结果通道
	once       sync.Once   // 只执行一次
}

// Worker 定义一个工作协程结构体
type Worker struct {
	taskChan   chan func() // 任务通道
	resultChan chan Result // 结果通道
	once       sync.Once   // 只执行一次
}

// NewGoroutinePool 创建一个协程池
func NewGoroutinePool(maxWorkers int) *GoroutinePool {
	return &GoroutinePool{
		maxWorkers: maxWorkers,
		taskChan:   make(chan func()),
		resultChan: make(chan Result),
	}
}

// Start 启动协程池
func (p *GoroutinePool) Start() {
	p.once.Do(func() {
		for i := 0; i < p.maxWorkers; i++ {
			worker := &Worker{
				taskChan:   make(chan func()),
				resultChan: make(chan Result),
			}
			p.workers = append(p.workers, worker)
			go worker.Run()
		}
	})
}

// Submit 提交任务到协程池
func (p *GoroutinePool) Submit(task func(), result Result) {
	p.taskChan <- task
	result.SetResult(<-p.resultChan)
}

// Wait 等待任务完成
func (p *GoroutinePool) Wait() {
	for range p.taskChan {
	}
}

// Close 关闭协程池
func (p *GoroutinePool) Close() {
	close(p.taskChan)
	close(p.resultChan)
}

type CustomResult struct {
	Result1 int
	Result2 string
}

func (c CustomResult) SetResult(i interface{}) {
	//TODO implement me
	panic("implement me")
}

// Run 工作协程运行函数
func (w *Worker) Run() {
	for task := range w.taskChan {
		task()
		w.resultChan <- &CustomResult{
			Result1: 1,
			Result2: "hello",
		}
	}
}

func goPoolt() {
	pool := NewGoroutinePool(2)
	pool.Start()

	var wg sync.WaitGroup
	var result CustomResult
	pool.Submit(func() {
		fmt.Println("Hello, world!")
	}, &result)

	wg.Wait()
	pool.Wait()
	pool.Close()
}
