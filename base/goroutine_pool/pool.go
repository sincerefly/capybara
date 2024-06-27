package goroutine_pool

import "github.com/alitto/pond"

func NewPool() *pond.WorkerPool {
	return pond.New(10, 0, pond.MinWorkers(10))
}
