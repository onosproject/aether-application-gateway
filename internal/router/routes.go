package router

import (
	"github.com/gin-gonic/gin"
	"github.com/onosproject/aether-application-gateway/internal/controllers/v1"
)

func SetAllRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		SetDeviceRoutes(api)
		SetApplicationRoutes(api)
	}
}

func SetDeviceRoutes(api *gin.RouterGroup) {
	devices := api.Group("/devices")
	{
		devices.GET("", v1.ListDevices)
		devices.POST("", v1.CreateDevice)
		devices.GET("/:id", v1.GetDevice)
		devices.PUT("/:id", v1.UpdateDevice)
		devices.DELETE("/:id", v1.DeleteDevice)
	}
}

func SetApplicationRoutes(api *gin.RouterGroup) {
	applications := api.Group("/applications")
	{
		applications.GET("", v1.ListApplications)
		applications.POST("", v1.CreateApplication)
		applications.GET("/:id", v1.GetApplication)
		applications.PUT("/:id", v1.UpdateApplication)
		applications.DELETE("/:id", v1.DeleteApplication)
	}
}
