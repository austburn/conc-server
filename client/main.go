package main

import "fmt"
import "os"
import "net"

func main() {
	conn, err := net.Dial("tcp", ":4321")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		fmt.Print("> ")
		inputBuffer := make([]byte, 256)
		n, err := os.Stdin.Read(inputBuffer)
		if err != nil {
			panic(err)
		}
		conn.Write(inputBuffer[:n-1])

		go waitForData(conn)
	}
}

func waitForData(conn net.Conn) {
	for {
		response := make([]byte, 256)
		_, err := conn.Read(response)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("\n%s\n", string(response))
		}
	}
}
