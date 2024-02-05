package signalctx

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"
)

// Context 實現關機
// 調用Context.Done()前建議使用以下:
//
//	Context.Increment()
//	defer Context.Decrement()
type Context struct {
	context.Context
	mu      sync.Mutex
	counter int64
	cancel  context.CancelFunc
	done    chan struct{}
}

func (c *Context) Done() <-chan struct{} {
	return c.Context.Done()
}

func (c *Context) Err() error {
	return c.Context.Err()
}

func (c *Context) Value(key interface{}) interface{} {
	return c.Context.Value(key)
}

func (c *Context) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter += 1
}

func (c *Context) Decrement() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter -= 1
	if c.counter == 0 {
		close(c.done)
	}
}

func (c *Context) Cancel() {
	c.cancel()
}

func (c *Context) AllDone() <-chan struct{} {
	return c.done
}

func (c *Context) Shutdown() <-chan os.Signal {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt, os.Kill)
	return shutdownSignal
}

func (c *Context) RunFunc(fn func()) {
	c.Increment()
	defer c.Decrement()

	for {
		select {
		case <-c.Done():
			slog.Info("Context done. Exiting...")
			return
		default:
		}

		fn()
	}
}

func NewContext() *Context {
	ctx := &Context{}
	ctx.Context, ctx.cancel = context.WithCancel(context.Background())
	ctx.done = make(chan struct{})
	return ctx
}
