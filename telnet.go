package main

import (
	"fmt"
	"net"
	"os"
	"time"

	// local
	"telnet/vars"
)

func showVersion() {
	fmt.Printf("%s", vars.ClearScreen)
	fmt.Printf("\n%s" + vars.MyProgname + " version: " + vars.MyVersion + "%s\n",
		vars.Yellow, vars.Off,
	)
	os.Exit(0)
}

func showHelp() {
	//fmt.Printf("%s", vars.ClearScreen)
	fmt.Printf("%s" + vars.MyInfo + "%s\n", vars.Blue, vars.Off)
	fmt.Printf("%s" + vars.MyDescription + "%s\n", vars.Yellow, vars.Off)
	fmt.Printf("%s " + vars.MyProgname + " -v, to show version%s\n",
		vars.Green, vars.Off,
	)
	fmt.Printf("%s " + vars.MyProgname + " -h, to show this help page%s\n",
		vars.Green, vars.Off,
	)
	fmt.Printf("%sUsage: " + vars.MyProgname + " <hostname or ip-address> <port number>%s\n",
		vars.Green, vars.Off,
	)
	fmt.Printf("%sExample: " + vars.MyProgname + " myShibaInu.com 443%s\n\n",
		vars.Green, vars.Off,
	)
	os.Exit(0)
}

func main() {
	timeout := 2 * time.Second
	if len(os.Args) < 2 {
		fmt.Printf("%ssyntax error...%s\n\n", vars.Red, vars.Off)
		showHelp()
	}
	switch os.Args[1] {
		case "-v":
			showVersion()
		case "-h":
			showHelp()
	}

	// read given argument
	if len(os.Args) != 3 {
		showHelp()
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
