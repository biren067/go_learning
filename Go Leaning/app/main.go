package main

import ( "fmt"
"log"
		"database/sql"
		_ "github.com/go-sql-driver/mysql" // "_" is use to execute init of mysql
)
func main(){
	fmt.Println("Connect started")

	db, err := sql.Open("mysql","root:root@tcp(localhost:3306)/biren_db")
	if err != nil{
		fmt.Println("Error: validate sql.Open() arguments")
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} 	else{
		fmt.Println("Connected to Database.")
	}
	
	
	id :=2
	name :="James"
	sqlStatement := "Insert into facility(id,name) values(?,?)"
	result, err := db.Exec(sqlStatement,id,name)
	if err!=nil{
		fmt.Println("Error: Not inserting data into database")
	} else {
		fmt.Println("Data is inserted into database.",result)
	}
}

// go get github.com/go-sql-driver/mysql@v1.6.0
// go mod init micro