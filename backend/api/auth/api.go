package auth

import (
	"crypto/md5"
	"errors"
	"fmt"
	"log"

	userApi "github.com/anqzi/memos/api/user"
	"github.com/anqzi/memos/store"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func NewAuthApi(group fiber.Router) fiber.Router {
	authApi := group.Group("/user")

	authApi.Post("/signin", SignIn)
	authApi.Post("/verify", IsValid)
	authApi.Get("/ping", ping)

	return authApi
}

func ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"ok": true,
	})
}

type SignInResult struct {
	User  CustomClaimsInfo `json:"user"`
	Token string           `json:"token"`
}

type CustomClaimsInfo struct {
	UserId    string `json:"user_id"`
	UserName  string `json:"username"`
	UserEmail string `json:"email"`
}

type CustomClaims struct {
	jwt.RegisteredClaims
	Info CustomClaimsInfo `json:"info"`
}

type ParserHeader struct {
	Authorization string `reqHeader:"Authorization"`
}

// 登录
func SignIn(c *fiber.Ctx) error {

	// 解析查询差数
	userQuery := userApi.UserQuery{}
	err := c.BodyParser(&userQuery)
	if err != nil {
		return err
	}

	// 根据邮箱地址查询用户
	_store, _ := c.Locals("store").(*store.Store)
	user, err := store.QueryUser(store.UserQuery{
		Email: userQuery.Email,
	}, *_store)

	if err != nil {
		return err
	}

	log.Println(user)

	passwordHash := fmt.Sprintf("%x", md5.Sum([]byte(userQuery.Password)))
	if passwordHash != user.PasswordHash {
		log.Println("密码错误", passwordHash, user.PasswordHash)
		return errors.New("密码错误")
	}

	// TODO 生成登录令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		Info: CustomClaimsInfo{
			UserName:  user.Name,
			UserEmail: user.Email,
			UserId:    user.Id,
		},
	})

	tokenStr, err := token.SignedString([]byte("aimemos-2025"))

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return c.JSON(SignInResult{
		User: CustomClaimsInfo{
			UserName:  user.Name,
			UserEmail: user.Email,
			UserId:    user.Id,
		},
		Token: tokenStr,
	})
}

func IsValid(c *fiber.Ctx) error {
	parserHeader := ParserHeader{}

	err := c.ReqHeaderParser(&parserHeader)
	if err != nil {
		return err
	}

	if len(parserHeader.Authorization) < 6 {
		return errors.New("未鉴权")
	}

	token, err := jwt.ParseWithClaims(parserHeader.Authorization[7:], &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("aimemos-2025"), nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("验证失败")
	}

	claims, ok := token.Claims.(*CustomClaims)

	if !ok {
		return errors.New("解析失败")
	}

	return c.JSON(claims.Info)

}
