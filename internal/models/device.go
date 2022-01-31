// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package models

// Device defines model for Device.
type Device struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	IMEI               int    `json:"imei"`
	Attached           bool   `json:"attached"`              // sd-core (1,0) (Subscribers_info metric)
	IP                 string `json:"ip"`                    // sd-core (Subscribers_info metric)
	SmallCellID        string `json:"small_cell_id"`         // sd-core which metric is this?
	CurrentMbpsUp      string `json:"current_mbps_up"`       // Prometheus
	CurrentMbpsDown    string `json:"current_mbps_down"`     // Prometheus
	Health             string `json:"health"`                // ("healthy" | "unhealthy" | "warning") sd-fabric
	NumFlows           string `json:"num_flows"`             //  sd-fabric
	NumFlowsWithIssues string `json:"num_flows_with_issues"` //  sd-fabric
	DeviceGroups       string `json:"device_groups"`
	SimICCID           int    `json:"sim_iccid"`
}
