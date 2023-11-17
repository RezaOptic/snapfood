package app

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"snapfood/config"
	"snapfood/repo/queries"
)

var (
	errorSent bool
)

func (app *App) initDatabases() {
	app.initPsql()
}

func (app *App) initPsql() {

	/// Construct the connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s",
		config.PsqlDB.Host, config.PsqlDB.Port, config.PsqlDB.User, config.PsqlDB.Password, config.PsqlDB.Sslmode)

	// Open a database connection
	var err error
	app.PsqlDB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Test the connection
	err = app.PsqlDB.Ping()
	if err != nil {
		panic(err)
	}
	err = createTables(app.PsqlDB)
	if err != nil {
		panic(err)
	}
}

func createTables(db *sql.DB) error {
	// Create the Instrument table
	_, err := db.Exec(queries.VendorsTable)
	if err != nil {
		return err
	}
	_, err = db.Exec(queries.UsersTable)
	if err != nil {
		return err
	}
	_, err = db.Exec(queries.AgentsTable)
	if err != nil {
		return err
	}
	_, err = db.Exec(queries.CouriersTable)
	if err != nil {
		return err
	}
	_, err = db.Exec(queries.OrdersTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(queries.TripsTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(queries.DelayReportsTable)
	if err != nil {
		return err
	}

	return err
}
func (app *App) initRedis() {
	//// logger
	//defer logger.ZSLogger.Sync()
	//app.Redis = redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "", // no password set
	//	DB:       0,  // use default DB
	//})
	//
	//if err := app.Redis.Ping(context.Background()); err != nil {
	//	logger.ZSLogger.Panic(err)
	//}
	//
	//logger.ZSLogger.Info(RedisSuccess)
}
