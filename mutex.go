package ctxsync

import (
	"sync"
	"context"
)

type Mutex struct {
	initOnce sync.Once
	ch       chan struct{}
}

func (mu *Mutex) init() {
	mu.initOnce.Do(func() {
		mu.ch = make(chan struct{}, 1)
	})
}

func (mu *Mutex) Lock(ctx context.Context) error {
	mu.init()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case mu.ch <- struct{}{}:
		return nil
	}
}

func (mu *Mutex) Unlock() {
	mu.init()

	select {
	case <-mu.ch:
	default:
	}
}
