package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "first_db"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disalbe", host, port, user, password, dbname)
	db, err := sql.Open("postgress", psqlconn)
	CheckError(err)
	defer db.Close()
	insertStmt := `insert into "Employe"("Name", "EmpID") values('ahmed',21)`
	_, e := db.Exec(insertStmt)
	CheckError(e)
	insertDynStmt := `insert into "Employee" ("Name", "EmpID") values($1, $2)`
	_, e = db.Exec(insertDynStmt, "Sheraz", 16)
	CheckError(e)
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
