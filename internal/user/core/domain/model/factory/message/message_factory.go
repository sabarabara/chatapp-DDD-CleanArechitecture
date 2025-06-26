package messagefactory

import (
	users "chatapp_demo/internal/user/core/domain/model/vo/user"
	dto "chatapp_demo/internal/user/core/dto/message"
	db_entity "chatapp_demo/internal/user/core/entity"
)

type MessageFactory interface{
	CreateMessage(messagedto *dto.MessageDTO,oppuserid *users.Userid,oppusername *users.Username)(*db_entity.MessageEntity,error)
}