package store

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type Memo struct {
	Id          string     `json:"id"`
	Content     string     `json:"content"`
	CreateAt    string     `json:"createAt"`
	UpdateAt    string     `json:"updateAt"`
	CreatorId   string     `json:"creatorId"`
	IsFixed     bool       `json:"isFixed"`
	Status      MemoStatus `json:"status"`
	CreatorName string     `json:"username"`
}

type MemoStatus string

const (
	MemoStatusDraft    MemoStatus = "private"
	MemoStatusActive   MemoStatus = "public"
	MemoStatusArchived MemoStatus = "archived"
)

type MemoCreate struct {
	Content   string
	CreatorId string
	Tags      []string
}

type MemoUpdate struct {
	Content   string
	Id        string
	CreatorId string
	Tags      []string
}

type MemoQuery struct {
	CreatorId string
	PageSize  int64
	PageNo    int64
}

type Tag struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TagQuery struct {
	CreatorId string `json:"creator_id"`
}

func CreateMemo(memoCreate MemoCreate, store Store) (Memo, error) {

	memo := Memo{
		Content:   memoCreate.Content,
		CreatorId: memoCreate.CreatorId,
	}

	tx, err := store.db.Begin()
	if err != nil {
		return Memo{}, err
	}
	defer tx.Rollback()
	// 查询已经存在的tags
	rows, err := tx.Query("select name from tags where creator_id=$1", memoCreate.CreatorId)
	if err != nil {
		return Memo{}, err
	}
	defer rows.Close()

	tagsMap := make(map[string]bool, len(memoCreate.Tags))
	for _, tag := range memoCreate.Tags {
		tagsMap[tag] = false
	}
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return Memo{}, err
		}
		tagsMap[name] = true
	}
	// 过滤不存在标签
	var tags []any

	for _, tag := range memoCreate.Tags {
		if !tagsMap[tag] {
			tags = append(tags, tag)
		}
	}
	if len(tags) > 0 {
		valuesStr := make([]any, 0, len(tags)*2)
		placeholderStr := make([]string, len(tags))
		for index, tag := range tags {
			valuesStr = append(valuesStr, tag, memoCreate.CreatorId)
			placeholderStr[index] = fmt.Sprintf("($%d,$%d)", 2*index+1, 2*index+2)
		}
		queryStmt := fmt.Sprintf(`insert into tags (name, creator_id) values %s`, strings.Join(placeholderStr, ","))
		log.Printf("%s", queryStmt)
		log.Println(valuesStr...)

		if _, err = tx.Exec(queryStmt, valuesStr...); err != nil {
			return Memo{}, err
		}
	}

	err = tx.QueryRow(
		"insert into memos (markdown, creator_id) values ($1, $2) returning id, created_at, updated_at;",
		memoCreate.Content,
		memoCreate.CreatorId,
	).Scan(&memo.Id, &memo.CreateAt, &memo.UpdateAt)

	if err != nil {
		return Memo{}, err
	}

	if err = tx.Commit(); err != nil {
		return Memo{}, err
	}

	return memo, err
}

func UpdateMemo(memoUpdate MemoUpdate, store Store) (Memo, error) {

	memo := Memo{
		Content:   memoUpdate.Content,
		CreatorId: memoUpdate.CreatorId,
		Id:        memoUpdate.Id,
	}

	tx, err := store.db.Begin()
	if err != nil {
		return Memo{}, err
	}
	defer tx.Rollback()
	// 查询已经存在的tags
	rows, err := tx.Query("select name from tags where creator_id=$1", memoUpdate.CreatorId)
	if err != nil {
		return Memo{}, err
	}
	defer rows.Close()

	tagsMap := make(map[string]bool, len(memoUpdate.Tags))
	for _, tag := range memoUpdate.Tags {
		tagsMap[tag] = false
	}
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return Memo{}, err
		}
		tagsMap[name] = true
	}
	// 过滤不存在标签
	var tags []any

	for _, tag := range memoUpdate.Tags {
		if !tagsMap[tag] {
			tags = append(tags, tag)
		}
	}
	if len(tags) > 0 {
		valuesStr := make([]any, 0, len(tags)*2)
		placeholderStr := make([]string, len(tags))
		for index, tag := range tags {
			valuesStr = append(valuesStr, tag, memoUpdate.CreatorId)
			placeholderStr[index] = fmt.Sprintf("($%d,$%d)", 2*index+1, 2*index+2)
		}
		queryStmt := fmt.Sprintf(`insert into tags (name, creator_id) values %s`, strings.Join(placeholderStr, ","))
		log.Printf("%s", queryStmt)
		log.Println(valuesStr...)

		if _, err = tx.Exec(queryStmt, valuesStr...); err != nil {
			return Memo{}, err
		}
	}

	err = tx.QueryRow(
		"update memos set markdown=$1 where id=$2 returning id, created_at, updated_at;",
		memoUpdate.Content,
		memoUpdate.Id,
	).Scan(&memo.Id, &memo.CreateAt, &memo.UpdateAt)

	if err != nil {
		return Memo{}, err
	}

	if err = tx.Commit(); err != nil {
		return Memo{}, err
	}

	return memo, err
}

func QueryMemos(memoQuery MemoQuery, store Store) ([]Memo, error) {

	var rows *sql.Rows
	var err error

	if memoQuery.PageNo != 0 && memoQuery.PageSize != 0 {
		rows, err = store.db.Query(
			"select memos.id, memos.markdown, memos.created_at, memos.updated_at, memos.creator_id, memos.status, users.username from memos left join users on memos.creator_id = users.id where creator_id=$1 order by created_at desc limit $2 offset $3",
			memoQuery.CreatorId,
			memoQuery.PageSize,
			(memoQuery.PageNo-1)*memoQuery.PageSize,
		)
	} else {
		rows, err = store.db.Query(
			"select memos.id, memos.markdown, memos.created_at, memos.updated_at, memos.creator_id, memos.status, users.username from memos left join users on memos.creator_id = users.id where creator_id=$1 order by created_at desc",
			memoQuery.CreatorId,
		)
	}

	if err != nil {
		return []Memo{}, err
	}

	defer rows.Close()

	memos := make([]Memo, 0)

	for rows.Next() {
		memo := Memo{}
		err = rows.Scan(
			&memo.Id,
			&memo.Content,
			&memo.CreateAt,
			&memo.UpdateAt,
			&memo.CreatorId,
			&memo.Status,
			&memo.CreatorName,
		)
		if err != nil {
			return []Memo{}, err
		}
		memos = append(memos, memo)
	}

	return memos, nil
}

func QueryMemoById(id string, creator_id string, store Store) (Memo, error) {

	memo := Memo{}

	err := store.db.QueryRow(
		"select memos.id, memos.markdown, memos.created_at, memos.updated_at, memos.creator_id, memos.status from memos where id=$1 and creator_id=$2",
		id,
		creator_id,
	).Scan(&memo.Id, &memo.Content, &memo.CreateAt, &memo.UpdateAt, &memo.CreatorId, &memo.Status)

	return memo, err
}

func QueryTags(tagQuery TagQuery, store Store) ([]Tag, error) {
	rows, err := store.db.Query(`select id, name from tags where creator_id=$1;`, tagQuery.CreatorId)

	if err != nil {
		return []Tag{}, err
	}

	defer rows.Close()

	tags := make([]Tag, 0)
	for rows.Next() {
		var tag Tag
		if err = rows.Scan(&tag.Id, &tag.Name); err != nil {
			return []Tag{}, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
