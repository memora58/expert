package routers

import (
	"gateway/controller"
	"gateway/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(service []interface{}) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors(), middleware.InitMiddleware(service), middleware.ErrorMiddleware())
	//store := cookie.NewStore([]byte("something-very-secret"))
	//ginRouter.Use(sessions.Sessions("mysession", store))

	rGroup := router.Group("api")
	{
		rGroup.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		// user service
		rGroup.POST("user/register", (&controller.UserController{}).Register)
		rGroup.POST("user/login", (&controller.UserController{}).Login)

		// jwt 教研
		rGroup.Use(middleware.JWT())
		{
			// task service
			rGroup.GET("task", (&controller.TaskController{}).GetTaskList)
			rGroup.POST("task", (&controller.TaskController{}).CreateTask)
			//rGroup.PUT("task")
			//rGroup.DELETE("task")
		}
	}
	return router
}
