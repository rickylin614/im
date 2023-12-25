package signalctx

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
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
	atomic.AddInt64(&c.counter, 1)
}

func (c *Context) Decrement() {
	if atomic.AddInt64(&c.counter, -1) == 0 {
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

func NewContext() *Context {
	ctx := &Context{}
	ctx.Context, ctx.cancel = context.WithCancel(context.Background())
	ctx.done = make(chan struct{})
	return ctx
}
