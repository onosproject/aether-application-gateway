package models

// Sim defines model for Sim.
type Sim struct {
	DisplayName *string `json:"display-name,omitempty" yaml:"display-name,omitempty"`
	Iccid       string  `json:"iccid"`
}
