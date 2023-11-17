package domain

type Vendor struct {
	VendorID   int    `json:"vendor_id"`
	VendorName string `json:"vendor_name"`
}

type Order struct {
	OrderID      int    `json:"order_id"`
	UserID       int    `json:"user_id"`
	VendorID     int    `json:"vendor_id"`
	OrderTime    string `json:"order_time"`
	TimeDelivery int    `json:"time_delivery"`
}

type User struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}

type Trip struct {
	TripID     int    `json:"trip_id"`
	OrderID    int    `json:"order_id"`
	CourierID  int    `json:"courier_id"`
	TripStatus string `json:"trip_status"`
}

type Courier struct {
	CourierID   int    `json:"courier_id"`
	CourierName string `json:"courier_name"`
}

type DelayReport struct {
	ReportID    int    `json:"report_id"`
	OrderID     int    `json:"order_id"`
	UserID      int    `json:"user_id"`
	AgentID     *int   `json:"agent_id"`
	VendorID    int    `json:"vendor_id"`
	DelayReason string `json:"delay_reason"`
	ReportTime  string `json:"report_time"`
	State       string `json:"state"`
}
