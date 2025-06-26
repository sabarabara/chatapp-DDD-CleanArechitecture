package repository

import (
	"chatapp_demo/internal/user/core/domain/model/vo/message"
	users "chatapp_demo/internal/user/core/domain/model/vo/user"
	"chatapp_demo/internal/user/core/domain/service/interacter/repository"
	db_entity "chatapp_demo/internal/user/core/entity"
	"gorm.io/gorm"
)

var _ repository.MessageRepository = (*MessageRepositoryImpl)(nil)

type MessageRepositoryImpl struct{
	db *gorm.DB
}


func (m *MessageRepositoryImpl)CreateMessage(messageentity *db_entity.MessageEntity)(*db_entity.MessageEntity,error){
	result := m.db.Save(messageentity)

	if result.Error != nil{
		return nil,result.Error
	}

	return messageentity,nil
}

func (m *MessageRepositoryImpl)FindMessageByUserIdandOppuserid(userid *users.Userid,oppuserid *users.Userid)([]*db_entity.MessageEntity,error){

	//大体どのくらいなんだろ30は重いか?よくわからん
	const getlim = 30


	resmessage := make([]*db_entity.MessageEntity,getlim,getlim)
	err := m.db.Where("(userid = ? AND oppuserid = ?) OR (userid = ? AND oppuserid = ?)",
										userid,oppuserid,oppuserid,userid).
							Order("date ASC").
							Limit(getlim).
							Find(&resmessage).Error

		if err != nil{
			return nil,err
		}
	return resmessage,nil
}

func (m *MessageRepositoryImpl)DeleteMessage(messageid *users.Userid)(string,error){
	result := m.db.Where("messageid = ?",messageid).
	          Delete(&db_entity.MessageEntity{})

		if result.Error != nil{
			return "No",result.Error
		}
		return "OK",nil
}

func (m *MessageRepositoryImpl)UpdateMessage(messageid users.Userid,message message.Message)(*gorm.DB,error){
	result := m.db.Model(&db_entity.MessageEntity{}).
	               Where("messageid = ?",messageid).
	               Update("message",message)

			if result.Error != nil{
				return nil,result.Error
			}

		return result,nil
}