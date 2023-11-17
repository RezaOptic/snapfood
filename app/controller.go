package app

import "snapfood/controller"

func (app *App) initController() {
	app.ServiceController = controller.NewServices(app.ServiceDelayLogic)
}
