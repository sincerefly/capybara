package fileitem

import (
	"github.com/sincerefly/capybara/base/goroutine_pool"
)

type ExecutorFunc func(fi FileItem) error

func PoolExecutor(store *Store, runner ExecutorFunc) {
	pool := goroutine_pool.NewPool()
	for _, fi := range store.GetItems() {
		innerFi := fi
		pool.Submit(func() {
			_ = runner(innerFi)
		})
	}

	// Stop the pool and wait for all submitted tasks to complete
	pool.StopAndWait()
}

func LoopExecutor(store *Store, runner ExecutorFunc) {
	for _, fi := range store.GetItems() {
		_ = runner(fi)
	}
}
