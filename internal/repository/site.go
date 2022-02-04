// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package repository

// Site is the response type from getting a site
// aether/v2.0.0/:target/enterprises/enterprise/:ent-id/site/:site-id
type Site struct {
	Description string `json:"description"`
	Device      []struct {
		Description string `json:"description"`
		DevID       string `json:"dev-id"`
		DisplayName string `json:"display-name"`
		Imei        int    `json:"imei"`
		SimCard     string `json:"sim-card"`
	} `json:"device"`
	DeviceGroup []struct {
		Device      []interface{} `json:"device"`
		DgID        string        `json:"dg-id"`
		DisplayName string        `json:"display-name"`
		IPDomain    string        `json:"ip-domain"`
		Mbr         struct {
			Downlink int `json:"downlink"`
			Uplink   int `json:"uplink"`
		} `json:"mbr"`
		TrafficClass string `json:"traffic-class"`
	} `json:"device-group"`
	DisplayName    string `json:"display-name"`
	ImsiDefinition struct {
		Enterprise int    `json:"enterprise"`
		Format     string `json:"format"`
		Mcc        string `json:"mcc"`
		Mnc        string `json:"mnc"`
	} `json:"imsi-definition"`
	IPDomain []struct {
		AdminStatus  string `json:"admin-status"`
		Description  string `json:"description"`
		DisplayName  string `json:"display-name"`
		Dnn          string `json:"dnn"`
		DNSPrimary   string `json:"dns-primary"`
		DNSSecondary string `json:"dns-secondary"`
		IPID         string `json:"ip-id"`
		Mtu          int    `json:"mtu"`
		Subnet       string `json:"subnet"`
	} `json:"ip-domain"`
	Monitoring struct {
		EdgeClusterPrometheusURL string `json:"edge-cluster-prometheus-url"`
		EdgeDevice               []struct {
			Description  string `json:"description"`
			DisplayName  string `json:"display-name"`
			EdgeDeviceID string `json:"edge-device-id"`
		} `json:"edge-device"`
		EdgeMonitoringPrometheusURL string `json:"edge-monitoring-prometheus-url"`
	} `json:"monitoring"`
	PriorityTrafficRule []struct {
		Application string `json:"application"`
		Description string `json:"description"`
		Device      string `json:"device"`
		DisplayName string `json:"display-name"`
		Endpoint    string `json:"endpoint"`
		Gbr         struct {
			Downlink int `json:"downlink"`
			Uplink   int `json:"uplink"`
		} `json:"gbr"`
		Mbr struct {
			Downlink int `json:"downlink"`
			Uplink   int `json:"uplink"`
		} `json:"mbr"`
		PtrID        string `json:"ptr-id"`
		TrafficClass string `json:"traffic-class"`
	} `json:"priority-traffic-rule"`
	SiteID string `json:"site-id"`
	Slice  []struct {
		DefaultBehavior string `json:"default-behavior"`
		Description     string `json:"description"`
		DeviceGroup     []struct {
			DeviceGroup string `json:"device-group"`
			Enable      bool   `json:"enable"`
		} `json:"device-group"`
		DisplayName string `json:"display-name"`
		Filter      []struct {
			Allow       bool   `json:"allow"`
			Application string `json:"application"`
		} `json:"filter"`
		Mbr struct {
			Downlink          int `json:"downlink"`
			DownlinkBurstSize int `json:"downlink-burst-size"`
		} `json:"mbr"`
		Sd      int    `json:"sd"`
		SliceID string `json:"slice-id"`
		Sst     int    `json:"sst"`
		Upf     string `json:"upf"`
	} `json:"slice"`
	SimCard   []SimCard `json:"sim-card"`
	SmallCell []struct {
		Address     string `json:"address"`
		DisplayName string `json:"display-name"`
		Enable      bool   `json:"enable"`
		SmallCellID string `json:"small-cell-id"`
		Tac         string `json:"tac"`
	} `json:"small-cell"`
	Upf []struct {
		Address     string `json:"address"`
		Description string `json:"description"`
		DisplayName string `json:"display-name"`
		Port        int    `json:"port"`
		UpfID       string `json:"upf-id"`
	} `json:"upf"`
}

// SimCard -
type SimCard struct {
	Description string `json:"description"`
	DisplayName string `json:"display-name"`
	Iccid       int    `json:"iccid"`
	Imsi        int    `json:"imsi"`
	SimID       string `json:"sim-id"`
}
