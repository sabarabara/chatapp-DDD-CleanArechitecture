package dto

type UserDTO struct{
	username string
	password string
	email string
}

func NewUserDTO (username string,password string,email string)(*UserDTO,error){
	return &UserDTO{username: username,
	                password: password,
									email: email},nil
}

func (u *UserDTO)GetUsername()(string){
	return u.username
}

func (u *UserDTO)GetPassword()(string){
	return u.password
}

func (u *UserDTO)GetEmail()(string){
	return u.email
}