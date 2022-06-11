package async

import (
	"fmt"
	"strconv"
)

type RoutinePool struct {
	queue chan func()
}

func NewRoutinePool(name string, queueSize int, routineNum int) *RoutinePool {
	pool := &RoutinePool{queue: make(chan func(), queueSize)}
	for i := 0; i < routineNum; i++ {
		go func(index int) {
			for task := range pool.queue {
				task()
			}
			fmt.Println(name + "-" + strconv.Itoa(index) + "协程退出")
		}(i)
	}
	return pool
}

func (pool *RoutinePool) Execute(task func()) {
	pool.queue <- task
}
func (pool *RoutinePool) Shutdown() {
	close(pool.queue)
}
