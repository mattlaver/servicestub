package main

import (
    "fmt"
)

func main() {
	var settings Settings = &SettingsImpl{}
    settings.Load("conf.json")
	fmt.Println(settings.PortNo())

	var routes Routes = &RoutesImpl{}
	routes.Load()

	var service Service = &ServiceImpl{}
	service.InitRoutes(routes.RouteList())
	service.Start(settings)
}
