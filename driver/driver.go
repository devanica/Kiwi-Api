package driver

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
	"os"
)

//DRIVER it is there to open database connection
var db *sql.DB

func logFatal(err error){
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB{
	pgUrl, err := pq.ParseURL(os.Getenv("POSTGRES_URL"))
	logFatal(err)

	log.Println(pgUrl)
	/*this gave us:
	dbname=kiwi
	host=localhost
	password=postgres
	port=5432
	user=postgres

	Open opens a database specified by its database driver name and a
	driver-specific data source name, usually consisting of at least a
	database name and connection information.*/
	db, err = sql.Open("postgres", "user=postgres password=postgres dbname=kiwi sslmode=disable")
	/*Open may just validate its arguments without creating a connection
	to the database. To verify that the data source name is valid, call
	Ping.*/
	err = db.Ping()
	//If there is an err the execution stops on the prev method, if not it continues.
	logFatal(err)

	return db
}