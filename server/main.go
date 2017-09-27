package main

import (
	"crypto/rand"
	"net"
	"reflect"
)

type Client struct {
	ID          []byte
	BroadcastCh chan []byte
	Conn        net.Conn
}

type Clients []Client

func (clients Clients) broadcast(id []byte) {
	for _, c := range clients {
		// c.Conn.Write(<-c.BroadcastCh)
		if !reflect.DeepEqual(c.ID, id) {
			c.Conn.Write([]byte("test"))
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":4321")
	if err != nil {
		panic(err)
	}
	clients := Clients{}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		broadcastCh := make(chan []byte)
		id := make([]byte, 8)
		rand.Read(id)
		client := Client{ID: id, BroadcastCh: broadcastCh, Conn: conn}
		clients = append(clients, client)
		go handleConnection(client)
		clients.broadcast(client.ID)
	}
}

func handleConnection(client Client) {
	for {
		inputBuffer := make([]byte, 256)
		n, err := client.Conn.Read(inputBuffer)
		if err != nil {
			panic(err)
		} else {
			client.Conn.Write(inputBuffer[:n])
		}
	}
}
