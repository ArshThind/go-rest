package util

import (
	"io/ioutil"
	"log"
	"os"
)

type query string

const (
	GET_USERS   query = "F:\\golang\\src\\user-service\\github.com.arshthind\\resources\\queries\\get-user-query.sql"
	ADD_USER    query = "F:\\golang\\src\\user-service\\github.com.arshthind\\resources\\queries\\add-user-query.sql"
	DELETE_USER query = "F:\\golang\\src\\user-service\\github.com.arshthind\\resources\\queries\\delete-user-query.sql"
	UPDATE_USER query = "F:\\golang\\src\\user-service\\github.com.arshthind\\resources\\queries\\update-user-query.sql"
)

var queries = make(map[query]string)

func init() {
	queries[GET_USERS] = loadQuery(GET_USERS)
	queries[ADD_USER] = loadQuery(ADD_USER)
	queries[DELETE_USER] = loadQuery(DELETE_USER)
	queries[UPDATE_USER] = loadQuery(UPDATE_USER)
	log.Println("Queries loaded successfully")
}

func loadQuery(query query) string {
	f, err := os.Open(string(query))
	if err != nil {
		log.Print("Error occurred while loading the query", err)
		panic("Error occurred while loading the query")
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Print("Error occurred while reading the query", err)
		panic("Error occurred while reading the query")
	}
	return string(data)
}

func GetQuery(query query) string {
	return queries[query]
}
