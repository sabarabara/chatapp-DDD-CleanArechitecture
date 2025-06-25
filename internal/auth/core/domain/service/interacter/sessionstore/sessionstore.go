package sessionstore

import (
	"chatapp_demo/internal/auth/core/domain/model/entity"
	"net/http"
)


type SessionStore interface{
	CreateSession(w http.ResponseWriter, r *http.Request,usersessiondto entity.UserSession)(*entity.UserSession, error)
	GetUserId()(uint64)
	GetUsername()(string)
	GetEmail()(string)
}