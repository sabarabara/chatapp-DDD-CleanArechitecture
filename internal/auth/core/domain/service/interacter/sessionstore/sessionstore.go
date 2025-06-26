package sessionstore

import (
	"chatapp_demo/internal/auth/core/domain/model/entity"
	users "chatapp_demo/internal/user/core/domain/model/vo/user"
	"net/http"
)


type SessionStore interface{
	CreateSession(w http.ResponseWriter, r *http.Request,usersessiondto entity.UserSession)(*entity.UserSession, error)
	GetUserId()(*users.Userid)
	GetUsername()(*users.Username)
	GetEmail()(*users.Email)
}