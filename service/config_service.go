package service

import (
	"encoding/json"
	"fmt"
	"reflect"

	"bwind.com/config-management-service/model"
)

var ConfigStore = make(map[string][]any)

type ConfigService struct{}

func NewConfigService() *ConfigService {
	return &ConfigService{}
}

func (c *ConfigService) CreateConfig(commonConfig model.CommonConfig) map[string]any {
	configName := commonConfig.ConfigName

	_, exists := ConfigStore[configName]
	if exists {
		return map[string]any{
			"success": false,
			"message": "Config already exists",
		}
	}

	var config any
	if err := json.Unmarshal(commonConfig.Data, &config); err != nil {
		return map[string]any{
			"success": false,
			"message": "Failed to unmarshal data",
		}
	}

	configs := []any{config}
	ConfigStore[configName] = configs

	return map[string]any{
		"success": true,
		"version": 1,
		"data":    config,
	}
}

func (c *ConfigService) UpdateConfig(schema string, config []byte) map[string]any {
	prototype, prototypeExists := model.SchemaRegistry[schema]
	if !prototypeExists {
		return map[string]any{
			"success": false,
			"message": "Schema not exists",
		}
	}

	configs, exists := ConfigStore[schema]
	if !exists {
		return map[string]any{
			"success": false,
			"message": "Config for this schema not exists",
		}
	}

	newInstance := reflect.New(reflect.TypeOf(prototype).Elem()).Interface()
	if err := json.Unmarshal(config, newInstance); err != nil {
		return map[string]any{
			"success": false,
			"message": "Failed to parse config",
		}
	}

	configs = append(configs, newInstance)
	ConfigStore[schema] = configs

	return map[string]any{
		"success": true,
		"version": len(configs),
		"data":    newInstance,
	}
}

func (c *ConfigService) RollbackConfig(schema string, oldVersion int) map[string]any {
	_, prototypeExists := model.SchemaRegistry[schema]
	if !prototypeExists {
		return map[string]any{
			"success": false,
			"message": "Schema not exists",
		}
	}

	configs := ConfigStore[schema]
	if oldVersion > len(configs) {
		return map[string]any{
			"success": false,
			"message": fmt.Sprintf("Cannot rollback %s schema: rollback version %d not exists", schema, oldVersion),
		}
	}

	config := configs[oldVersion-1]
	configs = append(configs, config)
	ConfigStore[schema] = configs

	return map[string]any{
		"success": true,
		"version": len(configs),
		"data":    config,
	}
}

func (c *ConfigService) FetchConfig(schema string) map[string]any {
	_, prototypeExists := model.SchemaRegistry[schema]
	if !prototypeExists {
		return map[string]any{
			"success": false,
			"message": "Schema not exists",
		}
	}

	configs, exists := ConfigStore[schema]
	if !exists {
		return map[string]any{
			"success": false,
			"message": fmt.Sprintf("Configuration for %s schema not found", schema),
		}
	}

	return map[string]any{
		"success": true,
		"version": len(configs),
		"data":    configs[len(configs)-1],
	}

}

func (c *ConfigService) ListConfig(schema string) map[string]any {
	_, prototypeExists := model.SchemaRegistry[schema]
	if !prototypeExists {
		return map[string]any{
			"success": false,
			"message": "Schema not exists",
		}
	}

	var configsWithVersion []any

	configs, exists := ConfigStore[schema]
	if !exists {
		return map[string]any{
			"success": false,
			"message": fmt.Sprintf("Configuration for %s schema not found", schema),
		}
	}

	for i, value := range configs {
		configsWithVersion = append(configsWithVersion, map[string]any{
			"config":  value,
			"version": (i + 1),
		})
	}

	return map[string]any{
		"success": true,
		"data":    configsWithVersion,
	}

}
