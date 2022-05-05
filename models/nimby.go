package models

// type NimbyStatus interface {}

type NimbyStatus struct {
	Value  *bool   `json:"value"`
	Mode   *string `json:"mode"`
	Reason *string `json:"reason"`
}

func NewNimbyStatus() *NimbyStatus {
	return &NimbyStatus{&[]bool{true}[0], &[]string{"auto"}[0], &[]string{"Default status"}[0]}
}

func (status *NimbyStatus) Merge(otherStatus *NimbyStatus) {
	if otherStatus.Value != nil {
		status.Value = otherStatus.Value
	}

	if otherStatus.Mode != nil {
		status.Mode = otherStatus.Mode
	}

	if otherStatus.Reason != nil {
		status.Reason = otherStatus.Reason
	}
}

func (status *NimbyStatus) GetValue() bool {
	return *status.Value
}

func (status *NimbyStatus) GetMode() string {
	return *status.Mode
}

func (status *NimbyStatus) GetReason() string {
	return *status.Reason
}
