package model

import "encoding/json"

var SchemaRegistry = map[string]any{
	"payment_config": &PaymentConfig{},
	"partner_config": &PartnerConfig{},
}

type CommonConfig struct {
	ConfigName string          `json:"config_name" validation:"required"`
	Data       json.RawMessage `json:"data" validation:"required"`
}

type PaymentConfig struct {
	MaxLimit int  `json:"max_limit" validation:"required"`
	Enabled  bool `json:"enabled" validation:"required"`
}

type PartnerConfig struct {
	Name    string `json:"name" validate:"required"`
	Enabled bool   `json:"enabled" validate:"required"`
}
