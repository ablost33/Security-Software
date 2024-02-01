package port_scanner

import (
	"net"
	"strconv"
	"sync"
	"time"
)

/*
	https://tutorialedge.net/projects/building-security-tools-in-go/building-port-scanner-go/
*/

type ScanResult struct {
	Port    string
	State   State
	Service string
}

func scanPort(wg *sync.WaitGroup, protocol, hostname string, port int, results chan<- ScanResult) {
	defer wg.Done()

	result := ScanResult{Port: strconv.Itoa(port) + "/" + protocol}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)
	if err != nil {
		result.State = Closed
		results <- result
		return
	}
	defer conn.Close()
	result.State = Open
	results <- result
}

func scan(hostname string, startPort, endPort int) []ScanResult {
	var wg sync.WaitGroup
	results := make(chan ScanResult, endPort-startPort+1)

	for i := startPort; i <= endPort; i++ {
		wg.Add(2)
		go scanPort(&wg, "udp", hostname, i, results)
		go scanPort(&wg, "tcp", hostname, i, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var scanResults []ScanResult
	for result := range results {
		scanResults = append(scanResults, result)
	}
	return scanResults
}

func InitialScan(hostname string) []ScanResult {
	return scan(hostname, 0, 1024)
}

func WideScan(hostname string) []ScanResult {
	return scan(hostname, 1025, 49152)
}
