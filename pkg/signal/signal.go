package signal

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func WithContext(ctx context.Context) context.Context {
	nCtx, cancel := context.WithCancel(ctx)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(c)
		select {
		case <-nCtx.Done():
		case <-c:
			cancel()
		}
	}()
	return nCtx
}
