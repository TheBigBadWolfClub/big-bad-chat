package brain

import "time"

type Type string

const (
	Direct  Type = "Direct"
	Meeting      = "Meeting"
	Public       = "Public"
)

type ChatRoomDto struct {
	ID
	Type `json:"type"`
}

type ToChatRoomDto interface {
	Dto() ChatRoomDto
}

type ID struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type UserId struct {
	ID
	ConnectionUuid string    `json:"connection_uuid"`
	LastConnection time.Time `json:"last_connection"`
}

func (u UserId) Dto() ChatRoomDto {
	return ChatRoomDto{
		ID: ID{
			UUID: u.UUID,
			Name: u.Name,
		},
		Type: Direct,
	}
}

type RoomID struct {
	ID
	Type `json:"type"`
}

func (r RoomID) Dto() ChatRoomDto {
	return ChatRoomDto{
		ID: ID{
			UUID: r.UUID,
			Name: r.Name,
		},
		Type: r.Type,
	}
}
