package main

import "url-shortener/internal/app"

const appConfigurationPath = "configs/app.json"

func main() {
	app.Run(appConfigurationPath)
}
