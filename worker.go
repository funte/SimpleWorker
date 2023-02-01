package SimpleWorker

import (
	"sync"
	"time"
)

// 可以启动停止的协程对象.
type Worker struct {
	Name       string
	IsRunning  bool
	quit       chan bool
	QuitSignal chan bool
	wg         sync.WaitGroup
	interval   time.Duration
	worker     func(w *Worker) bool
}

// 创建一个新的协程对象.
func NewWorker(
	name string, interval time.Duration, worker func(w *Worker) bool,
) *Worker {
	return &Worker{
		Name:       name,
		IsRunning:  false,
		quit:       make(chan bool),
		QuitSignal: make(chan bool),
		wg:         sync.WaitGroup{},
		interval:   interval,
		worker:     worker,
	}
}

func (this *Worker) prepareRun() {
	this.wg.Add(1)
	this.IsRunning = true
}

func (this *Worker) prepareStop() {
	this.IsRunning = false
	this.wg.Done()
	this.QuitSignal <- true
}

// 运行.
func (this *Worker) Run() {
	if this.IsRunning {
		return
	}

	this.wg.Add(1)
	this.IsRunning = true
	go func() {
		for {
			select {
			case <-this.quit:
				this.prepareStop()
				return
			default:
				// 先运行再等待.
				if !this.worker(this) {
					this.prepareStop()
					return
				}
				select {
				case <-this.quit:
					this.prepareStop()
					return
				case <-time.After(this.interval):
				}
			}
		}
	}()
}

// 停止.
func (this *Worker) Stop() {
	if !this.IsRunning {
		return
	}

	this.quit <- true
	this.wg.Wait()
}
