package config

import(
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

//connect to the database and return a pointer of the database as an instancr
func ConnectDatabase()(*sql.DB,error){
	host := "localhost"
	port := 3306
	user := "root"
	password := "gr00t@00"
	dbname := "eshop"
	

	//Create database connection string
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)

//Open connection to the database
db,err:=sql.Open("mysql",connStr)
if err!=nil{
	return nil ,err
}

//Ping the database to test the connection status
err=db.Ping()
if err!=nil{
	return nil,err
}
return db,err
}