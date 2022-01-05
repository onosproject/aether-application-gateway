package models

// Slice defines model for Slice.
type Slice struct {
	Applications *[]string `json:"applications,omitempty"`
	DeviceGroups *[]string `json:"device-groups,omitempty" yaml:"device-groups,omitempty"`
	DisplayName  string    `json:"display-name" yaml:"display-name"`
	SliceId      string    `json:"slice-id,omitempty" yaml:"slice-id,omitempty"`
}
