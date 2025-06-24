package factoryimpl

import (
	messagefactory "chatapp_demo/internal/user/core/domain/model/factory/message"
	"chatapp_demo/internal/user/core/domain/model/vo/message"
	dto "chatapp_demo/internal/user/core/dto/message"
	db_entity "chatapp_demo/internal/user/core/entity"
	"fmt"
)

var _ messagefactory.MessageFactory =(*MessageFactory)(nil)

type MessageFactory struct{}

func (m *MessageFactory)CreateMessage(messagedto *dto.MessageDTO)(*message.Message,error){
	msg := messagedto.GetMessageDTO()
	validmsg,err := message.NewMessage(msg)
	if err != nil{
		fmt.Println(err)
	}

	messageentity,err := db_entity.NewMessageEntity(validmsg)
	
}