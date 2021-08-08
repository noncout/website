package routers

import (
	"github.com/EDDYCJY/go-gin-example/middleware/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"website/pkg/setting"
	"website/pkg/upload"
	"website/routers/api"
	v1 "website/routers/api/v1"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	gin.SetMode(setting.ServerSetting.RunMode)

	router.GET("/auth", api.GetAuth)

	apiv1 := router.Group("/api/v1")

	router.POST("/upload", api.UploadImage)

	router.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	apiv1.Use(jwt.JWT())
	{
		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}
	return router
}
