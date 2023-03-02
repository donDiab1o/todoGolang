package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var _ = godotenv.Load()
var DataBase, dataBaseErr = sql.Open("postgres", getConnectionConfig())
var Port = os.Getenv("PORT")

func main() {

	checkConnectionDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/todo/delete", deleteTask)
	mux.HandleFunc("/todo/add", addTask)
	mux.HandleFunc("/tasks", showTask)

	log.Printf("Server start on %v \n", Port)

	errMuxServer := http.ListenAndServe((":" + Port), mux)
	log.Fatal(errMuxServer)

}

func getConnectionConfig() string {

	user := os.Getenv("USERDB")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOSTDB")
	sslmode := os.Getenv("SSLMODE")

	return fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", user, dbname, password, host, sslmode)
}

func checkConnectionDB() {
	if dataBaseErr != nil {
		panic(dataBaseErr)
	}

	dataBaseErr = DataBase.Ping()
	if dataBaseErr != nil {
		panic(dataBaseErr)
	}

	log.Printf("\nSuccessfully connected to database!\n")

	return
}
