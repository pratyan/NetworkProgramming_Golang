package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	//Listen
	//Accept
	//Handle connection -> thread ,, so that multiple clients can connect to this server

	dstream, err := net.Listen("tcp", ":8080") // listener will create a server

	// error check
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dstream.Close()

	// infinite loop to continously looking for connections
	for {
		con, err := dstream.Accept()

		//error check
		if err != nil {
			fmt.Println(err)
			return
		}

		// handle will do the further server stuff
		go handle(con) // goroutine so that multiple connections could be made
	}

}

func handle(con net.Conn) {
	for {
		data, err := bufio.NewReader(con).ReadString() // reading the string from the connection

		// error check
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(data) // printing the data that we have got from our connection

	}

	con.Close() // closing the connection after extracting all the string
}
