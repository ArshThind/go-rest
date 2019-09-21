package model

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
	UserType `json:"userType"`
}

type UserType string

const (
	CUSTOMER UserType = "CUSTOMER"
	VENDOR   UserType = "VENDOR"
)

func NewUser(id int, name, email string, phone int, userType UserType) User {
	return User{id, name, email, phone, userType}
}

func GetUserType(userType string) UserType {
	if userType == "C" {
		return CUSTOMER
	} else if userType == "V" {
		return VENDOR
	}
	return ""
}
