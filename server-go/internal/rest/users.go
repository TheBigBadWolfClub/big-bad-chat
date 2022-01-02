package server

import (
	"encoding/json"
	"net/http"
)

type createUser struct {
	Name string `json:"name"`
}

func (csr *ChatInfo) listUsers(writer http.ResponseWriter, _ *http.Request) {
	http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (csr *ChatInfo) newUser(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var user createUser
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if csr.chatter.AllClients().FindByName(user.Name) != nil {
		http.Error(writer, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	client := csr.chatter.AddUser(user.Name)
	marshal, err := json.Marshal(client.Info())
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, _ = writer.Write(marshal)
}
