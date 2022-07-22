package routes

import (
	"gokit/controllers"
	"gokit/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouters() *gin.Engine {

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(gin.Recovery())

	r.LoadHTMLGlob("public/*/*")

	r.GET("/", controllers.Default)

	//r.GET("/:id",controllers.ReadOne)

	r.GET("/token", controllers.GetToken)

	group := r.Group("/api").Use(middleware.CheckToken())
	{
		group.POST("/", controllers.Create)

		group.PUT("/", controllers.Update)

		group.DELETE("/:id", controllers.Delete)
	}

	return r

}
