package messagefactory

import (
	"chatapp_demo/internal/user/core/domain/model/vo/message"
	dto "chatapp_demo/internal/user/core/dto/message"
)

type MessageFactory interface{
	CreateMessage(message *dto.MessageDTO)(*message.Message,error)
}