package model

type NimbyStatus struct {
	Value  *bool   `json:"value"`
	Mode   *string `json:"mode"`
	Reason *string `json:"reason"`
}
