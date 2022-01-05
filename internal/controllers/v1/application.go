package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/onosproject/aether-application-gateway/internal/models"
	"net/http"
)

func CreateApplication(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
	var newApplication models.Application

	if err := c.BindJSON(&newApplication); err != nil {
		return
	}

	models.Applications = append(models.Applications, newApplication)
	c.IndentedJSON(http.StatusCreated, newApplication)
}

func GetApplication(c *gin.Context) {
	id := c.Param("id")

	for _, a := range models.Applications {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "application not found"})
}

func ListApplications(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Applications)
}

func UpdateApplication(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func DeleteApplication(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
