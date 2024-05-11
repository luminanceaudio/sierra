package maingroup

import (
	"context"
	"errors"
	"sync"
)

type MainGroup struct {
	wg     sync.WaitGroup
	errCh  chan error
	ctx    context.Context
	cancel context.CancelFunc
}

func New(ctx context.Context) *MainGroup {
	ctx, cancel := context.WithCancel(ctx)
	return &MainGroup{
		errCh:  make(chan error),
		ctx:    ctx,
		cancel: cancel,
	}
}

func (mg *MainGroup) Go(fn func(context.Context) error) {
	mg.wg.Add(1)
	go func() {
		defer mg.wg.Done()

		select {
		case <-mg.ctx.Done(): // In case context is already cancelled
			return
		default:
			if err := fn(mg.ctx); err != nil {
				mg.errCh <- err
				mg.cancel()
			}
		}
	}()
}

func (mg *MainGroup) Wait() error {
	// Close the error channel once all go routines finish
	go func() {
		mg.wg.Wait()
		close(mg.errCh)
	}()

	for {
		select {
		// Wait for context to be cancelled
		case <-mg.ctx.Done():
			return errors.New("context cancelled")

		// Wait for an error to be received, or for all routines to finish
		case err, ok := <-mg.errCh:
			if !ok {
				return nil // All routines finished successfully
			}
			return err
		}
	}
}
