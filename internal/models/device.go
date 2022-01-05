package models

// Device defines model for Device.
type Device struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Imei         string `json:"imei"`
	Location     string `json:"location,omitempty"`
	SerialNumber string `json:"serial_number"`
	Sim          string `json:"sim,omitempty"`
	Type         string `json:"type"`
	IP           string `json:"ip"`
	Imsi         string `json:"imsi"`
	DeviceGroup  string `json:"device_group"`
}

var Devices = []Device{
	{ID: "1", Name: "camera-1", Imei: "", Location: "", SerialNumber: "", Sim: "sim", Type: "camera", IP: "", Imsi: "", DeviceGroup: ""},
	{ID: "2", Name: "camera-2", Imei: "", Location: "", SerialNumber: "", Sim: "sim", Type: "camera", IP: "", Imsi: "", DeviceGroup: ""},
}
