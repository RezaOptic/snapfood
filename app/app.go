// Package app is the initialization of all the dependencies used in the project
package app

import (
	"database/sql"
	"math/rand"
	"snapfood/config"
	"snapfood/controller"
	"snapfood/logic"
	"snapfood/repo"
	"snapfood/services"
	"time"
)

type App struct {
	// databases connections
	PsqlDB            *sql.DB
	ServiceController controller.ServicesInterface
	ServiceDelayLogic logic.DelayLogicInterface
	ServiceTripRepo   repo.TripsRepoInterface
	ServiceOrderRepo  repo.OrdersRepoInterface
	ServiceDelayRepo  repo.DelayRepoInterface
	ThirdPartyService services.Interface
}

func NewApp(c *string) *App {
	config.Init(c)
	app := App{}
	return &app
}
func (app *App) Initialize() {
	app.initServices()
	app.initDatabases()
	app.initRepoLayers()
	app.initLogicLayers()
	app.initController()
	app.initRouter()
}

func (app *App) InitializeTest() *App {
	app.initServices()
	app.initDatabases()
	app.initRepoLayers()
	app.initLogicLayers()
	app.initController()
	return app
}

func (app *App) InitDependencies() *App {
	rand.Seed(time.Now().Unix())
	app.initServices()
	app.initDatabases()
	app.initRepoLayers()
	app.initLogicLayers()
	return app
}

func (app *App) InitRepoCacheTest() *App {
	app.initRedis()
	return app
}
