package interacter

import "chatapp_demo/internal/user/core/domain/model/vo/message"

type Geminier interface{
	CreateMessage(message message.Message)(message.Message,error)
}