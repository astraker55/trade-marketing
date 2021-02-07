package queries

// TruncateQuery is using for dropstate request
var TruncateQuery string = "TRUNCATE events"

// SelectQuery is QueryRow for getting info about events
var SelectQuery string = `SELECT date, views, clicks, cost, 
CASE 
	WHEN clicks != 0 THEN cost/clicks
	ELSE 0
END cpc, 
CASE 
	WHEN clicks != 0 THEN cost/views*1000
	ELSE 0
END cpm
FROM events 
WHERE DATE >= to_timestamp($1) and DATE <= to_timestamp($2) 
ORDER BY $3`

// InsertQuery for put data in database
var InsertQuery = `INSERT INTO events (date, views, clicks, cost)
 VALUES(:date, :views, :clicks, :cost)
`
