package brain

type ChatClient interface {
	AddUser(name string) *Client
	UserHasConnected(client *Client)
	ClientRooms(client *Client) Rooms
}

type ChatRoom interface {
	NewRoom(name string, roomType string) *Room
	RoomClients(room *Room) Clients
}

type Chatter interface {
	AllRooms() Rooms
	AllClients() Clients
	ChatClient
	ChatRoom
	Dtos
}

func NewChatter() Chatter {
	return &relation{
		rooms:   make(map[*Room]Clients),
		clients: make(map[*Client]Rooms),
	}
}

type relation struct {
	rooms   map[*Room]Clients
	clients map[*Client]Rooms
}

func (r *relation) AddUser(name string) *Client {
	client := r.AllClients().FindByName(name)

	if client == nil {
		client = NewClient(name, r)
		r.clients[client] = NewRooms()
	}

	r.CreateUser(client)
	return client
}

func (r *relation) UserHasConnected(client *Client) {
	r.NewConnection(client)
	r.sendChatInfo()
}

func (r *relation) ClientRooms(client *Client) Rooms {
	return r.clients[client]
}

func (r *relation) NewRoom(name string, roomType string) *Room {
	room := NewRoom(name, roomType, r)
	r.rooms[room] = NewClients()
	r.UpdateRoom(room.info.Dto())
	return room
}

func (r *relation) RoomClients(room *Room) Clients {
	return r.rooms[room]
}

func (r *relation) AllRooms() Rooms {
	all := NewRooms()
	for key := range r.rooms {
		all.Add(key)
	}
	return all
}

func (r *relation) AllClients() Clients {
	all := NewClients()
	for key := range r.clients {
		all.Add(key)
	}
	return all
}

type Dtos interface {
	ClientsDtos() []ChatRoomDto
	ClientRoomsDtos(client *Client) []ChatRoomDto
	RoomsDtos() []ChatRoomDto
	RoomClientsDtos(room *Room) []ChatRoomDto
}

func (r relation) ClientRoomsDtos(client *Client) []ChatRoomDto {
	dtos := make([]ChatRoomDto, 0)
	for k := range r.clients[client].List() {
		dtos = append(dtos, k.info.Dto())
	}
	return dtos
}

func (r relation) ClientsDtos() []ChatRoomDto {
	dtos := make([]ChatRoomDto, 0)
	for k := range r.clients {
		dtos = append(dtos, k.info.Dto())
	}
	return dtos
}

func (r relation) RoomClientsDtos(room *Room) []ChatRoomDto {
	dtos := make([]ChatRoomDto, 0)
	for k := range r.rooms[room].List() {
		dtos = append(dtos, k.info.Dto())
	}
	return dtos
}

func (r relation) RoomsDtos() []ChatRoomDto {

	dtos := make([]ChatRoomDto, 0)
	for k := range r.rooms {
		dtos = append(dtos, k.info.Dto())
	}
	return dtos
}

func (r relation) broadcast(msg *Message) {
	for key := range r.AllClients().List() {
		if !key.IsConnected() {
			break
		}
		msg.UserId = *key.info
		msg.ConnectionUuid = key.ConnectionUuid()
		key.send <- msg.encode()
	}
}

func (r relation) sendChatInfo() {

	dtos := make([]ChatRoomDto, 0)
	for k := range r.AllRooms().List() {
		dtos = append(dtos, k.info.Dto())
	}

	for k := range r.AllClients().List() {
		dtos = append(dtos, k.info.Dto())
	}

	r.UpdateChatRooms(dtos)
}
