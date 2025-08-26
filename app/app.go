package app

import (
	"encoding/json"
	"reflect"

	"bwind.com/config-management-service/controller"
	"bwind.com/config-management-service/helper"
	"bwind.com/config-management-service/model"
	"bwind.com/config-management-service/service"
	"github.com/gofiber/fiber/v2"
)

// Define API path, end-point, and controller mapping here
func NewApp() *fiber.App {
	app := fiber.New()

	//Load Service
	configService := service.NewConfigService()

	//Load Controller
	configController := controller.NewConfigController(configService)

	//Mapping end-point with controller
	configRouter := app.Group("/api/v1/config")
	configRouter.Post("/", ValidateFieldsMiddleware("create"), configController.CreateConfig)
	configRouter.Post("/:schema", ValidateFieldsMiddleware("update"), configController.UpdateConfig)
	configRouter.Post("/:schema/rollback/:version", configController.RollbackConfig)
	configRouter.Get("/:schema", configController.FetchConfig)
	configRouter.Get("/:schema/versions", configController.ListConfig)

	return app
}

// Validate the schema
func ValidateFieldsMiddleware(mode string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var prototype any
		var data []byte

		switch mode {
		case "create":
			var createConfig model.CommonConfig
			if err := c.BodyParser(&createConfig); err != nil {
				response := model.Response{
					Status:  "Fail",
					Message: "Invalid request body",
				}
				return c.Status(fiber.StatusBadRequest).JSON(response)
			}

			missingFields := helper.ValidateStruct(createConfig)
			if len(missingFields) > 0 {
				response := model.Response{
					Status:  "Fail",
					Message: "One or more field are required",
					Data:    missingFields,
				}

				return c.Status(fiber.StatusBadRequest).JSON(response)
			}

			proto, ok := model.SchemaRegistry[createConfig.ConfigName]
			if !ok {
				response := model.Response{
					Status:  "Fail",
					Message: "Unsupported config type: " + createConfig.ConfigName,
				}
				return c.Status(fiber.StatusBadRequest).JSON(response)
			}

			prototype = proto
			data = createConfig.Data
		case "update":
			configName := c.Params("schema")

			proto, ok := model.SchemaRegistry[configName]
			if !ok {
				response := model.Response{
					Status:  "Fail",
					Message: "Unsupported config type: " + configName,
				}
				return c.Status(fiber.StatusBadRequest).JSON(response)
			}

			prototype = proto
			data = c.Body()
		}

		newInstance := reflect.New(reflect.TypeOf(prototype).Elem()).Interface()
		if err := json.Unmarshal(data, newInstance); err != nil {
			response := model.Response{
				Status:  "Fail",
				Message: "invalid config data",
			}
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		missingFields := helper.ValidateStruct(newInstance)
		if len(missingFields) > 0 {
			response := model.Response{
				Message: "One or more field are required",
				Data:    missingFields,
			}

			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		return c.Next()
	}
}
