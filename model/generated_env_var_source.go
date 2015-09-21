package model

const (
	ENV_VAR_SOURCE_TYPE = "v1.EnvVarSource"
)

type EnvVarSource struct {
	FieldRef *ObjectFieldSelector `json:"fieldRef,omitempty" yaml:"field_ref,omitempty"`
}
