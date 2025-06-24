package dto

type ProfDTO struct{
	username string
	s3url string
}

func NewProfDTO(username string,s3url string)(*ProfDTO,error){
	return &ProfDTO{username:username,
	                s3url: s3url},nil
}


func (p*ProfDTO)getUsername()(string){
	return p.username
}

func (p*ProfDTO)getS3URL()(string){
	return p.s3url
}


func (p *ProfDTO)GetProfDTO()(*ProfDTO){
	return p;
}