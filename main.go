package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	timeout = time.Duration(1 * time.Second)
)

func main() {
	args := os.Args[1:]
	if err := validateArgs(args); err != nil {
		help()
		panic(err)
	}

	host, port, duration := args[0], args[1], args[2]

	d, err := time.ParseDuration(duration)
	if err != nil {
		help()
		panic(err)
	}

	for {
		latency, err := connect(host, port, d)
		if err != nil {
			fmt.Printf("connection error: %v\n", err)
		} else {
			fmt.Printf("connected in %dms\n", latency.Milliseconds())
		}
		time.Sleep(d)
	}
}

func validateArgs(args []string) error {
	if len(args) != 3 {
		return fmt.Errorf("invalid args")
	}

	return nil
}

func help() {
	fmt.Println("./ping <host> <port> <interval>")
}

func connect(host, port string, timeout time.Duration) (*time.Duration, error) {
	s := time.Now()
	c, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", host, port), timeout)
	if err != nil {
		return nil, err
	}
	latency := time.Now().Sub(s)
	c.Close()

	return &latency, nil
}
