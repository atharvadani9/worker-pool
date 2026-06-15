# Worker Pool

An in-memory worker pool in Go. Submit jobs as `func()` values; a fixed pool of goroutines executes them concurrently from a buffered queue.

## Usage

```go
pool := workerpool.New(4, 100) // 4 workers, queue capacity 100

pool.Submit(func() {
    fmt.Println("job executed")
})

pool.Shutdown() // blocks until all in-flight jobs finish
```

## API

| Method | Description |
|--------|-------------|
| `New(workers, queueSize int) *Pool` | Create a new pool and start workers |
| `Submit(job func()) error` | Enqueue a job; returns error if pool is shut down or queue is full |
| `Shutdown()` | Stop accepting jobs and wait for all in-flight jobs to complete |

## Running

```bash
go test -race ./...    # run tests with race detector
go run example/main.go # run the demo
```
