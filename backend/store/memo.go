package store

import "database/sql"

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
	MemoStatusDraft    MemoStatus = "draft"
	MemoStatusActive   MemoStatus = "active"
	MemoStatusArchived MemoStatus = "archived"
)

type MemoCreate struct {
	Content   string
	CreatorId string
}

type MemoQuery struct {
	CreatorId string
	PageSize  int64
	PageNo    int64
}

func CreateMemo(memoCreate MemoCreate, store Store) (Memo, error) {

	memo := Memo{
		Content:   memoCreate.Content,
		CreatorId: memoCreate.CreatorId,
	}

	err := store.db.QueryRow(
		"insert into memos (content, creator_id) values ($1, $2) returning id, create_at, update_at;",
		memoCreate.Content,
		memoCreate.CreatorId,
	).Scan(&memo.Id, &memo.CreateAt, &memo.UpdateAt)

	if err != nil {
		return Memo{}, err
	}

	return memo, err
}

func QueryMemos(memoQuery MemoQuery, store Store) ([]Memo, error) {

	var rows *sql.Rows
	var err error

	if memoQuery.PageNo != 0 && memoQuery.PageSize != 0 {
		rows, err = store.db.Query(
			"select memos.*, users.username from memos left join users on memos.creator_id = users.id where creator_id=$1 order by create_at desc limit $2 offset $3",
			memoQuery.CreatorId,
			memoQuery.PageSize,
			(memoQuery.PageNo-1)*memoQuery.PageSize,
		)
	} else {
		rows, err = store.db.Query(
			"select memos.*, users.username from memos left join users on memos.creator_id = users.id where creator_id=$1 order by create_at desc",
			memoQuery.CreatorId,
		)
	}

	if err != nil {
		return []Memo{}, err
	}

	memos := make([]Memo, 0)

	for rows.Next() {
		memo := Memo{}
		err = rows.Scan(&memo.Id,
			&memo.Content,
			&memo.CreateAt,
			&memo.UpdateAt,
			&memo.CreatorId,
			&memo.IsFixed,
			&memo.Status,
			&memo.CreatorName)
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
		"select * from memos where id=$1 and creator_id=$2",
		id,
		creator_id,
	).Scan(&memo.Id, &memo.Content, &memo.CreateAt, &memo.UpdateAt, &memo.CreatorId, &memo.IsFixed, &memo.Status)

	return memo, err
}
