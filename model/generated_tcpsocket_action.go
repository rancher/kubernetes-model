package model

const (
	TCPSOCKET_ACTION_TYPE = "v1.TCPSocketAction"
)

type TCPSocketAction struct {
	Port string `json:"port,omitempty" yaml:"port,omitempty"`
}
