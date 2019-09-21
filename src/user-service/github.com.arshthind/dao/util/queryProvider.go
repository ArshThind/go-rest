package util

import (
	"io/ioutil"
	"log"
	"os"
)

type query string

const (
	GET_USERS   query = "Absolute path to sql file"
	ADD_USER    query = "Absolute path to sql file"
	DELETE_USER query = "Absolute path to sql file"
	UPDATE_USER query = "Absolute path to sql file"
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
