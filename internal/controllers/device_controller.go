// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/onosproject/aether-application-gateway/internal/models"
	"github.com/onosproject/aether-application-gateway/internal/services"
	"net/http"
)

// DeviceController -
type DeviceController struct {
	deviceService services.DeviceService
}

// NewDeviceController creates a new DeviceController
func NewDeviceController(deviceService services.DeviceService) DeviceController {
	return DeviceController{
		deviceService: deviceService,
	}
}

// ListDevices returns a list of Devices for a site
func (dc *DeviceController) ListDevices(c *gin.Context) {
	ent := c.Param("enterprise")
	site := c.Param("site")

	devices, err := dc.deviceService.GetAllDevices(ent, site)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed finding devices"})
		return
	}
	c.JSON(http.StatusOK, devices)
}

// GetDevice returns a single device
func (dc *DeviceController) GetDevice(c *gin.Context) {
	ent := c.Param("enterprise")
	site := c.Param("site")
	id := c.Param("id")

	device, err := dc.deviceService.GetDevice(ent, site, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed finding device"})
		return
	}
	c.JSON(http.StatusOK, device)
}

// CreateDevice -
func (dc *DeviceController) CreateDevice(c *gin.Context) {
	var newDevice models.Device

	if err := c.BindJSON(&newDevice); err != nil {
		return
	}

	//models.DevicesStub = append(models.DevicesStub, newDevice)
	c.JSON(http.StatusCreated, newDevice)
}

// UpdateDevice -
func (dc *DeviceController) UpdateDevice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteDevice -
func (dc *DeviceController) DeleteDevice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ConnectDevice Connects a device to a slice.
// Application developer can choose to put it into a specific named DeviceGroup,
// or can name the slice, in which case the device will be added to an available group within that slice.
func (dc *DeviceController) ConnectDevice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DisconnectDevice -
func (dc *DeviceController) DisconnectDevice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
