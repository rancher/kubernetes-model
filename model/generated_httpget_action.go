package model

const (
	HTTPGET_ACTION_TYPE = "v1.HTTPGetAction"
)

type HTTPGetAction struct {
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	Port int32 `json:"port,omitempty" yaml:"port,omitempty"`

	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty"`
}
