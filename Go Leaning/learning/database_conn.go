package database_conn

import ( "fmt"
"log"
		"database/sql"
		_ "github.com/go-sql-driver/mysql" // "_" is use to execute init of mysql
)
func database_conn(){
	fmt.Println("Connect started")

	db, err := sql.Open("mysql","root:root@tcp(localhost:3306)/biren_db")
	if err != nil{
		fmt.Println("Error: validate sql.Open() arguments")
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to Database.")

}
// go get github.com/go-sql-driver/mysql@v1.6.0
// go mod init micro