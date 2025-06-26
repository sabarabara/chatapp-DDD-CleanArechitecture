package repository

import (
	users "chatapp_demo/internal/user/core/domain/model/vo/user"
	"chatapp_demo/internal/user/core/domain/service/interacter/repository"
	db_entity "chatapp_demo/internal/user/core/entity"

	"gorm.io/gorm"
)

var _ repository.UserRepository =(*UserRepositoryImpl)(nil)

type UserRepositoryImpl struct{
	db *gorm.DB
}

func (u *UserRepositoryImpl)FindByUser(email *users.Email)(*db_entity.UserEntity,error){
	var user *db_entity.UserEntity

	result := u.db.Where("email = ?",email).
	               First(&user)

	if result.Error != nil{
		return nil,result.Error
	}

	return user,nil
}

func (u*UserRepositoryImpl)CreateUser(userentity *db_entity.UserEntity)(*db_entity.UserEntity,error){
	result := u.db.Save(userentity)

	if result.Error != nil {
		return nil,result.Error
	}

	return userentity,nil
}

func (u*UserRepositoryImpl)DeleteUser(userid *users.Userid)(string,error){
	result := u.db.Where("userid = ?",userid).
	               Delete(&db_entity.UserEntity{})
		
		if result.Error != nil {
			return "NO",result.Error
		}
	return "Ok",nil
}

func (u*UserRepositoryImpl)UpdateUser(userentity *db_entity.UserEntity)(*db_entity.UserEntity,error){

	result := u.db.Save(userentity)
	if result.Error != nil{
		return nil,result.Error
	}

	return userentity,nil
}