package model

const (
	LIMIT_RANGE_ITEM_TYPE = "v1.LimitRangeItem"
)

type LimitRangeItem struct {
	Default map[string]interface{} `json:"default,omitempty" yaml:"default,omitempty"`

	Max map[string]interface{} `json:"max,omitempty" yaml:"max,omitempty"`

	Min map[string]interface{} `json:"min,omitempty" yaml:"min,omitempty"`

	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
