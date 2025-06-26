package authframework

import (
	authframework "chatapp_demo/internal/auth/framework/authFramework"
	users "chatapp_demo/internal/user/core/domain/model/vo/user"
)

var _ authframework.authFramework = (*AuthFramework)(nil)

type AuthFramework struct{}

func (a *AuthFramework)ComparePassword(password users.Password)(string,error){

	
}