// Author: Sam Moran
// Date: 12/3/2022
// Description:
/*
	This Go program runs a TCP port scanner. As a refresher, a TCP connection involves a three-way handshake (syn, syn-ack, ack) to establish a connection between a client and server. The connection can result in one of three ways: 

	1) Open Port: The three-way handshake results in an "ack" acknowledgment returned to the target server.

	2) Closed Port: The three-way handshake involves the client sending a "syn" packet but receiveing an "rst" response.

	3) Filtered Port: The "syn" packet is sent by the client, hits the firewall, and times out. 

	Based on the above results, you can determine which of the 65,535 possible ports are open and available to you to connect to. 

	This program runs a simple scan using Go's net package - net.DIal(network, address string). The first argument takes in a protocol (in this case, tcp) and then the target server. In this case, it's a free scanning service provided by the creator of Nmap. IF the connection is successful, "err" will return "nil" (although Dial returns "Conn" and "error", as long as "error" is returned as "nil", we know the connection succeeded).

	Core concepts of Go are also expounded on within the code logic for a beginner's benefit.

*/

package main

import(
	"fmt"
	"net"
	"sync"
)

// Edit this to view a smaller or larger port range
const PORT_RANGE = 1024


func main() {

	// The WaitGroup is a struct that will let us implement concurrency in a thread-safe way. 
	// Add() will increase the internal counter by one for each port being scanned at the moment.
	// Done() will decrement the counter by one once the port has gone through its scan.
	// Wait() will block execution of the goroutine in which it's called until the internal counter reaches zero.
	// This will ensure that our goroutine waits until the connection actually takes place so that it doesn't complete and exit once the loop finishes.
	var wg sync.WaitGroup

	// Iterate through port range
	for i := 1; i <= PORT_RANGE; i++ {

		wg.Add(1)

		go func(j int){

			defer wg.Done()

			// Format address for Go Dial function
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)

			// In the event port is closed or filtered through firewall, have the program continue
			if err != nil {
				return
			}

			// Close connection
			conn.Close()
			fmt.Printf("%d open\n", j)

		}(i)

		wg.Wait()

	}

}