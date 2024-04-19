package main

import "github.com/ferizoozoo/echo-tcp-server-in-go/server"

func main() {
	server := server.New("127.0.0.1", "123")
	server.Start()
}
