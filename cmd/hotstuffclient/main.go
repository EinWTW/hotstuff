package main

import (
	"context"
	"errors"
	"fmt"
	"io"

	"math"
	"os"

	"time"

	client "github.com/EinWTW/hotstuff/cmd/hotstuffclient/client"
	"github.com/EinWTW/hotstuff/internal/profiling"
	"github.com/spf13/pflag"
)

func usage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println()
	fmt.Println("Loads configuration from ./hotstuff.toml")
	fmt.Println()
	fmt.Println("Options:")
	pflag.PrintDefaults()
}

func main() {

	pflag.Usage = usage

	help := pflag.BoolP("help", "h", false, "Prints this text.")
	configFile := flag.String("config", "", "The path to the config file")
	cpuprofile := pflag.String("cpuprofile", "", "File to write CPU profile to")
	memprofile := pflag.String("memprofile", "", "File to write memory profile to")
	pflag.Uint32("self-id", 0, "The id for this replica.")
	pflag.Int("rate-limit", 0, "Limit the request-rate to approximately (in requests per second).")
	pflag.Int("payload-size", 0, "The size of the payload in bytes")
	pflag.Uint64("max-inflight", 10000, "The maximum number of messages that the client can wait for at once")
	pflag.String("input", "", "Optional file to use for payload data")
	pflag.Bool("benchmark", false, "If enabled, a BenchmarkData protobuf will be written to stdout.")
	pflag.Int("exit-after", 0, "Number of seconds after which the program should exit.")
	pflag.Bool("tls", false, "Enable TLS")
	pflag.Parse()

	if *help {
		pflag.Usage()
		os.Exit(0)
	}
	profileStop, err := profiling.StartProfilers(*cpuprofile, *memprofile, "", "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start profilers: %v\n", err)

	}
	defer func() {
		err := profileStop()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to stop profilers: %v\n", err)
			os.Exit(1)
		}
	}()

	client, err := client.InitHotstuffClient(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to InitHotstuffClient: %v\n", err)
		os.Exit(1)
	}
	data := []byte("")
	err = client.SendCommands(context.Background(), data)
	if err != nil && !errors.Is(err, io.EOF) {
		fmt.Fprintf(os.Stderr, "Failed to send commands: %v\n", err)
		client.Close()
		os.Exit(1)
	}
	client.Close()

	stats := client.GetStats()
	throughput := stats.Throughput
	latency := stats.LatencyAvg / float64(time.Millisecond)
	latencySD := math.Sqrt(stats.LatencyVar) / float64(time.Millisecond)

	fmt.Printf("Throughput (ops/sec): %.2f, Latency (ms): %.2f, Latency Std.dev (ms): %.2f\n",
		throughput,
		latency,
		latencySD,
	)

}
