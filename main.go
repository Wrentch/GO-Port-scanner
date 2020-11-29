package main

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

/*
TO DO: Single port scan flag
few ports flag (array)
x - debug mode/all (shows all results)
x - add timeout flag
*/

// ScanResult is a custom type that shows the port number and its state(Open/closed)
type ScanResult struct {
	Port  int
	State string
}

var (
	//test1 = kingpin.Flag("tes1", "test").String()

	scan     = kingpin.Command("scan", "Scans ports on a specific target")
	all      = kingpin.Flag("all", "Display all results. By default it will only display the open ports").Bool()
	target   = scan.Arg("target", "What target do you want to scan").Required().String()
	protocol = scan.Flag("protocol", "What protocol you want to use. Default is set to tcp").PlaceHolder("tcp").Short('p').Default("tcp").String()
	timeout  = scan.Flag("timeout", "Set the connection timeout. Amount of seconds").PlaceHolder("10s").Short('t').Duration()
	//timeoutInt, err = strconv.Atoi(*timeout)

	//drugoStanje = kingpin.Command("drugo", "heheh test 2")
)

func main() {
	var result ScanResult
	switch kingpin.Parse() {
	case "scan":
		if *timeout == 0 {
			*timeout = 10 * time.Second
		}

		open := ScanPort(*protocol, *target, 22, *timeout)
		fmt.Println(*target)
		//fmt.Println(*timeout)
		fmt.Printf("Specific port %v is %v \n", open.Port, open.State)

		for i := 1; i < 1024; i++ {

			result = ScanPort(*protocol, *target, i, *timeout)
			if *all == true {
				fmt.Printf("%v:%v - %v \n", *target, result.Port, result.State)
			} else if result.State == "Open" {
				fmt.Printf("%v:%v - %v \n", *target, result.Port, result.State)
				//fmt.Printf("Port %v is %v \n", result.Port, result.State)
			}

		}

	}
}

//ScanPort is a function the takes the adress port and the protocol. It gives back the port number and its state trough the ScanResult custom type
func ScanPort(protocol, hostname string, port int, timeout2 time.Duration) ScanResult {
	//combines adress and the port in one variable
	reult := ScanResult{Port: port}
	adress := hostname + ":" + strconv.Itoa(port)
	//Sends a connection the the adress at a declared port

	conncect, err := net.DialTimeout(protocol, adress, timeout2)

	//If the connection fails/gives an error that means its closed
	if err != nil {
		reult.State = "Closed"
		return reult
	}

	//closes the conncetion and gives the result
	defer conncect.Close()
	reult.State = "Open"
	return reult

}
