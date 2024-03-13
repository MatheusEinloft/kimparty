package main

import (
	"kimparty/config"
	"kimparty/internal/party"
	"kimparty/internal/server"
	"kimparty/internal/server/handler"
	"kimparty/internal/websocket"
)

func main() {
	port := config.GetPort()
	server := server.NewServer(port)

	partyService := party.NewService()
	upgrader := websocket.NewUpgrader()

	createPartyHandler := handler.NewCreatePartyHandler(partyService)
	findPartyHandler := handler.NewFindPartyHandler(partyService)
	joinPartyHandler := handler.NewJoinPartyHandler(partyService, upgrader)

	server.AddHandler(createPartyHandler)
	server.AddHandler(findPartyHandler)
	server.AddHandler(joinPartyHandler)

	server.Start()
}
