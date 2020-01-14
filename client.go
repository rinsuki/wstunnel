package main

import (
	"log"
	"net"

	"github.com/gorilla/websocket"
)

func client(bindIP string, websocketURL string) {
	listener, err := net.Listen("tcp", bindIP)
	if err != nil {
		panic(err)
	}
	for {
		tcp, err := listener.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		go func() {

			defer tcp.Close()
			ws, _, err := websocket.DefaultDialer.Dial(websocketURL, nil)
			if err != nil {
				log.Println(err)
				return
			}
			defer ws.Close()
			go func() {
				buf := make([]byte, 1024)
				for {
					len, err := tcp.Read(buf)
					if err != nil {
						log.Println(err)
						tcp.Close()
						ws.Close()
						break
					}
					log.Printf("C→S %d", len)
					ws.WriteMessage(websocket.BinaryMessage, buf[0:len])
				}
			}()
			for {
				msgType, buf, err := ws.ReadMessage()
				if err != nil {
					log.Println(err)
					tcp.Close()
					ws.Close()
					break
				}
				if msgType != websocket.BinaryMessage {
					log.Println("unknown msgType")
				}
				log.Printf("S→C %d", len(buf))
				tcp.Write(buf)
			}
		}()
	}
}
