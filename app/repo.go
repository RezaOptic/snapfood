package app

import "snapfood/repo"

func (app *App) initRepoLayers() {
	app.ServiceTripRepo = repo.NewTripsRepo(app.PsqlDB)
	app.ServiceOrderRepo = repo.NewOrdersRepo(app.PsqlDB)
	app.ServiceDelayRepo = repo.NewDelayRepo(app.PsqlDB)
}
