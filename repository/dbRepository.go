package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

/*
	connect to the Db
	sql.Open does not open a connection. A connection pool is maintained
*/
func init() {
	var err error
	/*
		http://go-database-sql.org/accessing.html
		Perhaps counter-intuitively, sql.Open() does not establish any connections to the database, nor does it validate
		driver connection parameters. Instead, it simply prepares the database abstraction for later use. The first actual
		connection to the underlying datastore will be established lazily, when it’s needed for the first time. If you want
		to check right away that the database is available and accessible (for example, check that you can establish a
		network connection and log in), use db.Ping() to do that, and remember to check for errors:

		Although it’s idiomatic to Close() the database when you’re finished with it, the sql.DB object is designed to be
		long-lived. Don’t Open() and Close() databases frequently. Instead, create one sql.DB object for each distinct
		datastore you need to access, and keep it until the program is done accessing that datastore. Pass it around as
		needed, or make it available somehow globally, but keep it open. And don’t Open() and Close() from a short-lived
		function. Instead, pass the sql.DB into that short-lived function as an argument.

		If you don’t treat the sql.DB as a long-lived object, you could experience problems such as poor reuse and sharing
		of connections, running out of available network resources, or sporadic failures due to a lot of TCP connections
		remaining in TIME_WAIT status. Such problems are signs that you’re not using database/sql as it was designed.


	*/
	Db, err = sql.Open("mysql", "root:root@/sample_jersey_app?autocommit=false")
	if err != nil {
		panic(err)
	}
}
