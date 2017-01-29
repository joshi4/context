package main

import (
	"time"

	"github.com/joshi4/context"
)

func main() {
	ctx := context.Background()
	timeout(ctx)
	deadline(ctx)
	value(ctx)
	cancel(ctx)
	context.Trail(ctx)
}

func timeout(ctx context.Context) {
	_, cancel := context.WithTimeout(ctx, 10*time.Second)
	cancel()
}

func cancel(ctx context.Context) {
	_, cancel := context.WithCancel(ctx)
	cancel()
}

func value(ctx context.Context) {
	_ = context.WithValue(ctx, "key", "value")
}

func deadline(ctx context.Context) {
	_, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	cancel()
}

// Current output of go run  main.go
//main.main: /Users/shantanu/go/src/github.com/joshi4/context/x/main.go 10
//main.timeout: /Users/shantanu/go/src/github.com/joshi4/context/x/main.go 19
//main.deadline: /Users/shantanu/go/src/github.com/joshi4/context/x/main.go 33
//main.value: /Users/shantanu/go/src/github.com/joshi4/context/x/main.go 29
//main.cancel: /Users/shantanu/go/src/github.com/joshi4/context/x/main.go 24
