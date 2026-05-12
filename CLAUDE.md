# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

An in-memory worker pool in Go — a job queue processor where workers pull `func()` jobs from a buffered channel and execute them concurrently.

## Commands

```bash
go test ./...          # run all tests
go test -race ./...    # run tests with race detector (preferred)
go test -run TestName  # run a single test
go run example/main.go # run the demo
```

## Architecture

**Data flow:**
```
Submit() → buffered chan (func()) → worker goroutines → execute job
                                                      ↓
Shutdown() → close(chan) → workers drain → WaitGroup.Done() → Shutdown() returns
```

**Files:**
- `pool.go` — `Pool` struct, `New(workers, queueSize int)`, `Submit(job func()) error`, `Shutdown()`
- `worker.go` — per-worker goroutine: ranges over job channel, executes jobs, calls `WaitGroup.Done()`
- `pool_test.go` — concurrency, shutdown ordering, submit-after-shutdown; always run with `-race`
- `example/main.go` — runnable demo

## Key Contracts

- `Submit` returns an error if called after `Shutdown` (guarded by `sync.Once` + closed flag)
- `Shutdown` blocks until all in-flight jobs finish (`sync.WaitGroup`)
- `Shutdown` is idempotent
- Workers drain the channel fully before exiting — no jobs are dropped on shutdown
