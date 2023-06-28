package helper

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/afiifatuts/bankmnc/model"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(pw, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	return err == nil
}

func IsRegistered(username string) bool {
	jsonData, err := os.Open("jsonData/user.json")

	if err != nil {
		return false
	}

	defer jsonData.Close()

	byteValue, _ := ioutil.ReadAll(jsonData)

	var users model.Users
	var status bool

	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].Username == username {
			status = true
			break
		} else {
			status = false
		}
	}
	return status

}

func GetPassword(username string) string {
	jsonData, err := os.Open("jsonData/user.json")
	if err != nil {
		return "failed open file"
	}

	defer jsonData.Close()
	//read opened jsonData
	byteValue, _ := ioutil.ReadAll(jsonData)

	var users model.Users
	var password string

	//unmarshal byteArray and get its pw
	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].Username == username {
			password = users.Users[i].Password
		} else {
			password = ""
		}
	}
	return password

}

func GetUser(username string) model.User {
	jsonData, err := os.Open("jsonData/user.json")
	if err != nil {
		return model.User{}
	}

	defer jsonData.Close()
	//read opened jsonData
	byteValue, _ := ioutil.ReadAll(jsonData)

	var users model.Users

	//unmarshal byteArray and get its pw
	json.Unmarshal(byteValue, &users)

	req := model.User{}

	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].Username == username {
			req.ID = users.Users[i].ID
			req.Username = users.Users[i].Username
			req.Password = users.Users[i].Password
			req.IsLogin = users.Users[i].IsLogin

		}
	}
	return req

}
