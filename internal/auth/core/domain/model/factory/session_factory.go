package factory

import "chatapp_demo/internal/auth/core/domain/model/entity"

type SessionFactory interface{
	CreateSession(usersessiondto entity.UserSession)()
}