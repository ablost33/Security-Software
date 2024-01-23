package main

import (
	"fmt"
	port_scanner "github.com/ablost33/Security-Software/port-scanner"
)

func main() {
	results := port_scanner.InitialScan("localhost")
	fmt.Println(results)
	fmt.Println("============================")
	results = port_scanner.WideScan("localhost")
	fmt.Println(results)
}
