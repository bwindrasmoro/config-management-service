package controller

import (
	"strconv"

	"bwind.com/config-management-service/model"
	"bwind.com/config-management-service/service"
	"github.com/gofiber/fiber/v2"
)

type ConfigController struct {
	configService service.ConfigService
}

func NewConfigController(configService *service.ConfigService) *ConfigController {
	return &ConfigController{
		configService: *configService,
	}
}

func (e *ConfigController) CreateConfig(c *fiber.Ctx) error {
	var createConfig model.CommonConfig
	if err := c.BodyParser(&createConfig); err != nil {
		response := model.Response{
			Status:  "Fail",
			Message: "Invalid request body",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result := e.configService.CreateConfig(createConfig)
	isSuccess, _ := result["success"].(bool)

	if !isSuccess {
		response := model.Response{
			Status:  "Fail",
			Message: result["message"].(string),
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := model.Response{
		Status:  "Success",
		Data:    result["data"],
		Version: result["version"].(int),
		Message: "Config sucessfully created",
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (e *ConfigController) UpdateConfig(c *fiber.Ctx) error {
	body := c.Body()
	schema := c.Params("schema")

	result := e.configService.UpdateConfig(schema, body)
	isSuccess, _ := result["success"].(bool)

	if !isSuccess {
		response := model.Response{
			Status:  "Fail",
			Message: result["message"].(string),
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := model.Response{
		Status:  "Success",
		Data:    result["data"],
		Version: result["version"].(int),
		Message: schema + " sucessfully updated",
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (e *ConfigController) RollbackConfig(c *fiber.Ctx) error {
	schema := c.Params("schema")
	version := c.Params("version")

	versionInt, err := strconv.Atoi(version)
	if err != nil {
		response := model.Response{
			Status:  "Fail",
			Message: "Old Version must be Integer",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result := e.configService.RollbackConfig(schema, versionInt)
	isSuccess, _ := result["success"].(bool)

	if !isSuccess {
		response := model.Response{
			Status:  "Fail",
			Message: result["message"].(string),
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := model.Response{
		Status:  "Success",
		Data:    result["data"],
		Version: result["version"].(int),
		Message: schema + " sucessfully updated",
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (e *ConfigController) FetchConfig(c *fiber.Ctx) error {
	schema := c.Params("schema")

	result := e.configService.FetchConfig(schema)
	isSuccess, _ := result["success"].(bool)

	if !isSuccess {
		response := model.Response{
			Status:  "Fail",
			Message: result["message"].(string),
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := model.Response{
		Status:  "Success",
		Version: result["version"].(int),
		Data:    result["data"],
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (e *ConfigController) ListConfig(c *fiber.Ctx) error {
	schema := c.Params("schema")

	result := e.configService.ListConfig(schema)
	isSuccess, _ := result["success"].(bool)

	if !isSuccess {
		response := model.Response{
			Status:  "Fail",
			Message: result["message"].(string),
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := model.Response{
		Status: "Success",
		Data:   result["data"],
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
