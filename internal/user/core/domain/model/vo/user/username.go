package users

type Username struct{
	username string
}

func NewUsername(username string)(* Username,error){
	return &Username{username: username},nil
}