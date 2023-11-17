package app

import "snapfood/logic"

func (app *App) initLogicLayers() {
	app.ServiceDelayLogic = logic.NewDelayLogic(app.ServiceTripRepo, app.ThirdPartyService, app.ServiceOrderRepo, app.ServiceDelayRepo)
}
