package sessionstore

import (
	"chatapp_demo/internal/auth/core/domain/model/entity"
	"chatapp_demo/internal/auth/core/domain/service/interacter/sessionstore"
	users "chatapp_demo/internal/user/core/domain/model/vo/user"
	sessionconfig "chatapp_demo/pkg/config"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

type SessionstoreImpl struct{
	userid uint64
	username string
	email string
}

var _ sessionstore.SessionStore = (*SessionstoreImpl)(nil)

func (s *SessionstoreImpl)CreateSession(w http.ResponseWriter, r *http.Request,usersession entity.UserSession)(*entity.UserSession,error){

	//init
	session, _ := sessionconfig.Store.Get(r, "usersession")

	session.Values["userid"] = usersession.GetUserId()
	session.Values["username"] = usersession.GetUsername()
	session.Values["email"] = usersession.GetEmail()

	//sessionstoreへの値の代入
	userid := session.Values["userid"].(uint64)
	username := session.Values["username"].(string)
	email := session.Values["email"].(string)

	//インスタンス化
	validUserid,err := users.NewUserid(userid)
	if err != nil{
		fmt.Println(err)
	}

	validUsername,err := users.NewUsername(username)
	if err != nil{
		fmt.Println(err)
	}

	validEmail,err := users.NewEmail(email)
	if err != nil{
		fmt.Println(err)
	}

	//整形
	validUserSession := entity.NewUserSession(validUserid,validUsername,validEmail)


	//save & return
	return validUserSession,sessions.Save(r,w)
}

func (s *SessionstoreImpl)GetUserId()(*users.Userid){
	userid,err :=  users.NewUserid(s.userid)
	if err != nil{
		fmt.Println(err)
	}
	return userid
}

func (s *SessionstoreImpl)GetUsername()(*users.Username){
	username,err :=  users.NewUsername(s.username)
	if err != nil{
		fmt.Println(err)
	}
	return username
}

func (s *SessionstoreImpl)GetEmail()(*users.Email){
	email,err :=  users.NewEmail(s.email)
	if err != nil{
		fmt.Println(err)
	}
	return email
}