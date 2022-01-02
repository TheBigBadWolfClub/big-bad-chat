package main

import (
	"github.com/TheBigBadWolfClub/big-bad-chat/server-go/internal"
	"github.com/TheBigBadWolfClub/big-bad-chat/server-go/internal/brain"
	server "github.com/TheBigBadWolfClub/big-bad-chat/server-go/internal/rest"

	"net/http"
)

func main() {

	httpRouter := internal.NewHTTPRouter().Mux

	// websocket
	chatter := brain.NewChatter()
	clientServer := brain.NewClientServer(chatter)

	httpRouter.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		clientServer.NewClientConnection(writer, request)
	})

	// static
	site := http.FileServer(http.Dir("./public"))
	httpRouter.Handle("/*", site)

	// rest
	rest := server.ChatServerEndpoint(chatter)
	httpRouter.Route("/api", rest)

	_ = http.ListenAndServe(":8080", httpRouter)
}
