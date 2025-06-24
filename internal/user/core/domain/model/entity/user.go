package entity

import (
	users "chatapp_demo/internal/user/core/domain/model/vo/user"
	"fmt"
)

type User struct{
	username *users.Username
	password *users.Password
	email *users.Email
}

func NewUser(name string,pw string,eml string)(*User,error){

	username,err:=users.NewUsername(name)
	if err != nil{
		fmt.Println(err)
	}
	password,err:=users.NewPassword(pw)
	if err != nil{
		fmt.Println(err)
	}

	email,err :=users.NewEmail(eml)
	if err != nil {
		fmt.Println(err)
	}

	return &User{
		username: username,
		password: password,
		email: email,
	},nil
}

func (u*User)GetUsername()(*users.Username,error){
	return u.username,nil
}

func (u*User)GetPassword()(*users.Password,error){
	return u.password,nil
}

func (u*User)GetEmail()(*users.Email,error){
	return u.email,nil
}