package main

import (
	"bwind.com/config-management-service/app"
)

func main() {
	app := app.NewApp()
	app.Listen(":3001")
}
