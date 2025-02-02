package resource

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/anqzi/memos/api/memo"
	"github.com/anqzi/memos/store"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// 定义类型

type ResourceCreate struct {
	Data      []byte    `json:"data"`
	Type      string    `json:"type"`
	Id        uuid.UUID `json:"id,omitempty"`
	CreatorId int64     `json:"creator_id,omitempty"`
}

type ResourceQuery struct {
	CreatorId int64 `json:"creator_id,omitempty"`
}

type Resource struct {
	Url string `json:"url"`
	Id  string `json:"id"`
}

// 定义函数

func NewResourceApi(group fiber.Router) fiber.Router {
	resourceApi := group.Group("/resource")

	// 新建
	resourceApi.Post("", memo.GetUserInfo, createResource)
	resourceApi.Get("/:id", getResource)
	resourceApi.Post("/all", memo.GetUserInfo, getResources)

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

	user_id := c.Locals("user_id").(string)

	resourceCreate.CreatorId, err = strconv.ParseInt(user_id, 10, 64)

	if err != nil {
		return err
	}

	_, err = store.CreateResource(store.Resource{
		Id:        resourceCreate.Id,
		Type:      resourceCreate.Type,
		Data:      resourceCreate.Data,
		CreatorId: resourceCreate.CreatorId,
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

func getResources(c *fiber.Ctx) error {

	user_id := c.Locals("user_id").(string)
	_store := c.Locals("store").(*store.Store)

	_user_id, err := strconv.ParseInt(user_id, 10, 64)

	if err != nil {
		return err
	}

	rows, err := store.QueryResources(store.Resource{
		CreatorId: _user_id,
	}, *_store)

	if err != nil {
		return err
	}

	resource := make([]Resource, 0)

	for _, row := range rows {
		resource = append(resource, Resource{
			Url: fmt.Sprintf("/api/resource/%s", row.Id),
			Id:  row.Id.String(),
		})
	}

	return c.JSON(resource)

}
