package main

import "herman-technical-julo/internal/app"

func setupAppContainer(dbs *databases) *app.Application {
	services := setupServices(dbs)
	return &app.Application{
		Services: services,
	}
}

func setupServices(dbs *databases) *app.Services {
	return &app.Services{}
}
