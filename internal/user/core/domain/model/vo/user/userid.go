package users

type Userid struct{
	userid uint64
}

func NewUserid(userid uint64)(*Userid,error){
	return &Userid{userid: userid},nil
}

func (u *Userid)GetUserId()(uint64,error){
	return u.userid,nil
}