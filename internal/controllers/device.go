// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/onosproject/aether-application-gateway/internal/models"
	"net/http"
)

// CreateDevice -
func (a *App) CreateDevice(c *gin.Context) {
	var newDevice models.Device

	if err := c.BindJSON(&newDevice); err != nil {
		return
	}

	models.Devices = append(models.Devices, newDevice)
	c.IndentedJSON(http.StatusCreated, newDevice)
}

// ListDevices -
func (a *App) ListDevices(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Devices)
}

// GetDevice should include Telemetry data:
//
// GetDeviceConnectivity(DeviceID) → (Throughput, Latency, Jitter, PacketDrop)
// Behind the scenes, does the mapping of DeviceID to IMSI, the mapping of IMSI to IPv4 address,
// and fetches fabric metrics from prometheus.
//
// GetDeviceStatus(DeviceID) → (State, Location, Priority, History)
// Behind the scenes, does the mapping of DeviceID to IMSI, queries SD-Core components prometheus such
// as SPGWC for subscriber-related data
func (a *App) GetDevice(c *gin.Context) {
	id := c.Param("id")

	for _, a := range models.Devices {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "device not found"})
}

// UpdateDevice -
func (a *App) UpdateDevice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteDevice -
func (a *App) DeleteDevice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ConnectDevice Connects a device to a slice.
// Application developer can choose to put it into a specific named DeviceGroup,
// or can name the slice, in which case the device will be added to an available group within that slice.
func (a *App) ConnectDevice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DisconnectDevice -
func (a *App) DisconnectDevice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
