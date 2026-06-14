package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	workerpool "workerpool/pool"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Enter the workers and queueSize")
		return
	}
	workers, _ := strconv.Atoi(os.Args[1])
	queueSize, _ := strconv.Atoi(os.Args[2])

	p := workerpool.New(workers, queueSize)

	for i := range 10 {
		p.Submit(func() {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			fmt.Println("Job", i)
		})
	}
	p.Shutdown()
}
