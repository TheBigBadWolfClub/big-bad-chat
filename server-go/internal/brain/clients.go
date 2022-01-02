package brain

type Clients interface {
	Add(conn *Client)
	FindByName(name string) *Client
	FindByUuid(uuid string) *Client
	Find(client Client) *Client
	Delete(conn *Client)
	List() map[*Client]bool
}

func NewClients() Clients {
	return &clientList{
		list: make(map[*Client]bool),
	}
}

type clientList struct {
	list map[*Client]bool
}

func (r clientList) Add(conn *Client) {
	r.list[conn] = true
}

func (r clientList) Delete(conn *Client) {
	delete(r.list, conn)
}

func (r clientList) List() map[*Client]bool {
	return r.list
}

func (r clientList) FindByName(name string) *Client {
	for conn := range r.list {
		if conn.info.Name == name {
			return conn
		}
	}
	return nil
}

func (r clientList) FindByUuid(uuid string) *Client {
	for conn := range r.list {
		if conn.info.UUID == uuid {
			return conn
		}
	}
	return nil
}

func (r clientList) Find(client Client) *Client {
	return r.FindByUuid(client.info.UUID)

}
