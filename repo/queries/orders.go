package queries

var (
	FindOrders = `SELECT  order_id,user_id,vendor_id,order_time, time_delivery FROM Orders WHERE order_id = $1`
)
