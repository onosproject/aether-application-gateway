package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/onosproject/aether-application-gateway/internal/models"
	"net/http"
)

func CreateDevice(c *gin.Context) {
	var newDevice models.Device

	if err := c.BindJSON(&newDevice); err != nil {
		return
	}

	models.Devices = append(models.Devices, newDevice)
	c.IndentedJSON(http.StatusCreated, newDevice)
}

func ListDevices(c *gin.Context) {
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
func GetDevice(c *gin.Context) {
	id := c.Param("id")

	for _, a := range models.Devices {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "device not found"})
}

// UpdateDevice handles connecting and disconnecting
// a Device from a Slice
func UpdateDevice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func DeleteDevice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
