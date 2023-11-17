package queries

var (
	FindOrderTrip = `SELECT  trip_id,order_id,courier_id,trip_status FROM Trips WHERE order_id = $1`
)
