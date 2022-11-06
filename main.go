package main

import (
	"Banking/app"
	"Banking/logger"
)

func main() {
	logger.Info("Start application...")
	app.Start()
}
