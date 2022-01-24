// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/onosproject/aether-application-gateway/internal/controllers"
)

// Setup initializes Gin engine
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
			devices.POST("/:id/connect", app.ConnectDevice)
			devices.POST("/:id/disconnect", app.DisconnectDevice)
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
