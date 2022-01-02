package server

import (
	"encoding/json"
	"net/http"
)

type createRoom struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (csr *ChatInfo) listRooms(writer http.ResponseWriter, _ *http.Request) {
	http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

}

func (csr *ChatInfo) newRoom(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var room createRoom
	err := decoder.Decode(&room)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if csr.chatter.AllRooms().Find(room.Name) != nil {
		http.Error(writer, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	newRoom := csr.chatter.NewRoom(room.Name, room.Type)
	marshal, err := json.Marshal(newRoom.Info())
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, _ = writer.Write(marshal)
}
