package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

// ScanResult is a custom type that shows the port number and its state(Open/closed)
type ScanResult struct {
	Port  int
	State string
}

func main() {
	//declaring variables
	var result ScanResult
	var hostname string

	fmt.Println("--Port Scanner--")
	//input for the adress that you want to scan
	SInput := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the adress that you want to scan:")
	SInput.Scan()
	hostname = SInput.Text()

	//Here you can scan only 1 specific port, if you wish to scan 2 or more just duplicate the function.
	open := ScanPort("tcp", hostname, 22)
	fmt.Printf("Specific port %v is %v \n", open.Port, open.State)

	//This loop will scan ports from 1 to 1024 and only print the ones that are open in a new line
	for i := 1; i < 1024; i++ {
		result = ScanPort("tcp", hostname, i)
		//If you want to see all the ports results just remove the if statement
		if result.State == "Open" {
			fmt.Printf("Port %v is %v \n", result.Port, result.State)
		}

	}

}

//ScanPort is a function the takes the adress port and the protocol. It gives back the port number and its state trough the ScanResult custom type
func ScanPort(protocol, hostname string, port int) ScanResult {
	//combines adress and the port in one variable
	reult := ScanResult{Port: port}
	adress := hostname + ":" + strconv.Itoa(port)
	//Sends a connection the the adress at a declared port
	conncect, err := net.DialTimeout(protocol, adress, 10*time.Second)

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

//This is a test that I didnt want to delete

//func InitalScan(hostname string) []ScanResult {
//var result []ScanResult
//for i := 1; i < 1024; i++ {
//result = append(result, ScanPort("tcp", hostname, i))
//}

//return result
//}
