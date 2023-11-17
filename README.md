To run the web service of the project, you can use the following command:

```bash
go run ./console/main.go http_server
```

In the above command, http_server is the command name.

1. Delay Service API

```bash
curl --location '127.0.0.1:9170/v1/service/delay' \
--form 'order_id="1"' \
--form 'user_id="1"'
```

Through this API, users can report a delay for their orders.

2. Assign Service API

```bash
curl --location '127.0.0.1:9170/v1/service/assign' \
--form 'agent_id="1"'
```

This API allows an agent to assign themselves the first order in the delay queue.

3. Reports API

```bash
curl --location '127.0.0.1:9170/v1/service/reports?vendor_id=1&from=2023-11-17T00%3A00%3A00Z&to=2023-11-17T23%3A59%3A59Z'
```

Through this API, you can retrieve a list of delays for a specific vendor within a specified time range.

To set up this project, you need a PostgreSQL database, and it is necessary to modify the database-related information
in the configuration file located at the following path:

``` 
./config/files/psql.toml