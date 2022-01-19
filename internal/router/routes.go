package router

import (
	"github.com/gin-gonic/gin"
	"github.com/onosproject/aether-application-gateway/internal/controllers"
)

func Setup() *gin.Engine {
	r := gin.Default()
	app := &controllers.App{}

	setDeviceRoutes(r, app)
	setApplicationRoutes(r, app)

	return r
}

func setDeviceRoutes(r *gin.Engine, app *controllers.App) {
	devices := r.Group("/devices")
	{
		{
			devices.GET("", app.ListDevices)
			devices.POST("", app.CreateDevice)
			devices.GET("/:id", app.GetDevice)
			devices.PUT("/:id", app.UpdateDevice)
			devices.DELETE("/:id", app.DeleteDevice)
		}
	}
}

func setApplicationRoutes(r *gin.Engine, app *controllers.App) {
	apps := r.Group("/applications")
	{
		{
			apps.GET("", app.ListApplications)
			apps.POST("", app.CreateApplication)
			apps.GET("/:id", app.GetApplication)
			apps.PUT("/:id", app.UpdateApplication)
			apps.DELETE("/:id", app.DeleteApplication)
		}
	}
}
