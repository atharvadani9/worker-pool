package workerpool

import (
	"errors"
	"sync"
	"workerpool/worker"
)

type Pool struct {
	jobs     chan func()
	wg       sync.WaitGroup
	shutdown bool
	once     sync.Once
	mu       sync.Mutex
}

func New(workers, queueSize int) *Pool {
	p := Pool{
		jobs: make(chan func(), queueSize),
	}

	for i := 0; i < workers; i++ {
		p.wg.Add(1)
		go worker.Start(p.jobs, &p.wg)
	}
	return &p
}

func (p *Pool) Submit(job func()) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.shutdown {
		return errors.New("Pool is full, wait your turn")
	}
	p.jobs <- job
	return nil
}

func (p *Pool) Shutdown() {
	p.once.Do(func() {
		p.mu.Lock()
		p.shutdown = true
		close(p.jobs)
		p.mu.Unlock()
		p.wg.Wait()
	})
}
