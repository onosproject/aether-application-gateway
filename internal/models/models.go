// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package models

// Application defines model for Application.
type Application struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// DeviceGroup defines model for DeviceGroup.
type DeviceGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Device defines model for Device.
type Device struct {
	ID                 string        `json:"id"`   // IMEI and/or device_serial_number
	Name               string        `json:"name"` // Display Name
	Description        string        `json:"description"`
	Attached           bool          `json:"attached"`              // sd-core API
	IP                 string        `json:"ip"`                    // sd-core API
	SmallCellID        string        `json:"small_cell_id"`         // sd-core API
	CurrentMbpsUp      string        `json:"current_mbps_up"`       // Prometheus
	CurrentMbpsDown    string        `json:"current_mbps_down"`     // Prometheus
	Health             string        `json:"health"`                // ("healthy" | "unhealthy" | "warning") sd-fabric
	NumFlows           string        `json:"num_flows"`             //  sd-fabric
	NumFlowsWithIssues string        `json:"num_flows_with_issues"` //  sd-fabric
	DeviceGroups       []DeviceGroup `json:"device_groups"`
	SIMIccID           string        `json:"sim_iccid"`
}

// DeviceGroups - temporary
var DeviceGroups = []DeviceGroup{
	{ID: "1", Name: "device-group-1"},
}

// Devices - temporary
var Devices = []Device{
	{
		ID:                 "1",
		Name:               "camera-1",
		Description:        "camera-1 description",
		Attached:           true,
		IP:                 "192.123.45.678",
		SmallCellID:        "",
		CurrentMbpsUp:      "",
		CurrentMbpsDown:    "",
		Health:             "",
		NumFlows:           "",
		NumFlowsWithIssues: "",
		DeviceGroups:       DeviceGroups,
		SIMIccID:           "",
	},
}

// Applications - temporary
var Applications = []Application{
	{ID: "1", Name: "app-1"},
	{ID: "2", Name: "app-2"},
}
