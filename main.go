package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 4 {
		switch os.Args[1] {
		case "server":
			server(os.Args[2], os.Args[3])
			os.Exit(0)
		case "client":
			client(os.Args[2], os.Args[3])
			os.Exit(0)
		}
	}
	fmt.Println("Usage: ")
	fmt.Printf("  %s server <bind ip> <tcp ip>\n", os.Args[0])
	fmt.Printf("  %s client <bind ip> <websocket url>\n", os.Args[0])
}
