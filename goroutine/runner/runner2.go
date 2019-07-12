package runner

import (
	"os"
	"os/signal"
	"time"
)

type Runner2 struct {
	tasks     []func(int)
	complete  chan error
	interrupt chan os.Signal
	timeout   <-chan time.Time
}

func New2(duration time.Duration) *Runner2 {
	return &Runner2{
		complete:  make(chan error),
		timeout:   time.After(duration),
		interrupt: make(chan os.Signal),
	}
}

func (r *Runner2) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner2) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner2) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner2) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
