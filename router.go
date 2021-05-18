package main

import (
	"data-view/config"
	"data-view/handler"
	"data-view/middleware"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"))

type Router struct {
	DataSource *handler.DataSourceHandler
	Image      *handler.ImageHandler
	DataView   *handler.DataViewHandler
}

func (r *Router) NewRouter() (router *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()

	if config.Instance.App.AllowCrossDomain {
		router.Use(middleware.Cors())
	}
	router.Use(middleware.ErrHandler())

	v1 := router.Group("/api/v1/")
	{
		// 数据源管理
		dataSource := v1.Group("dataSource/").Use(middleware.CheckUser())
		{
			dataSource.GET("/", r.DataSource.GetPage)
			dataSource.GET("/list", r.DataSource.GetList)
		}

		// 图片管理
		image := v1.Group("image/").Use(middleware.CheckUser())
		{
			image.GET("/", r.Image.GetPage)
			image.GET("/list", r.Image.GetList)
		}

		// 数据可视化大屏管理
		dataView := v1.Group("dataView/").Use(middleware.CheckUser())
		{
			dataView.GET("/", r.DataView.GetPage)
			dataView.GET("/:id", r.DataView.Get)
			dataView.GET("/chartData", r.DataView.GetChartData)
		}
	}

	return
}
