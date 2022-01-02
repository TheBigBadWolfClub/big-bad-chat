package brain

import "log"

type Notificator interface {
	NewConnection(client *Client)
	UpdateChatRooms(collection []ChatRoomDto)
	UpdateRoom(dto ChatRoomDto)
	CreateUser(client *Client)
}

func (r relation) NewConnection(client *Client) {
	msgClient, _ := MsgClient(NewConnection, client)
	r.broadcast(msgClient)
}

func (r relation) CreateUser(client *Client) {
	msgClient, _ := MsgClient(NotifCreateUser, client)
	r.broadcast(msgClient)

	dtos := []ChatRoomDto{client.Info().Dto()}
	r.UpdateChatRooms(dtos)
}

func (r relation) UpdateChatRooms(collection []ChatRoomDto) {
	msg, err := MsgChatRoomDto(UpdateChatRooms, collection)
	if err != nil {
		log.Printf("fail to create message UpdateChatRooms: %s\n", UpdateChatRooms)
		return
	}
	r.broadcast(msg)
}

func (r *relation) UpdateRoom(dto ChatRoomDto) {
	dtos := []ChatRoomDto{dto}
	r.UpdateChatRooms(dtos)
}
