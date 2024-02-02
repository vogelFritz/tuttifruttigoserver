package main

import (
	"encoding/json"
	"net"

	si "github.com/vogelFritz/tuttifruttigoserver/socketinterface"
)

func addEventListeners(server *si.Server) {
	server.AddEventListener("nombreJugador", func(nombre string, socket net.Conn) {

	})
	server.AddEventListener("nuevaSala", func(roomData string, socket net.Conn) {
		var s sala
		json.Unmarshal([]byte(roomData), &s)
		server.AddToRoom(s.Nombre, socket)
		server.Emit(si.EmissionParams{
			Event: "nuevaSala",
			Data:  roomData,
		})
	})
	server.AddEventListener("unirse", func(roomData string, socket net.Conn) {
		var jsonMap map[string]string
		json.Unmarshal([]byte(roomData), &jsonMap)
		server.AddToRoom(jsonMap["room"], socket)
		server.Emit(si.EmissionParams{
			Room:  jsonMap["room"],
			Event: "unirse",
			Data:  jsonMap["player"],
		})
	})
}
