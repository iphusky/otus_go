package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := start(); err != nil {
		log.Fatal(err)
	}
}

func start() error {
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	args := flag.Args()

	if len(args) != 2 {
		return ErrNotEnoughArgs
	}

	address := net.JoinHostPort(args[0], args[1])

	client := NewTelnetClient(address, *timeout, os.Stdin, os.Stdout)

	if err := client.Connect(); err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "...Connected to %s\n", address)
	defer client.Close()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	go func() {
		if err := client.Send(); err != nil {
			log.Println("send error: ", err)
		}
		fmt.Fprintf(os.Stderr, "...EOF")
		stop()
	}()

	go func() {
		if err := client.Receive(); err != nil {
			log.Println("receive error: ", err)
		}
		fmt.Fprintf(os.Stderr, "...Connection was closed by peer")
		stop()
	}()

	<-ctx.Done()

	return nil
}
