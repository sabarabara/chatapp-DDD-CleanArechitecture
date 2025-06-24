package users

type Password struct{
	password string
}

func NewPassword(password string)(*Password,error){
	return &Password{password: password},nil
}