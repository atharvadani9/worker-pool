package worker

import "sync"

func Start(ch chan func(), wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range ch {
		job()
	}
}
