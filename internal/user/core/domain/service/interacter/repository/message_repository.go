package repository

import (
	users "chatapp_demo/internal/user/core/domain/model/vo/user"
	db_entity "chatapp_demo/internal/user/core/entity"
)

type MessageRepository interface{
	CreateMessage(messageentity *db_entity.MessageEntity)(*db_entity.MessageEntity,error)
	FindMessageByUserIdandOppuserid(userid *users.Userid,oppuserid *users.Userid)([]*db_entity.MessageEntity,error)
	deleteMessage(messageid *users.Userid)(string,error)
	UpdateMessage(messageentity *db_entity.MessageEntity)(*db_entity.MessageEntity,error)
}