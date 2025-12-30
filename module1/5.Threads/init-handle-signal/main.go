package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("init-handle-signal started")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM)
	fmt.Println("init-handle-signal waiting for signal")
	<-sig
	fmt.Println("init-handle-signal received signal")
}
