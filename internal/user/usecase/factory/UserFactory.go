package factoryimpl

import (
	"chatapp_demo/internal/user/core/domain/model/entity"
	userfactory "chatapp_demo/internal/user/core/domain/model/factory/user"
	"chatapp_demo/internal/user/core/domain/service/generater"
	"chatapp_demo/internal/user/core/domain/service/interacter/repository"
	dto "chatapp_demo/internal/user/core/dto/user"
	"fmt"

	db_entity "chatapp_demo/internal/user/core/entity"
)


var _ userfactory.UserFactory =(*UserFactoryImpl)(nil)

type UserFactoryImpl struct{
	IDGen generater.Generator
	userRepo repository.UserRepository
}

func (f *UserFactoryImpl)CreateUser(userdto *dto.UserDTO)(*db_entity.UserEntity,error){

	//dtoからの値取得
	username := userdto.GetUsername()
	password := userdto.GetPassword()
	email := userdto.GetEmail()


	//voにおけるバリデーション
	user,err := entity.NewUser(username,password,email)
	if err != nil{
		fmt.Println(err)
	}

	//generaterによる特有のid作成
	userid,err := f.IDGen.Next("user-seq")
	if err != nil{
		fmt.Println(err)
	}

	//dbのentity作成
	userentity,err := db_entity.NewUserEntity(userid,user)
	if err != nil{
		fmt.Println(err)
	}

	return userentity,nil
}