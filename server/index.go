package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	if dbErr != nil && dbErr != db.Ping() {
		panic(dbErr)
	}
	defer db.Close()


	log.Print("connected to db")

	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/todo/delete", deleteTask)
	mux.HandleFunc("/todo/update", updateTask)
	mux.HandleFunc("/todo/add", addTask)
	mux.HandleFunc("/tasks", showTask)

	log.Printf("Server start on %d \n", 4000)

	err1 := http.ListenAndServe(":4000", mux)
	log.Fatal(err1)

}

var ConnStr = "user=process.env.DB_NAME dbname=todo password=root host=localhost sslmode=disable"
var db, dbErr = sql.Open("postgres", ConnStr)
