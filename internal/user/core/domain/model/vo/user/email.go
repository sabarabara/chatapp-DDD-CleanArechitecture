package users

type Email struct{
	email string
}

func NewEmail(email string)(*Email,error){
	return &Email{email: email},nil
}