package repository

import users "chatapp_demo/internal/user/core/domain/model/vo/user"

type MessageRepository interface{
	FindAllMessage(user *users.Userid,oppuser *users.Userid)()
	CreateMessage()()
	DeleteMessage()()
}