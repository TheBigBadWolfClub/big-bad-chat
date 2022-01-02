package brain

import "time"

type Messages interface {
	Append(msg ChatMessage)
	Get()
}

type ChatMessage struct {
	UUID     string    `json:"uuid"`
	Time     time.Time `json:"time"`
	Text     []string  `json:"text"`
	Sender   ID        `json:"sender"`
	RoomUuid string    `json:"room_uuid"`
}

func NewMessages() Messages {
	return &store{list: make([]ChatMessage, 0)}
}

type store struct {
	list []ChatMessage
}

func (s store) Append(msg ChatMessage) {
	last := s.list[len(s.list)-1]
	// if same userID and less 20seconds, append text to previous msg
	if sameSender(msg.Sender.UUID, last.Sender.UUID) && less20seconds(msg.Time, last.Time) {
		last.Text = append(last.Text, msg.Text...)
	}
	s.list = append(s.list, msg)
}

func (s store) Get() {
	panic("implement me")
}

func (s store) last() (last *ChatMessage, ok bool) {
	if len(s.list) == 0 {
		return nil, false
	}
	return &s.list[len(s.list)-1], true
}

func (s store) fist() (first *ChatMessage, ok bool) {
	if len(s.list) == 0 {
		return nil, false
	}
	return &s.list[0], true

}

// if same user msg in lower interval, append text
func sameSender(newSender string, prevSender string) bool {
	return newSender == prevSender
}

func less20seconds(newSender time.Time, prevSender time.Time) bool {
	return false
}
