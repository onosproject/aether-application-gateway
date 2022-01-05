package models

// DeviceGroup defines model for DeviceGroup.
type DeviceGroup struct {
	DeviceGroupId string    `json:"device-group-id" yaml:"device-group-id"`
	Devices       *[]string `json:"devices,omitempty"`
	DisplayName   string    `json:"display-name" yaml:"display-name"`
}
