# SimpleWorker
golang 协程的简单封装, 可以随意启动停止, 并设置运行间隔.

示例:
```go
package main

import (
	"fmt"
	"time"

	"github.com/funte/SimpleWorker"
)

func main() {
	count := 0
  // 创建一个名叫 wahaha 的协程对象, 一秒运行一次.
	worker := SimpleWorker.NewWorker("wahaha", time.Second,
		func(w *SimpleWorker.Worker) bool {
      // 打印名称.
			fmt.Println(w.Name)
			count++
      // 循环三次自动停止此协程.
			return count < 3
		})
	worker.Run()
	<-worker.QuitSignal
}
```
