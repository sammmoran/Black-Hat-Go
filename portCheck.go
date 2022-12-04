// Author: Sam Moran
// Date: 12/3/2022
// Description:
/*
	This Go program runs a TCP port scanner. As a refresher, a TCP connection involves a three-way handshake (client sends "syn" packet, server responds with "syn-ack" packet, client responds with "ack" packet) to establish a connection between a client and server. The connection can result in one of three ways: 

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
	"sort"
)

// Edit this to view a smaller or larger port range
const PORT_RANGE = 1024


// The worker will take a channel to receive work from and the WaitGroup will track when that unit of work has been completed by that worker
func worker(ports, results chan int){

	for p :=range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	
	}

}




func main() {


	// Create channel
	// The channel will provide work to a pool of 100 workers.
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	// The WaitGroup is a struct that will let us implement concurrency in a thread-safe way. 
	// Add() will increase the internal counter by one for each port being scanned at the moment.
	// Done() will decrement the counter by one once the port has gone through its scan.
	// Wait() will block execution of the goroutine in which it's called until the internal counter reaches zero.
	// This will ensure that our goroutine waits until the connection actually takes place so that it doesn't complete and exit once the loop finishes.
	// var wg sync.WaitGroup


	for i := 0; i <= cap(ports); i++ {

		go worker(ports, results)

	}

	go func() {

		// Iterate through port range
		for i := 1; i <= PORT_RANGE; i++ {

			ports <- i

		}
	}()

	for i := 0; i < PORT_RANGE; i++ {

		port := <-results

		if port != 0{
			openports = append(openports, port)
		}

	}
	close(ports)
	close(results)
	sort.Ints(openports)
	
	for _, port := range openports {
		fmt.Println("%d open\n", port)
	}

}