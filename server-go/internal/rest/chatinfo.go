package server

import (
	"github.com/TheBigBadWolfClub/big-bad-chat/server-go/internal/brain"
	"github.com/go-chi/chi/v5"
)

type ChatInfo struct {
	chatter brain.Chatter
}

func ChatServerEndpoint(chat brain.Chatter) func(router chi.Router) {
	root := ChatInfo{chatter: chat}

	return func(router chi.Router) {

		router.Get("/rooms", root.listRooms)
		router.Post("/rooms", root.newRoom)

		router.Get("/users", root.listUsers)
		router.Post("/users", root.newUser)
	}
}
