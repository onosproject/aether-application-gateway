// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package services

import (
	"errors"
	"fmt"
	"github.com/onosproject/aether-application-gateway/internal/models"
	"github.com/onosproject/aether-application-gateway/internal/repository"
	"log"
)

// DeviceService -
type DeviceService struct {
	siteRepo      repository.SiteRepository
	analyticsRepo repository.AnalyticsRepository
}

// NewDeviceService creates a new NewDeviceService
func NewDeviceService(sr repository.SiteRepository, ar repository.AnalyticsRepository) DeviceService {
	return DeviceService{
		siteRepo:      sr,
		analyticsRepo: ar,
	}
}

// GetAllDevices call to get all devices
func (c *DeviceService) GetAllDevices(enterpriseID, siteID string) ([]models.Device, error) {
	var listDeviceResponse []models.Device

	site, err := c.siteRepo.GetSite(enterpriseID, siteID)
	if err != nil {
		log.Println("error getting site info")
	}

	simMap := mapSimToIccid(site)

	deviceGroupMap := mapDeviceToDeviceGroups(site)

	deviceMap := make(map[string]models.Device)
	for _, d := range site.Device {
		newDev := models.Device{
			ID:           d.DevID,
			Name:         d.DisplayName,
			Description:  d.Description,
			IMEI:         d.Imei,
			SimICCID:     simMap[d.SimCard],
			DeviceGroups: deviceGroupMap[d.DevID],
		}
		deviceMap[d.DevID] = newDev
		listDeviceResponse = append(listDeviceResponse, newDev)
	}

	return listDeviceResponse, nil
}

// GetDevice -> call to get a single device
//
// → (Throughput, Latency, Jitter, PacketDrop)
// Behind the scenes, does the mapping of DeviceID to IMSI, the mapping of IMSI to IPv4 address,
// and fetches fabric metrics from prometheus.
//
// → (State, Location, Priority, History)
// Behind the scenes, does the mapping of DeviceID to IMSI, queries SD-Core components prometheus such
// as SPGWC for subscriber-related data
func (c *DeviceService) GetDevice(enterpriseID, siteID, id string) (models.Device, error) {
	site, err := c.siteRepo.GetSite(enterpriseID, siteID)
	if err != nil {
		log.Println("error getting site info")
	}

	simMap := mapSimToIccid(site)

	deviceGroupMap := mapDeviceToDeviceGroups(site)

	var device models.Device
	for _, d := range site.Device {
		if d.DevID == id {

			device = models.Device{
				ID:           id,
				Name:         d.DisplayName,
				Description:  d.Description,
				IMEI:         d.Imei,
				SimICCID:     simMap[d.SimCard],
				DeviceGroups: deviceGroupMap[id],
			}

			return device, nil
		}
	}

	return models.Device{}, errors.New("device not found")
}

func mapSimToIccid(site repository.Site) map[string]int {
	simMap := make(map[string]int)
	for _, s := range site.SimCard {
		simMap[s.SimID] = s.Iccid
	}
	return simMap
}

// TODO: handle multiple device groups for a device
func mapDeviceToDeviceGroups(site repository.Site) map[string]string {
	deviceGroupMap := make(map[string]string)
	for _, dg := range site.DeviceGroup {
		if len(dg.Device) > 0 {
			for _, d := range dg.Device {
				deviceID := d.(map[string]interface{})["device-id"]
				deviceIDStr := fmt.Sprintf("%v", deviceID)
				deviceGroupMap[deviceIDStr] = dg.DgID
			}
		}
	}
	return deviceGroupMap
}
