package brain

type Rooms interface {
	Add(room *Room)
	Find(name string) *Room
	Delete(room *Room)
	List() map[*Room]bool
}

func NewRooms() Rooms {
	return &roomsList{
		list: make(map[*Room]bool),
	}
}

type roomsList struct {
	list map[*Room]bool
}

func (r roomsList) Add(room *Room) {
	r.list[room] = true
}

func (r roomsList) Find(name string) *Room {
	for room := range r.list {
		if room.info.Name == name {
			return room
		}
	}
	return nil
}

func (r roomsList) Delete(room *Room) {
	delete(r.list, room)
}

func (r roomsList) List() map[*Room]bool {
	return r.list
}
