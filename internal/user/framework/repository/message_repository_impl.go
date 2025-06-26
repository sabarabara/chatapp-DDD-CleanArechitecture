package repository

import (
	"chatapp_demo/internal/user/core/domain/service/interacter/repository"
	db_entity "chatapp_demo/internal/user/core/entity"
)

var _ repository.MessageRepository = (*MessageRepositoryImpl)(nil)

type MessageRepositoryImpl struct{}

func (m *MessageRepositoryImpl)CreateMessage(db_entity.MessageEntity)(string,error){

}

func (m *MessageRepositoryImpl)FindByMessage()(){
	
}