package repository

import (
	"chatapp_demo/internal/user/core/domain/model/vo/message"
	users "chatapp_demo/internal/user/core/domain/model/vo/user"
	db_entity "chatapp_demo/internal/user/core/entity"

	"gorm.io/gorm"
)

type MessageRepository interface{
	CreateMessage(messageentity *db_entity.MessageEntity)(*db_entity.MessageEntity,error)
	FindMessageByUserIdandOppuserid(userid *users.Userid,oppuserid *users.Userid)([]*db_entity.MessageEntity,error)
	DeleteMessage(messageid *users.Userid)(string,error)
	UpdateMessage(messageid users.Userid,message message.Message)(*gorm.DB,error)
}