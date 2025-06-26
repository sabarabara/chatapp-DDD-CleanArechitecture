package factoryimpl

import (
	"chatapp_demo/internal/auth/core/domain/service/interacter/sessionstore"
	messagefactory "chatapp_demo/internal/user/core/domain/model/factory/message"
	"chatapp_demo/internal/user/core/domain/model/vo/message"
	users "chatapp_demo/internal/user/core/domain/model/vo/user"
	"chatapp_demo/internal/user/core/domain/service/generater"
	dto "chatapp_demo/internal/user/core/dto/message"
	db_entity "chatapp_demo/internal/user/core/entity"
	"fmt"
	"time"
)

var _ messagefactory.MessageFactory =(*MessageFactory)(nil)

type MessageFactory struct{
	IDGen generater.Generator
	SessionStore sessionstore.SessionStore
}

func (m *MessageFactory)CreateMessage(messagedto *dto.MessageDTO,oppuserid *users.Userid,oppusername *users.Username)(*db_entity.MessageEntity,error){
	msg := messagedto.GetMessageDTO()
	validmsg,err := message.NewMessage(msg)
	if err != nil{
		fmt.Println(err)
	}

	//generaterによる特有のid作成
	messageid,err := m.IDGen.Next("user-seq")
	if err != nil{
		fmt.Println(err)
	}

	//dateの作成
	date := time.Now()

	//sessionから値の取得
	userid := m.SessionStore.GetUserId()
	username := m.SessionStore.GetUsername()

	//相手のuserid作成
	oppid := oppuserid
	
	//entityの作成
	messageentity,err := db_entity.NewMessageEntity(messageid,validmsg,username,userid,oppid,date)
	
	return messageentity,nil
}