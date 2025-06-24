package userfactory

import (
	dto "chatapp_demo/internal/user/core/dto/user"
	db_entity "chatapp_demo/internal/user/core/entity"
)

type UserFactory interface{
	CreateUser(userdto *dto.UserDTO)(*db_entity.UserEntity,error)
}