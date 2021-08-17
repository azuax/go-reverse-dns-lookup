package main

import (
	"flag"
	"fmt"
	"os"
)

type Params struct {
	cidr     string
	nThreads int
}

func ParseOpts(args []string) (PARAMS Params) {
	if len(args) <= 3 || args[1] != "-cidr" {
		fmt.Printf("Usage %s -cidr <CIDR RANGE>", args[0])
		os.Exit(1)
	}
	flag.StringVar(&PARAMS.cidr, "cidr", "", "IP range in CIDR format. Example: 192.168.0.1/24")
	flag.IntVar(&PARAMS.nThreads, "t", 20, "Max number of concurrent requests")
	flag.Parse()

	return
}
