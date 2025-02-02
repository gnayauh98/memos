package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kehuay/aimemos/api/auth"
	"github.com/kehuay/aimemos/api/memo"
	"github.com/kehuay/aimemos/api/resource"
	"github.com/kehuay/aimemos/api/user"
	"github.com/kehuay/aimemos/store"
	"github.com/kehuay/aimemos/store/db"
)

func main() {

	dbConfig := db.NewDefaultDBConfig()

	var host = flag.String("host", dbConfig.Host, "db host")
	var port = flag.String("port", dbConfig.Port, "db port")
	var dbuser = flag.String("user", dbConfig.User, "db user")
	var dbname = flag.String("dbname", dbConfig.DBName, "db database")
	var dbpassword = flag.String("password", dbConfig.Password, "db password")

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
