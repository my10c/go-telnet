package main

import (
	"fmt"
	"net"
	"os"
	"time"

	// local
	"telnet/vars"
)

func main() {
	timeout := 2 * time.Second
	// read given argument
	if len(os.Args) != 3 {
		fmt.Println("syntax: host port")
		return
	}
	host := os.Args[1]
	port := os.Args[2]
	address := host + ":" + port
	_, err := net.LookupHost(host)
	if err != nil {
		fmt.Printf("%sError: %v%s\n", vars.Red, err, vars.Off)
		os.Exit(1)
	}
	fmt.Printf("%sConnecting to %s...%s\n", vars.Purple, address, vars.Off)
	// conn, err := net.Dial("tcp", address)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		fmt.Printf("%sError: %v%s\n", vars.Red, err, vars.Off)
		fmt.Printf("%sGives address:port (%s) might not be open%s\n", vars.Red, address, vars.Off)
		os.Exit(1)
	} else {
		fmt.Printf("%sGives address:port (%s) is open%s\n", vars.Green,  address, vars.Off)
	}
	defer conn.Close()
	os.Exit(0)
}
