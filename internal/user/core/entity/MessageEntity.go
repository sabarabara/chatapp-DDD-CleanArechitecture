package db_entity

import (
	"chatapp_demo/internal/user/core/domain/model/vo/message"
	users "chatapp_demo/internal/user/core/domain/model/vo/user"
)

type MessageEntity struct{
	messageId string
	message *message.Message
	username *users.Username
	userId *users.Userid
	oppuserId *users.Userid
}

func NewMessageEntity (messageId string,message *message.Message,username *users.Username,userid *users.Userid,oppuserid *users.Userid)(*MessageEntity,error){
	return &MessageEntity{messageId: messageId,
		                    message :message,
	                      username: username,
										    userId:userid,
											  oppuserId:oppuserid},nil
}

func (m *MessageEntity)getMessageEntity()(*MessageEntity){
	return m;
}