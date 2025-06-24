package router

import "github.com/gin-gonic/gin"

func RegisterMessageRouter(r *gin.Engine){

	messageGroup := r.Group("/message")
	{
		messageGroup.POST("/post",PostMsgCtl)
		messageGroup.GET("/get-one",GetOneMsgCtl)
		messageGroup.GET("/get-all",GetAllMsgCtl)
		messageGroup.DELETE("/delete",DeleteMsgCtl)
	}
}