package store

import (
	"fmt"
	"log"
)

type UserCreate struct {
	Name         string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type UserUpdate struct {
	Name         string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type UserQuery struct {
	Name         string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type User struct {
	Id           string
	Name         string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

func CreateUser(userCreate UserCreate, store Store) error {
	_, err := store.db.Exec(
		"insert into users (username, email, password_hash) values ($1, $2, $3)",
		userCreate.Name,
		userCreate.Email,
		userCreate.PasswordHash,
	)
	return err
}

func UpdateUser(userUpdate UserUpdate, user User, store Store) error {
	_, err := store.db.Exec("")
	return err
}

func QueryUser(userQuery UserQuery, store Store) (User, error) {
	// 构造查询表达式
	queryStr := ""

	count := 0

	if userQuery.Name != "" {
		queryStr = fmt.Sprintf("username='%s'", userQuery.Name)
		count += 1
	}
	if userQuery.Email != "" {
		if count > 0 {
			queryStr += " and "
		}
		queryStr += fmt.Sprintf("email='%s'", userQuery.Email)
		count += 1
	}
	if userQuery.PasswordHash != "" {
		if count > 0 {
			queryStr += " and "
		}
		queryStr += fmt.Sprintf("password_hash='%s'", userQuery.PasswordHash)
		count += 1
	}

	if count != 0 {
		queryStr = "select id, username, email, password_hash from users where " + queryStr + ";"
	} else {
		queryStr = "select id, username, email, password_hash from users;"
	}

	log.Println("查询字符串: ", queryStr)

	user := User{}

	err := store.db.QueryRow(queryStr).Scan(&user.Id, &user.Name, &user.Email, &user.PasswordHash)

	return user, err
}
