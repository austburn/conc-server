package main

import "fmt"
import "os"

func main() {
	for {
		fmt.Print("> ")
		inputBuffer := make([]byte, 256)
		n, err := os.Stdin.Read(inputBuffer)
		if err != nil {
			panic(err)
		}
		fmt.Println("Echo", string(inputBuffer[:n-1]))
	}
}
