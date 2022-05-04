package model

type NimbyStatus interface{}

type nimbyStatus struct {
	Value  *bool   `json:"value"`
	Mode   *string `json:"mode"`
	Reason *string `json:"reason"`
}

func NewNimbyStatus() nimbyStatus {
	return nimbyStatus{&[]bool{true}[0], &[]string{"auto"}[0], &[]string{"Default status"}[0]}
}

func (status *nimbyStatus) Merge(otherStatus *nimbyStatus) {
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

func (status *nimbyStatus) GetValue() bool {
	return *status.Value
}

func (status *nimbyStatus) GetMode() string {
	return *status.Mode
}

func (status *nimbyStatus) GetReason() string {
	return *status.Reason
}
