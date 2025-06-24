package repository

import (
	users "chatapp_demo/internal/user/core/domain/model/vo/user"
	db_entity "chatapp_demo/internal/user/core/entity"
)

type UserRepository interface{
	FindByUser(email *users.Email)()
	FindByUsers()()
	CreateUser(userentity *db_entity.UserEntity)(*db_entity.UserEntity,error)
	DeleteUser()()
}