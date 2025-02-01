package memo

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kehuay/aimemos/api/auth"
	"github.com/kehuay/aimemos/store"
	"github.com/kehuay/mark-parser/parser"
	"github.com/kehuay/mark-parser/parser/token"
	"github.com/kehuay/mark-parser/render"
)

func NewMemoApi(group fiber.Router) fiber.Router {
	memoApi := group.Group("/memo", GetUserInfo)

	memoApi.Put("", CreateMemo)
	memoApi.Patch("", UpdateMemo)
	memoApi.Post("/all", QueryMemos)
	memoApi.Post("/tags", QueryTags)
	memoApi.Post("/:id", QueryMemoById)

	return memoApi
}

type MemoCreate struct {
	Content string `json:"content"`
}

type MemoUpdate struct {
	Content string `json:"content"`
	Id      string `json:"id"`
}

type MemoQuery struct {
	PageNo   int64 `json:"pageNo,omitempty"`
	PageSize int64 `json:"pageSize,omitempty"`
}

func GetUserInfo(c *fiber.Ctx) error {

	parserHeader := auth.ParserHeader{}

	err := c.ReqHeaderParser(&parserHeader)
	if err != nil {
		return err
	}
	if len(parserHeader.Authorization) < 6 {
		return errors.New("未鉴权")
	}
	// log.Println(parserHeader.Authorization[7:])

	token, err := jwt.ParseWithClaims(parserHeader.Authorization[7:], &auth.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("aimemos-2025"), nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("验证失败")
	}

	claims, ok := token.Claims.(*auth.CustomClaims)

	if !ok {
		return errors.New("解析失败")
	}

	log.Println("用户信息: ", claims.Info)

	c.Locals("user_id", claims.Info.UserId)
	c.Locals("user_name", claims.Info.UserName)

	return c.Next()
}

func CreateMemo(c *fiber.Ctx) error {

	_store, _ := c.Locals("store").(*store.Store)
	userId, _ := c.Locals("user_id").(string)
	userName, _ := c.Locals("user_name").(string)

	// 获取memo的内容
	memoCreate := MemoCreate{}
	err := c.BodyParser(&memoCreate)

	if err != nil {
		return err
	}

	if len(memoCreate.Content) == 0 {
		return errors.New("内容为空")
	}

	// log.Println(memoCreate.Content)

	text := []byte(memoCreate.Content)
	tokens := parser.Parser(text)

	memo, err := store.CreateMemo(store.MemoCreate{
		Content:   memoCreate.Content,
		CreatorId: userId,
		Tags:      token.GetTags(text, tokens),
	}, *_store)

	if err != nil {
		return err
	}

	memo.CreatorName = userName

	memo.Content = render.RenderToHtml(text, tokens)

	return c.JSON(memo)
}

func UpdateMemo(c *fiber.Ctx) error {

	_store, _ := c.Locals("store").(*store.Store)
	userId, _ := c.Locals("user_id").(string)
	userName, _ := c.Locals("user_name").(string)

	// 获取memo的内容
	memoUpdate := MemoUpdate{}
	err := c.BodyParser(&memoUpdate)

	if err != nil {
		return err
	}

	if len(memoUpdate.Content) == 0 {
		return errors.New("内容为空")
	}

	text := []byte(memoUpdate.Content)
	tokens := parser.Parser(text)

	memo, err := store.UpdateMemo(store.MemoUpdate{
		Content:   memoUpdate.Content,
		Id:        memoUpdate.Id,
		CreatorId: userId,
		Tags:      token.GetTags(text, tokens),
	}, *_store)

	if err != nil {
		return err
	}

	memo.CreatorName = userName

	memo.Content = render.RenderToHtml(text, tokens)

	return c.JSON(memo)
}

func QueryMemos(c *fiber.Ctx) error {
	_store, _ := c.Locals("store").(*store.Store)
	userId, _ := c.Locals("user_id").(string)

	memoQuery := MemoQuery{}
	c.BodyParser(&memoQuery)

	memos, err := store.QueryMemos(store.MemoQuery{
		CreatorId: userId,
		PageNo:    memoQuery.PageNo,
		PageSize:  memoQuery.PageSize,
	}, *_store)

	if err != nil {
		return err
	}

	// markdown转换为html
	for index := range memos {
		text := []byte(memos[index].Content)
		tokens := parser.Parser(text)
		memos[index].Content = render.RenderToHtml(text, tokens)
	}

	return c.JSON(memos)
}

func QueryMemoById(c *fiber.Ctx) error {

	id := c.Params("id")
	_store, _ := c.Locals("store").(*store.Store)
	userId, _ := c.Locals("user_id").(string)

	if len(id) == 0 {
		return errors.New("请上传ID")
	}

	memo, err := store.QueryMemoById(id, userId, *_store)

	if err != nil {
		return err
	}

	return c.JSON(memo)
}

func QueryTags(c *fiber.Ctx) error {
	user_id, _ := c.Locals("user_id").(string)
	_store, _ := c.Locals("store").(*store.Store)

	tags, err := store.QueryTags(store.TagQuery{
		CreatorId: user_id,
	}, *_store)

	if err != nil {
		return err
	}

	return c.JSON(tags)
}
