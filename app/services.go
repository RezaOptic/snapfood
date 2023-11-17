package app

import "snapfood/services"

func (app *App) initServices() {
	app.ThirdPartyService = services.NewServices()
}
