package brain

import (
	"fmt"
	"github.com/google/uuid"
)

func MsgClient(actionType ActionType, client *Client) (*Message, error) {
	switch actionType {
	case NewConnection:
		return newConnection(client), nil
	case NotifCreateUser:
		return UpdateUser(client), nil
	default:
		return nil, fmt.Errorf("unsuported message type for client: %s", actionType)
	}
}

func MsgChatRoomDto(actionType ActionType, collection []ChatRoomDto) (*Message, error) {
	switch actionType {
	case UpdateChatRooms:
		return updateRooms(collection), nil
	default:
		return nil, fmt.Errorf("unsuported message type for rooms: %s", actionType)
	}
}

func updateRooms(collection []ChatRoomDto) *Message {
	return &Message{
		UUID:      uuid.NewString(),
		Action:    UpdateChatRooms,
		ChatRooms: collection,
	}
}

func newConnection(client *Client) *Message {
	return &Message{
		ConnectionUuid: client.ConnectionUuid(),
		UUID:           uuid.NewString(),
		Action:         string(NewConnection),
		UserId:         *client.info,
	}
}

func UpdateUser(client *Client) *Message {
	return &Message{
		ConnectionUuid: client.ConnectionUuid(),
		UUID:           uuid.NewString(),
		Action:         string(NotifCreateUser),
		UserId:         *client.info,
		ChatRooms:      client.chatter.ClientRoomsDtos(client),
	}
}
