package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	m        sync.Mutex
	resource chan io.Closer
	factory  func() (io.Closer, error)
	closed   bool
}

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size not correct.")
	}
	return &Pool{
		resource: make(chan io.Closer, size),
		factory:  fn,
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resource:
		if !ok {
			return nil, errors.New("pool is closed.")
		}
		return r, nil
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return
	}
	select {
	case p.resource <- r:
		log.Println("release in queue.")
	default:
		log.Println("release of closing")
		r.Close()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	close(p.resource)
	for r := range p.resource {
		r.Close()
	}
}
