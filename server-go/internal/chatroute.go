package internal

import "net/http"

type ChatRoute interface {
	Path() string
	Handler(http.ResponseWriter, *http.Request)
}
