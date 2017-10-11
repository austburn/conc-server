package main

import (
	"io"
	"math/rand"
	"net"
	"reflect"
)

type Client struct {
	ID          []byte
	BroadcastCh chan []byte
	Conn        net.Conn
}

type Clients []Client

func (clients *Clients) broadcast(id []byte, msg []byte) {
	for _, c := range *clients {
		// c.Conn.Write(<-c.BroadcastCh)
		if !reflect.DeepEqual(c.ID, id) {
			c.Conn.Write(msg)
		}
	}
}

func randId(b *[]byte) {
	c := make([]byte, len(*b))
	for i := range c {
		c[i] = byte(rand.Intn(26) + 65)
	}
	*b = c
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
		randId(&id)
		client := Client{ID: id, BroadcastCh: broadcastCh, Conn: conn}
		clients = append(clients, client)
		go handleConnection(client, &clients)
	}
}

func handleConnection(client Client, clients *Clients) {
	for {
		inputBuffer := make([]byte, 256)
		n, err := client.Conn.Read(inputBuffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		} else {
			divider := []byte{':', ' '}
			msg := append(client.ID, divider...)
			msg = append(msg, inputBuffer[:n]...)
			client.Conn.Write(msg)
			clients.broadcast(client.ID, msg)
		}
	}
}
