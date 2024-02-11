package main

import (
	"encoding/json"
	"math/rand"

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
	server.AddEventListener("fieldSuggestion", func(fieldsSuggestion string, socket si.Socket) {
		server.Emit(si.EmissionParams{
			Room:  socket.Room,
			Event: "fieldSuggestion",
			Data:  fieldsSuggestion,
		})
	})
	server.AddEventListener("voted", func(voteData string, socket si.Socket) {
		server.Emit(si.EmissionParams{
			Room:  socket.Room,
			Event: "voted",
			Data:  voteData,
		})
	})
	server.AddEventListener("startGame", func(_ string, socket si.Socket) {
		server.Emit(si.EmissionParams{
			Room:  socket.Room,
			Event: "startGame",
		})
		newLetter := string(rune(65 + rand.Int()%26))
		server.Emit(si.EmissionParams{
			Room:  socket.Room,
			Event: "newLetter",
			Data:  newLetter,
		})
	})
	server.AddEventListener("stop", func(userData string, socket si.Socket) {
		server.Emit(si.EmissionParams{
			Room:  socket.Room,
			Event: "stop",
			Data:  userData,
		})
	})
	server.AddEventListener("userFieldValues", func(userData string, socket si.Socket) {
		server.Emit(si.EmissionParams{
			Room:  socket.Room,
			Event: "userFieldValues",
			Data:  userData,
		})
	})
}
