package dto


type UserSessionDTO struct{
	userid uint64
	username string
	email string
}

func NewUserSession(userid uint64,username string,email string)(*UserSessionDTO){
	return &UserSessionDTO{userid: userid,
	                       username: username,
												email:email}
}

func (u *UserSessionDTO)GetUserId()(uint64){
	return u.userid
}

func (u *UserSessionDTO)GetUsername()(string){
	return u.username
}
func (u *UserSessionDTO)GetEmail()(string){
	return u.email
}