package queries

var (
	VendorsTable = `CREATE TABLE IF NOT EXISTS Vendors (
    vendor_id SERIAL PRIMARY KEY,
    vendor_name VARCHAR(255)
);`
	UsersTable = `CREATE TABLE IF NOT EXISTS Users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255)
);`
	AgentsTable = `CREATE TABLE IF NOT EXISTS Agents (
    agent_id SERIAL PRIMARY KEY,
    username VARCHAR(255)
);`
	OrdersTable = `CREATE TABLE IF NOT EXISTS Orders (
    order_id SERIAL PRIMARY KEY,
    user_id INT,
    vendor_id INT,
    order_time TIMESTAMP,
    time_delivery INT,
    FOREIGN KEY (user_id) REFERENCES Users(user_id),
    FOREIGN KEY (vendor_id) REFERENCES Vendors(vendor_id)
);`
	OrdersIndex1 = `CREATE INDEX idx_user_id ON Orders(user_id);`
	OrdersIndex2 = `CREATE INDEX idx_vendor_id ON Orders(vendor_id);`
	OrdersIndex3 = `CREATE INDEX idx_order_time ON Orders(order_time);`

	TripsTable = `CREATE TABLE IF NOT EXISTS Trips (
    trip_id SERIAL PRIMARY KEY,
    order_id INT,
    courier_id INT,
    trip_status VARCHAR(20), 
    FOREIGN KEY (order_id) REFERENCES Orders(order_id),
    FOREIGN KEY (courier_id) REFERENCES Couriers(courier_id)
);`
	TripsIndex1 = `CREATE INDEX idx_order_id ON Trips(order_id);`
	TripsIndex2 = `CREATE INDEX idx_courier_id ON Trips(courier_id);`

	CouriersTable = `CREATE TABLE IF NOT EXISTS Couriers (
    courier_id SERIAL PRIMARY KEY,
    courier_name VARCHAR(255)
);`
	DelayReportsTable = `CREATE TABLE IF NOT EXISTS DelayReports (
    report_id SERIAL PRIMARY KEY,
    order_id INT,
    user_id INT, 
    agent_id INT, 
    vendor_id INT, 
    delay_reason TEXT,
    report_time TIMESTAMP,
    state VARCHAR(20),
    FOREIGN KEY (order_id) REFERENCES Orders(order_id),
    FOREIGN KEY (user_id) REFERENCES Users(user_id),
    FOREIGN KEY (agent_id) REFERENCES Agents(agent_id),
    FOREIGN KEY (vendor_id) REFERENCES Vendors(vendor_id)
);`
	DelayReportsIndex1 = `CREATE INDEX idx_order_id ON DelayReports(order_id);`
	DelayReportsIndex2 = `CREATE INDEX idx_user_id ON DelayReports(user_id);`
	DelayReportsIndex3 = `CREATE INDEX idx_report_time ON DelayReports(report_time);`
	DelayReportsIndex4 = `CREATE INDEX idx_agent_id ON DelayReports(agent_id);`
	DelayReportsIndex5 = `CREATE INDEX idx_vendor_id ON DelayReports(vendor_id);`
)
