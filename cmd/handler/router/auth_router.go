package router

import "github.com/gin-gonic/gin"

func RegisterAuthRouter(r *gin.Engine){

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login",LoginCtl)
		authGroup.GET("/logout",LogoutCtl)
	}
}