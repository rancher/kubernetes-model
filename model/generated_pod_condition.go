package model

const (
	POD_CONDITION_TYPE = "v1.PodCondition"
)

type PodCondition struct {
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
