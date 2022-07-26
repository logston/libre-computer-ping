package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	timeout = time.Duration(1 * time.Second)
)

var (
	connLatency = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "latency",
			Help: "Connection latency (in ms) to address.",
		},
		[]string{"host"},
	)
	connAttempts = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "attempts",
			Help: "Attempted connections to address.",
		},
		[]string{"host"},
	)
	connSuccesses = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "successes",
			Help: "Successful connections to address.",
		},
		[]string{"host"},
	)
)

func main() {
	args := os.Args[1:]
	if err := validateArgs(args); err != nil {
		help()
		panic(err)
	}
	duration := args[0]
	addresses := args[1:]

	d, err := time.ParseDuration(duration)
	if err != nil {
		help()
		panic(err)
	}

	for _, address := range addresses {
		go loop(address, d)
	}

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

func validateArgs(args []string) error {
	if len(args) <= 1 {
		return fmt.Errorf("invalid args")
	}

	for _, address := range args[1:] {
		if !strings.Contains(address, ":") {
			return fmt.Errorf("address %q is not valid", address)
		}
	}

	return nil
}

func help() {
	fmt.Println("./ping <interval> <host>:<port> <host>:<port>...")
}

func loop(address string, d time.Duration) {

	for {
		connAttempts.WithLabelValues(address).Inc()
		latency, err := connect(address, d)
		if err != nil {
			fmt.Printf("connection error: %v\n", err)
		} else {
			connSuccesses.WithLabelValues(address).Inc()
			connLatency.WithLabelValues(address).Set(float64(latency.Milliseconds()))
			fmt.Printf("connected to %s in %dms\n", address, latency.Milliseconds())
		}

		time.Sleep(d)
	}
}

func connect(address string, timeout time.Duration) (*time.Duration, error) {
	s := time.Now()
	c, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return nil, err
	}
	latency := time.Now().Sub(s)
	c.Close()

	return &latency, nil
}
