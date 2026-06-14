package workerpool

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestBasicSubmit(t *testing.T) {
	p := New(2, 4)
	var count int64

	for range 5 {
		p.Submit(func() {
			atomic.AddInt64(&count, 1)
		})
	}

	p.Shutdown()

	if count != 5 {
		t.Errorf("expected 5, got %d", count)
	}
}

func TestSubmitAfterShutdown(t *testing.T) {
	p := New(2, 4)
	p.Shutdown()
	err := p.Submit(func() {})
	if err == nil {
		t.Errorf("Submit called after Shutdown, error should have been returned")
	}
}

func TestConcurrentSubmits(t *testing.T) {
	p := New(4, 100)
	var count int64
	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 10 {
				p.Submit(func() {
					atomic.AddInt64(&count, 1)
				})
			}
		}()
	}

	wg.Wait()
	p.Shutdown()

	if count != 100 {
		t.Errorf("expected 100, got %d", count)
	}
}
func TestDoubleShutdown(t *testing.T) {
	p := New(2, 4)
	p.Shutdown()
	p.Shutdown()
}

func TestShutdownDrainsJobs(t *testing.T) {
	p := New(2, 4)
	var count int64

	for range 6 {
		p.Submit(func() {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			atomic.AddInt64(&count, 1)
		})
	}

	p.Shutdown()
	if count != 6 {
		t.Errorf("expected 6, got %d", count)
	}
}
