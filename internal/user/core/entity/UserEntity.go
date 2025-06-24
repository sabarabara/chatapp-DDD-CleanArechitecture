package db_entity

import (
	"chatapp_demo/internal/user/core/domain/model/entity"
	users "chatapp_demo/internal/user/core/domain/model/vo/user"
	"fmt"
)

type UserEntity struct{
	userId *users.Userid
	username *users.Username
	password *users.Password
	email *users.Email
}

func NewUserEntity(userid *users.Userid,user *entity.User)(*UserEntity,error){

	username,err :=user.GetUsername()
	if err != nil{
		fmt.Println(err)
	}

	password,err :=user.GetPassword()
	if err != nil{
		fmt.Println(err)
	}
	
	email,err:= user.GetEmail()
	if err != nil{
		fmt.Println(err)
	}


	return &UserEntity{
		userId: userid,
		username: username,
		password: password,
		email:email},nil
}


func (u *UserEntity)GetUserEntity()(*UserEntity){
	return u;
}