package brain

import (
	"encoding/json"
	"log"
)

type ActionType string

const (
	NewConnection   ActionType = "new-connection"
	NotifCreateUser ActionType = "notif-create-user"
	UpdateChatRooms            = "update-chat-rooms"

	NewTextAction   = "new-text"
	JoinRoomAction  = "join-room"
	LeaveRoomAction = "leave-room"

	// TODO unk no match
	NotifNewRoom = "notif-new-room"
)

type Message struct {
	UUID           string `json:"uuid"`
	ConnectionUuid string `json:"connection_uuid"`
	UserId         UserId `json:"user_id"`
	Action         string `json:"action"`

	RoomId RoomID  `json:"room_id"`
	Sender *Client `json:"sender"`

	ChatMessage `json:"chat"`
	ChatRooms   []ChatRoomDto `json:"chatRooms"`
}

func (message *Message) encode() []byte {
	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}

	return jsonData
}
