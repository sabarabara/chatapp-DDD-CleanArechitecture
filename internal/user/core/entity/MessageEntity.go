package db_entity

import (
	"chatapp_demo/internal/user/core/domain/model/vo/message"
	users "chatapp_demo/internal/user/core/domain/model/vo/user"
	"time"
)

type MessageEntity struct{
	messageId *users.Userid
	message *message.Message
	username *users.Username
	userId *users.Userid
	oppuserId *users.Userid
	date time.Time
}

func NewMessageEntity (messageId *users.Userid,message *message.Message,username *users.Username,userid *users.Userid,oppuserid *users.Userid,date time.Time)(*MessageEntity,error){
	return &MessageEntity{messageId: messageId,
		                    message :message,
	                      username: username,
										    userId:userid,
											  oppuserId:oppuserid,
												date:date},nil
}

func (m *MessageEntity)GetMessageEntity()(*MessageEntity){
	return m;
}