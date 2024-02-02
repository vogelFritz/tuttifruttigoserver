package main

import (
	"encoding/json"

	si "github.com/vogelFritz/tuttifruttigoserver/socketinterface"
)

func addEventListeners(server *si.Server) {
	server.AddEventListener("nuevaSala", func(roomData string, socket si.Socket) {
		var s sala
		json.Unmarshal([]byte(roomData), &s)
		server.AddToRoom(s.Nombre, socket)
		server.Emit(si.EmissionParams{
			Event: "nuevaSala",
			Data:  roomData,
		})
	})
	server.AddEventListener("unirse", func(roomData string, socket si.Socket) {
		var jsonMap map[string]string
		json.Unmarshal([]byte(roomData), &jsonMap)
		server.AddToRoom(jsonMap["room"], socket)
		server.Emit(si.EmissionParams{
			Room:  jsonMap["room"],
			Event: "unirse",
			Data:  jsonMap["player"],
		})
	})
	server.AddEventListener("ready", func(playerName string, socket si.Socket) {
		server.Emit(si.EmissionParams{
			Room:  socket.Room,
			Event: "ready",
			Data:  playerName,
		})
	})
}
