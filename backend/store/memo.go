package store

import (
	"database/sql"
	"fmt"
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
	Tags      []string
}

type Tag struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatorId string `json:"creator_id"`
}

type TagMemo struct {
	TagId     string `json:"tag_id"`
	TagName   string `json:"tag_name"`
	CreatorId string `json:"creator_id"`
	MemoId    string `json:"memo_id"`
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

	err = tx.QueryRow(
		"insert into memos (markdown, creator_id) values ($1, $2) returning id, created_at, updated_at;",
		memoCreate.Content,
		memoCreate.CreatorId,
	).Scan(&memo.Id, &memo.CreateAt, &memo.UpdateAt)

	if err != nil {
		return Memo{}, err
	}

	if len(memoCreate.Tags) > 0 {
		valuesStr := make([]any, 0, len(memoCreate.Tags)*2)
		placeholderStr := make([]string, len(memoCreate.Tags))
		for index, tag := range memoCreate.Tags {
			valuesStr = append(valuesStr, tag, memoCreate.CreatorId)
			placeholderStr[index] = fmt.Sprintf("($%d,$%d)", 2*index+1, 2*index+2)
		}
		queryStmt := fmt.Sprintf(`insert into tags (name, creator_id) values %s on conflict (name, creator_id) do update set name=excluded.name, creator_id=excluded.creator_id returning id, name;`, strings.Join(placeholderStr, ","))

		rows, err := tx.Query(queryStmt, valuesStr...)

		if err != nil {
			return Memo{}, err
		}

		defer rows.Close()

		memoTagPairStr := make([]any, 0, len(memoCreate.Tags)*2)

		for rows.Next() {

			tagMemo := TagMemo{
				MemoId: memo.Id,
				// CreatorId: memo.CreatorId,
			}

			err = rows.Scan(&tagMemo.TagId, &tagMemo.TagName)

			if err != nil {
				return Memo{}, err
			}

			memoTagPairStr = append(memoTagPairStr, tagMemo.TagId, tagMemo.MemoId)
		}

		// 插入memo-tag
		queryStmt = fmt.Sprintf(`insert into memo_tag (tag_id, memo_id) values %s on conflict (tag_id, memo_id) do update set tag_id=excluded.tag_id, memo_id=excluded.memo_id;`, strings.Join(placeholderStr, ","))

		_, err = tx.Exec(queryStmt, memoTagPairStr...)

		if err != nil {
			return Memo{}, err
		}
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
	var valuesStr []any
	var placeholderStr []string
	if len(memoUpdate.Tags) > 0 {
		valuesStr = make([]any, 0, len(memoUpdate.Tags)*2)
		placeholderStr = make([]string, len(memoUpdate.Tags))
		for index, tag := range memoUpdate.Tags {
			valuesStr = append(valuesStr, tag, memoUpdate.CreatorId)
			placeholderStr[index] = fmt.Sprintf("($%d,$%d)", 2*index+1, 2*index+2)
		}

	}

	// 插入的新tag与老tag
	// tagMemos := make([]TagMemo, 0, len(memoUpdate.Tags))
	tagsMap := make(map[string]bool, len(memoUpdate.Tags))

	if len(memoUpdate.Tags) > 0 {
		// 插入tags
		queryStmt := fmt.Sprintf(`insert into tags (name, creator_id) values %s on conflict (name, creator_id) do update set name=excluded.name, creator_id=excluded.creator_id returning id, name, creator_id;`, strings.Join(placeholderStr, ","))

		rows, err := tx.Query(queryStmt, valuesStr...)
		if err != nil {
			return Memo{}, err
		}
		defer rows.Close()

		memoTagPairStr := make([]any, 0, len(memoUpdate.Tags)*2)

		for rows.Next() {

			tagMemo := TagMemo{
				MemoId: memoUpdate.Id,
			}

			err = rows.Scan(&tagMemo.TagId, &tagMemo.TagName, &tagMemo.CreatorId)

			if err != nil {
				return Memo{}, err
			}

			tagsMap[tagMemo.TagId] = true
			memoTagPairStr = append(memoTagPairStr, tagMemo.TagId, tagMemo.MemoId)

			// tagMemos = append(tagMemos, tagMemo)
		}

		// 插入memo-tag
		queryStmt = fmt.Sprintf(`insert into memo_tag (tag_id, memo_id) values %s on conflict (tag_id, memo_id) do update set tag_id=excluded.tag_id, memo_id=excluded.memo_id;`, strings.Join(placeholderStr, ","))

		_, err = tx.Exec(queryStmt, memoTagPairStr...)

		if err != nil {
			return Memo{}, err
		}
	}

	// log.Println(tagsMap)

	// 查询当前存在的tag-memo对
	rows, err := tx.Query("select tag_id, memo_id from memo_tag where memo_id=$1", memoUpdate.Id)
	if err != nil {
		return Memo{}, err
	}
	defer rows.Close()

	deleteTagMemoIds := make([]any, 0)
	deleteTagMemoPlaceholder := make([]string, 0)
	countIndex := 2
	deleteTagMemoIds = append(deleteTagMemoIds, memoUpdate.Id)
	for rows.Next() {
		tagMemo := TagMemo{}
		err = rows.Scan(&tagMemo.TagId, &tagMemo.MemoId)
		if err != nil {
			return Memo{}, err
		}
		if !tagsMap[tagMemo.TagId] {
			deleteTagMemoPlaceholder = append(
				deleteTagMemoPlaceholder,
				fmt.Sprintf("$%d", countIndex),
			)
			countIndex += 1
			deleteTagMemoIds = append(deleteTagMemoIds, tagMemo.TagId)

		}
	}

	if len(deleteTagMemoIds) > 1 {
		queryStmt := fmt.Sprintf("delete from memo_tag where memo_id=$1 and tag_id in (%s);", strings.Join(deleteTagMemoPlaceholder, ","))

		// log.Println(queryStmt)
		_, err = tx.Exec(
			queryStmt,
			deleteTagMemoIds...,
		)

		if err != nil {
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

	tags := ""
	queryResultFiled := "memos.id, memos.markdown, memos.created_at, memos.updated_at, memos.creator_id, memos.status, users.username"
	queryTags := ""
	queryWhere := "creator_id=$1"
	if len(memoQuery.Tags) > 0 {
		tags = fmt.Sprintf("memo_tag.tag_id in (%s)", strings.Join(memoQuery.Tags, ","))
		queryWhere += fmt.Sprintf(" and %s", tags)
		queryTags = "join memo_tag on memos.id = memo_tag.memo_id"
	}

	if memoQuery.PageNo == 0 {
		memoQuery.PageNo = 1
	}
	if memoQuery.PageSize == 0 {
		memoQuery.PageSize = 10
	}

	rows, err = store.db.Query(
		fmt.Sprintf(
			"select %s from memos %s join users on memos.creator_id = users.id where %s group by memos.id, users.username order by created_at desc limit $2 offset $3",
			queryResultFiled,
			queryTags,
			queryWhere,
		),
		memoQuery.CreatorId,
		memoQuery.PageSize,
		(memoQuery.PageNo-1)*memoQuery.PageSize,
	)

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
