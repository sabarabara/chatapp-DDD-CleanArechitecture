package entity

import users "chatapp_demo/internal/user/core/domain/model/vo/user"

type UserSession struct{
	userid *users.Userid
	username *users.Username
	email *users.Email
}

func NewUserSession(userid *users.Userid,username *users.Username,email *users.Email)*UserSession{
	return &UserSession{userid: userid,
	                    username: username,
											email: email}
}

func (u *UserSession)GetUserId()(*users.Userid){
	return u.userid
}

func (u *UserSession)GetUsername()(*users.Username){
	return u.username
}

func (u *UserSession)GetEmail()(*users.Email){
	return u.email
}