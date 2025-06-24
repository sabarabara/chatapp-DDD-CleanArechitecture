package router

import "github.com/gin-gonic/gin"

func RegisterProfRouter(r *gin.Engine){
	profGroup := r.Group("prof")
	{
		profGroup.GET("get",ProfGetCtl)
	}
}