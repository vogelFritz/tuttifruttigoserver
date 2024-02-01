package main

import (
	si "github.com/vogelFritz/tuttifruttigoserver/socketinterface"
)

func main() {
	var server si.Server
	server.Init("192.168.0.91:8000")
	addEventListeners(&server)
	server.WaitForClients()
}
