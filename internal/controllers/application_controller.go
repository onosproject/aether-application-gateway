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

// App -
type App struct {
}

// ApplicationController -
type ApplicationController struct {
}

// NewApplicationController creates a new ApplicationController
func NewApplicationController() ApplicationController {
	return ApplicationController{}
}

// CreateApplication -
func (a *App) CreateApplication(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
	var newApplication models.Application

	if err := c.BindJSON(&newApplication); err != nil {
		return
	}

	models.Applications = append(models.Applications, newApplication)
	c.IndentedJSON(http.StatusCreated, newApplication)
}

// GetApplication -
func (a *App) GetApplication(c *gin.Context) {
	id := c.Param("id")

	for _, a := range models.Applications {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "application not found"})
}

// ListApplications -
func (a *App) ListApplications(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Applications)
}

// UpdateApplication -
func (a *App) UpdateApplication(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteApplication -
func (a *App) DeleteApplication(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
