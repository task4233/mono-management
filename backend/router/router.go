/*
Package router は, Routesの設定をするパッケージです。

docker-composeの環境下では
	import app/router
でimportできます。
*/
package router

import (
	"github.com/gin-gonic/gin"

	// internal
	"app/internal"
)

/*
Create funcは, APIのエンドポイント(Routes)を設定するメソッドです。

	internal.Create()
で使用できます。
*/
func Create() *gin.Engine {
	router := gin.Default()
	// Routes
	// grouping
	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			// mono
			mono := v1.Group("mono")
			{
				// CRUD
				mono.GET("/", internal.GetMonos)
				mono.POST("/new", internal.CreateMono)
				mono.PUT("/:monoId", internal.UpdateMonos)
				mono.DELETE("/:monoId", internal.DeleteMono)
			}
			// user
			user := v1.Group("user")
			{
				user.GET("/info", internal.GetInfo)
				user.POST("/login", internal.Login)
				user.POST("/logout", internal.Logout)
				user.POST("new", internal.CreateAccount)
				user.POST("update", internal.UpdateAccount)
				user.DELETE("delete", internal.DeleteAccount)
			}
			// tag
			tag := v1.Group("tag")
			{
				tag.GET("/", internal.GetTags)
				tag.POST("/new", internal.CreateTag)
				tag.PUT("/:tagId", internal.UpdateTag)
				tag.DELETE("/:tagId", internal.DeleteTag)
			}
			// search
			search := v1.Group("search")
			{
				search.POST("/", internal.SearchMonos)
			}
		}
	}
	return router
}
