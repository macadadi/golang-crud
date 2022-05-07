package dbhandler

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DBconnect()(db *sql.DB){
	if err := godotenv.Load(".env");err != nil{
		log.Fatal("could not load environment variables")
		return
	}
	dbinfo := fmt.Sprintf("host=%v port=%v dbname=%v user=%v sslmode=disable",os.Getenv("host"),os.Getenv("port"),os.Getenv("dbname"),os.Getenv("user"))

	db, err := sql.Open("postgres",dbinfo)
	if err != nil{
		log.Fatal(" Could not open database")
		return
	}
	if err := db.Ping();err != nil{
		panic(err)
	}
	return db
}