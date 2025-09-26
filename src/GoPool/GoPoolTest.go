package GoPool

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

//工作协程的结构体以及其方法
/*
graph TD
	A[启动Dispatcher] --> B[初始化Worker池]
	B --> C[Worker持续注册到池]
	C --> D{有新任务?}
	D -- 是 --> E[获取空闲Worker通道]
	E --> F[发送任务到Worker]
	F --> G[执行任务并Done]
	G --> C
	D -- 否 --> H{所有任务完成?}
	H -- 是 --> I[结束]
	H -- 否 --> D
*/

// Worker 工作协程
type Worker struct {
	ID         int
	taskChan   chan *TaskWithResult      // 任务通道
	workerPool chan chan *TaskWithResult // 工作协程通道
	quit       chan bool                 // 退出通道
}

// NewWorker 创建一个新的工作协程
func NewWorker(id int, workerPool chan chan *TaskWithResult) *Worker {
	return &Worker{
		ID:         id,
		taskChan:   make(chan *TaskWithResult),
		workerPool: workerPool,
		quit:       make(chan bool),
	}
}

// Start 启动工作协程
func (w *Worker) Start() {
	go func() {
		for {
			// 将各自的taskChan推送到workerPool
			w.workerPool <- w.taskChan
			select { //等待阶段
			case task := <-w.taskChan:
				// 从任务通道中获取任务并执行
				result, err := task.Task()
				if err != nil {
					task.Error <- err
					close(task.Result)
				} else {
					task.Result <- result
					close(task.Error)
				}
			case <-w.quit:
				return
			}
		}
	}()
}

// Stop 停止工作协程
func (w *Worker) Stop() {
	// 发送退出信号
	go func() {
		w.quit <- true
	}()
}

// ************************************************ //
// Dispatcher 调度器结构图及其方法

// Dispatcher 调度器，用于管理多个worker
type Dispatcher struct {
	workerPool chan chan *TaskWithResult
	maxWorkers int
	wg         sync.WaitGroup
	//优化新增任务队列
	taskQueue chan *TaskWithResult // 新增任务队列
	queueSize int                  // 任务队列大小
}

// NewDispatcher 创建一个新的调度器
func NewDispatcher(maxWorkers, queueSize int) *Dispatcher {
	return &Dispatcher{
		workerPool: make(chan chan *TaskWithResult, maxWorkers),
		maxWorkers: maxWorkers,
		//优化新增任务队列
		taskQueue: make(chan *TaskWithResult, queueSize), // 带缓冲队列
		queueSize: queueSize,
	}
}

// Run 启动调度器，创建并启动多个worker
func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		// 创建并启动一个worker
		worker := NewWorker(i, d.workerPool)
		worker.Start()
	}
	//优化新增任务队列
	//统一分配任务
	go func() {
		for task := range d.taskQueue {
			go func(t *TaskWithResult) {
				workerChan := <-d.workerPool // 获取空闲worker的任务通道
				workerChan <- t              //发送任务到worker
			}(task)
		}
	}()
}

// Submit 提交一个任务到调度器
func (d *Dispatcher) Submit(task func() (interface{}, error), timeout time.Duration) (interface{}, error) {
	taskWithResult := &TaskWithResult{
		Task:   task,
		Result: make(chan interface{}, 1),
		Error:  make(chan error, 1),
	}
	select {
	case d.taskQueue <- taskWithResult:
		d.wg.Add(1)
		defer d.wg.Done()
		select {
		case result := <-taskWithResult.Result:
			return result, nil
		case err := <-taskWithResult.Error:
			return nil, err
		case <-time.After(timeout): // 任务执行超时
			return nil, errors.New("task execution timeout")
		}
	case <-time.After(timeout): // 提交队列超时
		return nil, errors.New("task queue full")
	}

	//优化前代码
	//d.wg.Add(1)
	//go func() {
	//	taskChan := <-d.workerPool      // 获取空闲worker的任务通道
	//	taskChan <- func() { // 发送任务到worker
	//		defer d.wg.Done()
	//		task()
	//	}
	//}()
}

func (d *Dispatcher) Wait() {
	d.wg.Wait()
}

// **********************支持返回值************************** //

// TaskWithResult 带返回值的任务类型
type TaskWithResult struct {
	Task   func() (interface{}, error) // 返回结果和错误的函数
	Result chan interface{}            // 结果通道
	Error  chan error                  // 错误通道
}

func GoPoolsDemo() {
	numTasks := 10

	// Create a new dispatcher.
	dispatcher := NewDispatcher(3, 10)
	dispatcher.Run()

	// Submit tasks to the dispatcher.
	for i := 0; i < numTasks; i++ {
		taskID := i
		result, _ := dispatcher.Submit(func() (interface{}, error) {
			fmt.Printf("Task %d is running\n", taskID)
			time.Sleep(time.Second) // Simulate task execution time
			return fmt.Printf("Task %d is done\n", taskID), nil
		}, time.Second)

		fmt.Println("result::", result)
	}

	// Wait for all tasks to be completed.
	dispatcher.Wait()
	fmt.Println("All tasks are completed")
}
