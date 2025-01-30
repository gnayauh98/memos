package resource

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kehuay/aimemos/store"
)

// 定义类型

type ResourceCreate struct {
	Data []byte    `json:"data"`
	Type string    `json:"type"`
	Id   uuid.UUID `json:"id,omitempty"`
}

// 定义函数

func NewResourceApi(group fiber.Router) fiber.Router {
	resourceApi := group.Group("/resource")

	// 新建
	resourceApi.Post("", createResource)
	resourceApi.Get("/:id", getResource)

	return resourceApi
}

func createResource(c *fiber.Ctx) error {

	resourceCreate := ResourceCreate{}

	resourceCreate.Type = c.FormValue("type")

	file, err := c.FormFile("files")

	if err != nil {
		return errors.New("文件内容不存在")
	}

	reader, err := file.Open()
	if err != nil {
		return errors.New("文件打开失败")
	}
	defer reader.Close()

	buff := new(bytes.Buffer)

	_, err = io.Copy(buff, reader)

	if err != nil {
		return errors.New("文件读取失败")
	}

	resourceCreate.Data = buff.Bytes()

	// log.Println(file.Size)
	// log.Println(n)

	resourceCreate.Id = uuid.New()

	_store := c.Locals("store").(*store.Store)

	_, err = store.CreateResource(store.Resource{
		Id:   resourceCreate.Id,
		Type: resourceCreate.Type,
		Data: resourceCreate.Data,
	}, *_store)

	// log.Println(resourceCreate)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"code": 1000,
		"data": fiber.Map{
			"id":  resourceCreate.Id,
			"url": fmt.Sprintf("/api/resource/%s", resourceCreate.Id),
		},
	})
}

func getResource(c *fiber.Ctx) error {
	id := c.Params("id")

	_store := c.Locals("store").(*store.Store)

	resource, err := store.QueryResource(id, *_store)

	// log.Println(resource)

	if err != nil {
		return err
	}

	c.Set("content-type", "image/png")

	return c.Send(resource.Data)
}
