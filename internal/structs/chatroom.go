package structs

type Chatroom struct {
	Messages []string
	Users    []*User
}

func (ch *Chatroom) PushMessage(message string) {
	ch.Messages = append(ch.Messages, message)
}
