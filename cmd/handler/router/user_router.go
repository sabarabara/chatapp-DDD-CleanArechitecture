package router

import "github.com/gin-gonic/gin"

func RegisterUserRouts(r *gin.Engine)(error){

	userGroup := r.Group("/user")
	{
		userGroup.POST("/post",CreateUsercrl)
		userGroup.GET("/get",GetUserCtl)
		userGroup.PUT("/put",PutUserCtl)
		userGroup.DELETE("/delete",DeleteUserctl)
	}

	return nil
}