package router

import "github.com/gin-gonic/gin"

func RegisterImageRouter(r *gin.Engine){
	imageGroup := r.Group("/image")
	{
		imageGroup.POST("/post".Imagectl)
	}
}