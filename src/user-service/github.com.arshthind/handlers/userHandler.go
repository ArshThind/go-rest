package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
	"user-service/github.com.arshthind/dao"
	"user-service/github.com.arshthind/model"
)

const (
	GET    string = "GET"
	PUT    string = "PUT"
	POST   string = "POST"
	DELETE string = "DELETE"
)

var numRegex *regexp.Regexp = regexp.MustCompile("[0-9]+")

type UserHandler struct{}

func (handler UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	switch r.Method {
	case GET:
		fetchUsers(w)
	case PUT:
		updateUser(w, r)
	case POST:
		createUser(w, r)
	case DELETE:
		deleteUser(w, r)
	default:
		w.WriteHeader(405)
	}
}

func fetchUsers(w http.ResponseWriter) {
	start := time.Now()
	users, err := dao.GetUsers()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Println("Error occurred while serializing data", err)
		w.WriteHeader(500)
		return
	}
	defer log.Println("Time taken: ", time.Since(start))
	log.Println(w.Write(jsonData))
}

func createUser(writer http.ResponseWriter, request *http.Request) {
	var user = model.User{}
	data, err := extractBody(request, &user)
	if err != nil {
		writer.WriteHeader(500)
		return
	}

	if _, err := dao.AddUser(&user); err != nil {
		log.Println("Error occurred while adding user to the table", err)
		writer.WriteHeader(500)
		return
	}

	writer.WriteHeader(200)
	_, _ = writer.Write(data)
}

func updateUser(writer http.ResponseWriter, request *http.Request) {
	user := &model.User{}
	data, err := extractBody(request, user)
	if err != nil {
		writer.WriteHeader(500)
		return
	}
	userID, err, badRequest := parseUserID(writer, request)
	if badRequest {
		return
	}
	if err != nil {
		log.Println("Error occurred while parson body", err)
		writer.WriteHeader(500)
		return
	}

	count, err := dao.UpdateUser(userID, user)
	if err != nil {
		writer.WriteHeader(500)
		log.Println("Error occurred while updating user", err)
		return
	}
	if count == 0 {
		writer.WriteHeader(404)
		return
	}
	writer.WriteHeader(200)
	writer.Write(data)
}

func deleteUser(writer http.ResponseWriter, request *http.Request) {
	parsedUserID, err, badRequest := parseUserID(writer, request)
	if badRequest {
		return
	}
	if err != nil {
		log.Println("Error occurred while parsing user id", err)
		writer.WriteHeader(500)
		return
	}
	count, err := dao.DeleteUser(parsedUserID)
	if err != nil {
		log.Println("Error occurred while deleting user", err)
		writer.WriteHeader(500)
		return
	}

	if count == 0 {
		writer.WriteHeader(404)
		return
	}
	writer.WriteHeader(200)
}

func extractBody(request *http.Request, user *model.User) ([]byte, error) {
	body := request.Body
	defer body.Close()
	data, err := ioutil.ReadAll(body)
	if err != nil {
		log.Println("Error occurred during body de-serialization", err)
		return nil, err
	}

	if parseErr := json.Unmarshal(data, user); parseErr != nil {
		log.Println("Error occurred during json parsing", err)
		return nil, parseErr
	}
	return data, nil
}

func parseUserID(writer http.ResponseWriter, request *http.Request) (int, error, bool) {
	path := request.URL.Path
	pathParams := strings.Split(path, "/")
	userID := pathParams[len(pathParams)-1]
	if !numRegex.MatchString(userID) {
		writer.WriteHeader(400)
		log.Print("Invalid user id:", userID)
		return 0, nil, true
	}
	parsedUserID, err := strconv.Atoi(userID)
	return parsedUserID, err, false
}
