package dto

type MessageDTO struct{
	message string
}

func NewMessageDTO(message string)(*MessageDTO,error){
	return &MessageDTO{message: message},nil
}

func (m *MessageDTO)GetMessageDTO()(string){
	return m.message
}