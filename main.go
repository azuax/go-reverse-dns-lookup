package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

func incIP(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] > 0 {
			break
		}
	}
}

func getIPbyCIDR(cidr string) (ips []string, err error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); incIP(ip) {
		ips = append(ips, ip.String())
	}

	// we remove nw and broadcast address
	switch {
	case len(ips) < 2:
		return ips, nil

	default:
		return ips[1 : len(ips)-1], nil
	}

}

func worker(ips chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ip := <-ips
	lookup, err := net.LookupAddr(ip)
	if err != nil {
		// fmt.Println(err)
		return
	}
	for _, addr := range lookup {
		fmt.Printf("%s\t\t\t%s\n", ip, addr)
	}

}

func main() {
	params := ParseOpts(os.Args)

	WG := new(sync.WaitGroup)
	ips := make(chan string, params.nThreads)

	ipRange, err := getIPbyCIDR(params.cidr)
	if err != nil {
		panic("There was an error on the CIDR")
	}

	for _, ip := range ipRange {
		ips <- ip
		WG.Add(1)
		go worker(ips, WG)
	}

	WG.Wait()

}
