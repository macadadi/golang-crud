package services

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	Id         int    `json:"id"`
	Age        int    `json:"age"`
	Last_name  string `json:"lastName"`
	First_name string `json:"firstName"`
}

func GetUsers(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var albums []album
		sqlQuery := "SELECT * FROM users ORDER BY id"
		data, err := db.Query(sqlQuery)
		if err != nil {
			panic(err)
		}
		for data.Next() {
			var album album
			if err := data.Scan(&album.Id, &album.Age, &album.First_name, &album.Last_name); err != nil {
				log.Fatal("could not fetch data")
			}
			albums = append(albums, album)

		}
		c.IndentedJSON(http.StatusOK, albums)

	}
}

func GetSingleUser(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var album album
		id := c.Param("id")
		sqlquery := "SELECT * FROM users WHERE id= $1"

		data, err := db.Query(sqlquery, id)

		if err != nil {
			log.Fatal("could not fetch a record with that id")
		}
		for data.Next() {
			if err := data.Scan(&album.Id, &album.Age, &album.Last_name, &album.First_name); err != nil {
				log.Fatal("no data was found")
				return
			}

		}
		c.IndentedJSON(http.StatusOK, album)
	}
}

func AddNewUser(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var album album
		sqlquery := "INSERT INTO users(Age,Last_name,First_name) VALUES($1,$2,$3)"
		if err := c.BindJSON(&album); err != nil {
			log.Fatal(err)
		}
		_, err := db.Exec(sqlquery, album.Age, album.Last_name, album.First_name)
		if err != nil {
			log.Fatalf("Error occured %v", err)
			return
		}
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "successfully added the data"})
	}
}

func UpdateUser(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var album album
		id := c.Param("id")
		sqlquery := "UPDATE users SET age=$1,Last_name=$2 ,First_name=$3 WHERE id=$4"

		if err := c.BindJSON(&album); err != nil {
			log.Fatalf("binding err %v", err)
		}
		_, err := db.Exec(sqlquery, album.Age, album.Last_name, album.First_name, id)

		if err != nil {
			log.Fatalf("execution error %v", err)
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "successfully updated the data"})
	}
}

func Delete(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		sqlquery := "DELETE FROM users WHERE id=$1"
		_, err := db.Exec(sqlquery, id)
		if err != nil {
			log.Fatalf("could not delete the record %v", err)

		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "successfully deleted the record"})
	}
}
