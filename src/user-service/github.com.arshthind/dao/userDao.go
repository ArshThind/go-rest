package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"user-service/github.com.arshthind/dao/util"
	"user-service/github.com.arshthind/model"
)

var (
	ds  *sql.DB
	err error
)

func init() {
	ds, err = sql.Open("mysql", "user:password@tcp(hostname:port)/db")
	if err != nil {
		panic("DB initialization failed")
	}
	if err = ds.Ping(); err != nil {
		panic("DB ping failed. Please check your connection details")
	}
	log.Println("DB initialization successful")
}

func GetUsers() ([]model.User, error) {
	var users = make([]model.User, 0)
	rs, err := ds.Query(util.GetQuery(util.GET_USERS))
	if err != nil {
		log.Println("Error occurred while retrieving users from the db", err)
		return users, err
	}
	defer rs.Close()
	for rs.Next() {
		var userID, phoneNum int
		var userName, email string
		var userType string
		if err = rs.Scan(&userID, &userName, &email, &phoneNum, &userType); err != nil {
			log.Println("Error occurred during result set mapping", err)
			return users, err
		}
		users = append(users, model.NewUser(userID, userName, email, phoneNum, model.GetUserType(userType)))
	}
	return users, nil
}

func AddUser(user *model.User) (int, error) {
	query := util.GetQuery(util.ADD_USER)
	res, err := ds.Exec(query, user.Name, user.Email, user.Phone, string(user.UserType))
	id, _ := res.LastInsertId()
	return int(id), err
}

func DeleteUser(id int) (int, error) {
	query := util.GetQuery(util.DELETE_USER)
	res, err := ds.Exec(query, id)
	count, _ := res.RowsAffected()
	return int(count), err
}

func UpdateUser(id int, user *model.User) (int, error) {
	query := util.GetQuery(util.UPDATE_USER)
	rs, err := ds.Exec(query, user.Name, user.Email, user.Phone, user.UserType, id)
	count, _ := rs.RowsAffected()
	return int(count), err
}
