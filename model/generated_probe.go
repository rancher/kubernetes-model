package model

const (
	PROBE_TYPE = "v1.Probe"
)

type Probe struct {
	Exec ExecAction `json:"exec,omitempty" yaml:"exec,omitempty"`

	HttpGet HTTPGetAction `json:"httpGet,omitempty" yaml:"http_get,omitempty"`

	InitialDelaySeconds int64 `json:"initialDelaySeconds,omitempty" yaml:"initial_delay_seconds,omitempty"`

	TcpSocket TCPSocketAction `json:"tcpSocket,omitempty" yaml:"tcp_socket,omitempty"`

	TimeoutSeconds int64 `json:"timeoutSeconds,omitempty" yaml:"timeout_seconds,omitempty"`
}
