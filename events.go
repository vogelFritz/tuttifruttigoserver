package main

import (
	"encoding/json"
	"fmt"
	"net"

	si "github.com/vogelFritz/tuttifruttigoserver/socketinterface"
)

func addEventListeners(server *si.Server) {
	server.AddEventListener("nombreJugador", func(nombre string, socket net.Conn) {

	})
	server.AddEventListener("nuevaSala", func(roomData string, socket net.Conn) {
		server.AddToRoom(roomData, socket)
		fmt.Println(roomData)
		var s sala
		json.Unmarshal([]byte(roomData), &s)
		fmt.Println(s.Nombre)
		server.Emit(si.EmissionParams{
			Event: "nuevaSala",
			Data:  s.Nombre,
		})
	})
	server.AddEventListener("unirse", func(roomName string, socket net.Conn) {
		server.AddToRoom(roomName, socket)
		s := sala{Nombre: roomName, Jugadores: []string{"pedro", "juan"}}
		encodedData, _ := json.Marshal(s)
		server.Emit(si.EmissionParams{
			Room:  roomName,
			Event: "unirse",
			Data:  string(encodedData),
		})
	})
}
