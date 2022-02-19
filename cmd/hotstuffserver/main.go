package main

import (
	"fmt"
	"os"

	srv "github.com/EinWTW/hotstuff/cmd/hotstuffserver/clientsrv"
	"github.com/EinWTW/hotstuff/internal/profiling"
	"github.com/spf13/pflag"
)

func usage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println()
	fmt.Println("Loads configuration from ./hotstuff.toml and file specified by --config")
	fmt.Println()
	fmt.Println("Options:")
	pflag.PrintDefaults()
}

func main() {
	pflag.Usage = usage

	// some configuration options can be set using flags
	help := pflag.BoolP("help", "h", false, "Prints this text.")
	configFile := pflag.String("config", "", "The path to the config file")
	cpuprofile := pflag.String("cpuprofile", "", "File to write CPU profile to")
	memprofile := pflag.String("memprofile", "", "File to write memory profile to")
	fullprofile := pflag.String("fullprofile", "", "File to write fgprof profile to")
	traceFile := pflag.String("trace", "", "File to write execution trace to")
	pflag.Uint32("self-id", 0, "The id for this replica.")
	pflag.Int("view-change", 100, "How many views before leader change with round-robin pacemaker")
	pflag.Int("batch-size", 100, "How many commands are batched together for each proposal")
	pflag.Int("view-timeout", 1000, "How many milliseconds before a view is timed out")
	pflag.String("privkey", "", "The path to the private key file")
	pflag.String("cert", "", "Path to the certificate")
	pflag.Bool("print-commands", false, "Commands will be printed to stdout")
	pflag.Bool("print-throughput", false, "Throughput measurements will be printed stdout")
	pflag.Int("interval", 1000, "Throughput measurement interval in milliseconds")
	pflag.Bool("tls", false, "Enable TLS")
	pflag.String("client-listen", "", "Override the listen address for the client server")
	pflag.String("peer-listen", "", "Override the listen address for the replica (peer) server")
	pflag.Parse()

	if *help {
		pflag.Usage()
		os.Exit(0)
	}

	profileStop, err := profiling.StartProfilers(*cpuprofile, *memprofile, *traceFile, *fullprofile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start profilers: %v\n", err)
		os.Exit(1)
	}

	defer func() {
		err := profileStop()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to stop profilers: %v\n", err)
			os.Exit(1)
		}
	}()

	config := ""
	if configFile != nil {
		config = *configFile
	}
	srv.InitHotstuffServer(config)
}
