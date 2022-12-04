// Author: Sam Moran
// Date: 12/3/202
// Description:
/*
	This Go program runs a TCP port scanner. As a refresher, a TCP connection involves a three-way handshake (syn, syn-ack, ack) to establish a connection between a client and server. The connection can result in one of three ways: 

	1) Open Port: The three-way handshake results in an "ack" acknowledgment returned to the target server.

	2) Closed Port: The three-way handshake involves the client sending a "syn" packet but receiveing an "rst" response.

	3) Filtered Port: The "syn" packet is sent by the client, hits the firewall, and times out. 

	Based on the above results, you can determine which of the 65,535 possible ports are open and available to you to connect to. 

	This program runs a simple scan using Go's net package - net.DIal(network, address string). The first argument takes in a protocol (in this case, tcp) and then the target server. In this case, it's a free scanning service provided by the creator of Nmap. IF the connection is successful, "err" will return "nil" (although Dial returns "Conn" and "error", as long as "error" is returned as "nil", we know the connection succeeded).
*/

package main

import(
	"fmt"
	"net"
)

func main() {

	_, err := net.Dial("tcp", "scanme.nmap.org:80")

	if err == nil {

		fmt.Println("Connection Successful!")

	}

}