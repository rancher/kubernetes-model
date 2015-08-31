package model

const (
	SERVICE_SPEC_TYPE = "v1.ServiceSpec"
)

type ServiceSpec struct {
	ClusterIP string `json:"clusterIP,omitempty" yaml:"cluster_ip,omitempty"`

	DeprecatedPublicIPs []string `json:"deprecatedPublicIPs,omitempty" yaml:"deprecated_public_ips,omitempty"`

	Ports []ServicePort `json:"ports,omitempty" yaml:"ports,omitempty"`

	Selector map[string]interface{} `json:"selector,omitempty" yaml:"selector,omitempty"`

	SessionAffinity string `json:"sessionAffinity,omitempty" yaml:"session_affinity,omitempty"`

	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
