package generater

import users "chatapp_demo/internal/user/core/domain/model/vo/user"

type Generator interface {
	Next(sequenceName string) (*users.Userid, error)
}
