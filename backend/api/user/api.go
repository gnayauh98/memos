package user

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kehuay/aimemos/store"
)

// 定义类型

type UserCreate struct {
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserQuery struct {
	Name     string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type User struct{}

// 定义函数

func NewUserApi(group fiber.Router) fiber.Router {
	userApi := group.Group("/user")

	// 新建
	userApi.Put("", registryUser)
	// 查询
	userApi.Post("", queryUser)

	return userApi
}

func registryUser(c *fiber.Ctx) error {

	bodyBytes := c.Request().Body()

	userCreate := UserCreate{}

	err := json.Unmarshal(bodyBytes, &userCreate)

	if err != nil {
		return err
	}

	_store, _ := c.Locals("store").(*store.Store)

	err = store.CreateUser(store.UserCreate{
		Name:         userCreate.Name,
		Email:        userCreate.Email,
		PasswordHash: fmt.Sprintf("%x", md5.Sum([]byte(userCreate.Password))),
	}, *_store)

	if err != nil {
		return c.JSON(fiber.Map{
			"code":    "1001",
			"message": "注册失败",
			"err":     err.Error(),
		})
	}

	userCreate.Password = ""

	return c.JSON(fiber.Map{
		"code":    "1000",
		"message": "注册成功",
		"data": fiber.Map{
			"user": userCreate,
		},
	})
}

func queryUser(c *fiber.Ctx) error {

	// 处理application/json数据
	bodyBytes := c.Request().Body()

	userQuery := UserQuery{}

	err := json.Unmarshal(bodyBytes, &userQuery)

	if err != nil {
		return err
	}

	log.Println(userQuery)

	_store, _ := c.Locals("store").(*store.Store)

	userQueryStore := store.UserQuery{
		Email: userQuery.Email,
		Name:  userQuery.Name,
	}

	if userQuery.Password != "" {
		passwordHash := md5.Sum([]byte(userQuery.Password))
		userQueryStore.PasswordHash = fmt.Sprintf("%x", passwordHash)
	}

	user, err := store.QueryUser(userQueryStore, *_store)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"code":    "1000",
		"message": "查询成功",
		"data": fiber.Map{
			"user": user,
		},
	})
}
