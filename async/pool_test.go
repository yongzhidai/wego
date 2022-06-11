package async

import (
	"fmt"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := NewRoutinePool("test", 10, 10)
	pool.Execute(func() {
		fmt.Println("任务一开始执行")
		time.Sleep(time.Second)
		fmt.Println("任务一开始结束")
	})
	pool.Execute(func() {
		fmt.Println("任务二开始执行")
		time.Sleep(time.Second)
		fmt.Println("任务二开始结束")
	})

	pool.Shutdown()

	//防止因程序退出线程池里的协程退出
	time.Sleep(10 * time.Second)
}
