package model

const (
	CONTAINER_STATE_WAITING_TYPE = "v1.ContainerStateWaiting"
)

type ContainerStateWaiting struct {
	Reason string `json:"reason,omitempty" yaml:"reason,omitempty"`
}
