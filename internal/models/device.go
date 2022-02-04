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
	Attached           int    `json:"attached"`                        // sd-core (1,0) (Subscribers_info metric)
	IP                 string `json:"ip,omitempty"`                    // sd-core (Subscribers_info metric)
	SmallCellID        string `json:"small_cell_id,omitempty"`         // sd-core which metric is this?
	CurrentMbpsUp      string `json:"current_mbps_up,omitempty"`       // Prometheus
	CurrentMbpsDown    string `json:"current_mbps_down,omitempty"`     // Prometheus
	Health             string `json:"health,omitempty"`                // sd-fabric ("healthy" | "unhealthy" | "warning")
	NumFlows           string `json:"num_flows,omitempty"`             //  sd-fabric
	NumFlowsWithIssues string `json:"num_flows_with_issues,omitempty"` //  sd-fabric
	DeviceGroups       string `json:"device_groups"`
	SimICCID           int    `json:"sim_iccid"`
}
