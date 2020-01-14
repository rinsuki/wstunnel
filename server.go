package main

import (
	"log"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Server struct {
	DestAddress string
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	tcp, err := net.Dial("tcp", s.DestAddress)
	if err != nil {
		log.Println(err)
		return
	}
	defer tcp.Close()
	go func() {
		buf := make([]byte, 1024)
		for {
			len, err := tcp.Read(buf)
			if err != nil {
				log.Println(err)
				conn.Close()
				tcp.Close()
				break
			}
			log.Printf("S→C %d", len)
			conn.WriteMessage(websocket.BinaryMessage, buf[0:len])
		}
	}()
	for {
		msgType, buf, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			conn.Close()
			tcp.Close()
			break
		}
		if msgType != websocket.BinaryMessage {
			log.Println("unknown msgType")
		}
		log.Printf("C→S %d", len(buf))
		tcp.Write(buf)
	}
}

func server(bindIP string, destIP string) {
	s := Server{
		DestAddress: destIP,
	}
	http.HandleFunc("/ws", s.handler)
	log.Fatal(http.ListenAndServe(bindIP, nil))
}
