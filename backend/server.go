package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kehuay/aimemos/api/auth"
	"github.com/kehuay/aimemos/api/memo"
	"github.com/kehuay/aimemos/api/user"
	"github.com/kehuay/aimemos/store"
)

func main() {

	_store, err := store.NewStore()

	if err != nil {
		log.Fatal("数据库配置失败")
		return
	}

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("store", &_store)
		return c.Next()
	})

	appApi := app.Group("/api")

	user.NewUserApi(appApi)
	auth.NewAuthApi(appApi)
	memo.NewMemoApi(appApi)

	app.Listen(":3000")
}
