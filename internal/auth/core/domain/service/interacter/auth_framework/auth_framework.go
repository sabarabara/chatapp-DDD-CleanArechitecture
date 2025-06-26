package authframework

import users "chatapp_demo/internal/user/core/domain/model/vo/user"

type AuthFramework interface{
	ComparePassword(password users.Password)(string,error)
}