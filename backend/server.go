package main

import (
	"flag"
	"log"

	"github.com/anqzi/memos/api/auth"
	"github.com/anqzi/memos/api/memo"
	"github.com/anqzi/memos/api/resource"
	"github.com/anqzi/memos/api/user"
	"github.com/anqzi/memos/store"
	"github.com/anqzi/memos/store/db"
	"github.com/gofiber/fiber/v2"
)

func main() {

	dbConfig := db.NewDefaultDBConfig()

	var host = flag.String("host", dbConfig.Host, "db host")
	var port = flag.String("port", dbConfig.Port, "db port")
	var dbuser = flag.String("user", dbConfig.User, "db user")
	var dbname = flag.String("db", dbConfig.DBName, "db database")
	var dbpassword = flag.String("pass", dbConfig.Password, "db password")

	flag.Parse()

	dbConfig.Host = *host
	dbConfig.Port = *port
	dbConfig.User = *dbuser
	dbConfig.DBName = *dbname
	dbConfig.Password = *dbpassword

	_store, err := store.NewStore(dbConfig)

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
	resource.NewResourceApi(appApi)

	app.Listen(":3000")
}
