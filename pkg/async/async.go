package async

import (
	"context"
	"fmt"
	"os"
	"sync"
)

type Handle func(ctx context.Context, wg *sync.WaitGroup)

type Async struct {
	ctx context.Context
	wg  *sync.WaitGroup
}

func NewAsync(ch <-chan os.Signal) *Async {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-ch
		fmt.Println("准备退出")
		cancel()
	}()
	var wg sync.WaitGroup
	return &Async{ctx: ctx, wg: &wg}
}

func (a *Async) Register(f Handle) {
	a.wg.Add(1)

	go f(a.ctx, a.wg)
}

func (a *Async) Wait() {
	a.wg.Wait()
}
