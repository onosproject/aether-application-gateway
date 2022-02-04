// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/onosproject/aether-application-gateway/internal/config"
	"github.com/onosproject/aether-application-gateway/internal/controllers"
	"github.com/onosproject/aether-application-gateway/internal/repository"
	"github.com/onosproject/aether-application-gateway/internal/services"
	promApiV1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"net/http"
)

// Setup initializes Gin engine
func Setup(hc *http.Client, roc *config.RocConfig, promV1 promApiV1.API) *gin.Engine {
	r := gin.Default()

	sr := repository.NewSiteRepository(hc, roc)
	ar := repository.NewAnalyticsRepository(promV1)
	ds := services.NewDeviceService(sr, ar)
	dc := controllers.NewDeviceController(ds)

	setDeviceRoutes(r, dc)

	return r
}

func setDeviceRoutes(r *gin.Engine, dc controllers.DeviceController) {
	v1 := r.Group("/v1/:enterprise/:site/devices")
	{
		v1.GET("", dc.ListDevices)
		v1.GET("/:id", dc.GetDevice)
		v1.POST("", dc.CreateDevice)
		v1.PUT("/:id", dc.UpdateDevice)
		v1.DELETE("/:id", dc.DeleteDevice)
		v1.POST("/:id/connect", dc.ConnectDevice)
		v1.POST("/:id/disconnect", dc.DisconnectDevice)
	}
}

//func setApplicationRoutes(r *gin.Engine, app *controllers.App) {
//	apps := r.Group("/applications")
//	{
//		{
//			apps.GET("", app.ListApplications)
//			apps.POST("", app.CreateApplication)
//			apps.GET("/:id", app.GetApplication)
//			apps.PUT("/:id", app.UpdateApplication)
//			apps.DELETE("/:id", app.DeleteApplication)
//		}
//	}
//}
