package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

/*
TO DO:
x - Single port scan flag
x - few ports flag (array)
x - debug mode/all (shows all results)
x - add timeout flag

x -Range scan -> add range flags and arguments
x- Show all options for scan protocols
x - add file export for results
x - add custom export file name
*/

// ScanResult is a custom type that shows the port number and its state(Open/closed)
type ScanResult struct {
	Port  int64
	State string
}

var (
	testing = kingpin.Command("testing", "for testing")
	scan    = kingpin.Command("scan", "Scans ports on a specific target")

	all      = kingpin.Flag("all", "Display all results. By default it will only display the open ports").Short('a').Bool()
	export   = kingpin.Flag("export", "Results will be exporeted to a file").Bool()
	specific = scan.Flag("specific", "Scans only a few specific ports that you specifed").Short('s').Bool()

	exportFileName = kingpin.Flag("name", "Choose a name for the result file export, including the file extension.").PlaceHolder("output.txt").Default("output.txt").String()

	target = scan.Arg("target", "What target do you want to scan").Required().String()
	ports  = scan.Arg("ports", "What ports do you want to scan").Int64List()

	protocol = scan.Flag("protocol", "What protocol you want to use. Default is set to tcp. Options are: tcp, tcp4 (IPv4-only), tcp6 (IPv6-only), udp, udp4 (IPv4-only), udp6 (IPv6-only), ip, ip4 (IPv4-only), ip6 (IPv6-only), unix, unixgram and unixpacket").PlaceHolder("tcp").Short('p').Default("tcp").String()
	timeout  = scan.Flag("timeout", "Set the connection timeout. Amount of seconds").PlaceHolder("10s").Duration()

	start = scan.Flag("start", "Using a port range scan, whats the port you want to start with").Default("1").PlaceHolder("1").String()
	end   = scan.Flag("end", "Using a port range scan, whats the port you want to end with").Default("1024").PlaceHolder("1024").String()
	//drugoStanje = kingpin.Command("drugo", "heheh test 2")
)

func main() {
	var result ScanResult

	switch kingpin.Parse() {
	case "scan":

		if *timeout == 0 {
			*timeout = 10 * time.Second
		}

		if *export == true {

			if _, err := os.Stat(*exportFileName); err == nil {
				fmt.Printf("File named %v already exists. Remove the file or change the name with --name\n", *exportFileName)
				os.Exit(0)
			}

			file, err := os.Create(*exportFileName)
			if err != nil {
				return
			}

			defer file.Close()

			fmt.Print(Letters())
			fmt.Printf("** Target: %v ** \n", *target)

			file.WriteString(Letters())
			file.WriteString(fmt.Sprintf("** Target: %v ** \n", *target))

			//specific scan
			if *specific == true {
				if len(*ports) == 0 {
					*ports = append(*ports, 22, 80, 8080)
				}

				fmt.Println("Scanning specific ports:", *ports)
				file.WriteString(fmt.Sprintln("Scanning specific ports:", *ports))

				for i, s := range *ports {
					result = ScanPort(*protocol, *target, s, *timeout)
					if *all == true {
						fmt.Printf("%v - %v:%v - %v \n", i, *target, result.Port, result.State)
						file.WriteString(fmt.Sprintf("%v - %v:%v - %v \n", i, *target, result.Port, result.State))
					} else if result.State == "Open" {
						fmt.Printf("%v - %v:%v - %v \n", i, *target, result.Port, result.State)
						file.WriteString(fmt.Sprintf("%v - %v:%v - %v \n", i, *target, result.Port, result.State))
						//fmt.Printf("Port %v is %v \n", result.Port, result.State)
					}
				}
				//open := ScanPort(*protocol, *target, 22, *timeout)

				//fmt.Println(*timeout)
				//fmt.Printf("Specific port %v is %v \n", open.Port, open.State)

				//port range scan
			} else {
				n, err := strconv.ParseInt(*start, 10, 64)
				if err != nil {
					log.Println(err)
				}
				e, err2 := strconv.ParseInt(*end, 10, 64)
				if err2 != nil {
					log.Println(err2)
				}

				if n > e {
					fmt.Println("End value is bigger than the starting value")
					file.WriteString(fmt.Sprintln("End value is bigger than the starting value"))
				}

				var i int64
				for i = n; i <= e; i++ {
					result = ScanPort(*protocol, *target, i, *timeout)
					if *all == true {
						fmt.Printf("%v:%v - %v \n", *target, result.Port, result.State)
						file.WriteString(fmt.Sprintf("%v:%v - %v \n", *target, result.Port, result.State))
					} else if result.State == "Open" {
						fmt.Printf("%v:%v - %v \n", *target, result.Port, result.State)
						file.WriteString(fmt.Sprintf("%v:%v - %v \n", *target, result.Port, result.State))
						//fmt.Printf("Port %v is %v \n", result.Port, result.State)
					}
				}
			}
		} else {
			fmt.Print(Letters())
			fmt.Printf("** Target: %v ** \n", *target)

			//specific scan
			if *specific == true {
				if len(*ports) == 0 {
					*ports = append(*ports, 22, 80, 8080)
				}

				fmt.Println("Scanning specific ports:", *ports)

				for i, s := range *ports {
					result = ScanPort(*protocol, *target, s, *timeout)
					if *all == true {
						fmt.Printf("%v - %v:%v - %v \n", i, *target, result.Port, result.State)
					} else if result.State == "Open" {
						fmt.Printf("%v - %v:%v - %v \n", i, *target, result.Port, result.State)
						//fmt.Printf("Port %v is %v \n", result.Port, result.State)
					}
				}
				//open := ScanPort(*protocol, *target, 22, *timeout)

				//fmt.Println(*timeout)
				//fmt.Printf("Specific port %v is %v \n", open.Port, open.State)

				//port range scan
			} else {
				n, err := strconv.ParseInt(*start, 10, 64)
				if err != nil {
					log.Println(err)
				}
				e, err2 := strconv.ParseInt(*end, 10, 64)
				if err2 != nil {
					log.Println(err2)
				}

				if n > e {
					fmt.Println("End value is bigger than the starting value")
				}

				var i int64
				for i = n; i <= e; i++ {
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

	case "testing":

		fmt.Println(*specific)
		fmt.Println(*ports)
		fmt.Println(*export)
	}
}

//ScanPort is a function the takes the adress port and the protocol. It gives back the port number and its state trough the ScanResult custom type
func ScanPort(protocol, hostname string, port int64, timeout2 time.Duration) ScanResult {
	//combines adress and the port in one variable
	reult := ScanResult{Port: port}
	portSting := strconv.FormatInt(port, 10)
	adress := hostname + ":" + portSting
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

//Yes I know this is unnecessary but it looks cool so I dont care

//Letters is a simple functions that prints out cool ascii text
func Letters() string {
	l1 := fmt.Sprintln("                                 _                                         ")
	l2 := fmt.Sprintln("                                | |                                        ")
	l3 := fmt.Sprintln("   __ _  ___    _ __   ___  _ __| |_   ___  ___ __ _ _ __  _ __   ___ _ __ ")
	l4 := fmt.Sprintln("  / _` |/ _ \\  | '_ \\ / _ \\| '__| __| / __|/ __/ _` | '_ \\| '_ \\ / _ \\ '__|")
	l5 := fmt.Sprintln(" | (_| | (_) | | |_) | (_) | |  | |_  \\__ \\ (_| (_| | | | | | | |  __/ |   ")
	l6 := fmt.Sprintln("  \\__, |\\___/  | .__/ \\___/|_|   \\__| |___/\\___\\__,_|_| |_|_| |_|\\___|_|   ")
	l7 := fmt.Sprintln("   __/ |       | |                                                         ")
	l8 := fmt.Sprintln("  |___/        |_|                                                         ")
	l9 := fmt.Sprintln("")
	lf := l1 + l2 + l3 + l4 + l5 + l6 + l7 + l8 + l9
	return lf
}

/*
if *export == true {
	if _, err := os.Stat(*ExportFileName); err == nil {
		fmt.Printf("File named %v already exists. Remove the file or change the name with --name\n", *ExportFileName)
		os.Exit(0)
	}
}

if *export == true {
	file, err := os.Create(*ExportFileName)
	if err != nil {
		return
	}
	defer file.Close()

}
*/
