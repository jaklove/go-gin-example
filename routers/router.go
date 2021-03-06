package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jacklove/go-gin-example/docs"
	"github.com/jacklove/go-gin-example/middleware/jwt"
	"github.com/jacklove/go-gin-example/pkg/export"
	"github.com/jacklove/go-gin-example/pkg/qrcode"
	"github.com/jacklove/go-gin-example/pkg/setting"
	"github.com/jacklove/go-gin-example/pkg/upload"
	"github.com/jacklove/go-gin-example/routers/api"
	v1 "github.com/jacklove/go-gin-example/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.AppSetting.RUN_MODE)

	r.GET("/auth",api.GetAuth)
	r.POST("/upload", api.UploadImage)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// 获取标签列表
		apiv1.GET("/tags",v1.GetTags)

		//新建标签
		apiv1.POST("/tags",v1.AddTag)

		//更新指定标签
		apiv1.PUT("/tags/:id",v1.EditTag)

		//删除指定标签
		apiv1.DELETE("/tags/:id",v1.DeleteTag)

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
		//生成文章海报
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	//导出标签
	r.POST("/tags/export", v1.ExportTag)

	return r
}
