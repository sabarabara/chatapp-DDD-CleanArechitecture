package message

type Message struct{
	message string
}

func NewMessage(message string)(*Message,error){
	return &Message{message: message},nil
}

func (m *Message)GetMessage()(string){
	return m.message
}