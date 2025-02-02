package store

import (
	"github.com/google/uuid"
)

type Resource struct {
	Id        uuid.UUID `json:"id"`
	Type      string    `json:"type"`
	Data      []byte    `json:"data"`
	CreatorId int64     `json:"creator_id"`
}

func CreateResource(resource Resource, store Store) (Resource, error) {

	err := store.db.QueryRow(
		"insert into resources (id, type, raw,creator_id) values ($1, $2, $3, $4) returning id, type, raw",
		resource.Id,
		resource.Type,
		resource.Data,
		resource.CreatorId,
	).Scan(&resource.Id, &resource.Type, &resource.Data)

	return resource, err
}

func QueryResource(id string, store Store) (Resource, error) {
	resource := Resource{}

	err := store.db.QueryRow(
		"select id,type,raw from resources where id = $1;",
		id,
	).Scan(&resource.Id, &resource.Type, &resource.Data)

	if err != nil {
		return resource, err
	}

	return resource, nil
}

func QueryResources(resource Resource, store Store) ([]Resource, error) {
	rows, err := store.db.Query(
		"select id,creator_id,type from resources where creator_id=$1 order by created_at desc;",
		resource.CreatorId,
	)

	if err != nil {
		return []Resource{}, err
	}

	rowList := make([]Resource, 0)

	for rows.Next() {
		resource := Resource{}
		err = rows.Scan(&resource.Id, &resource.CreatorId, &resource.Type)
		if err != nil {
			return rowList, err
		}
		rowList = append(rowList, resource)
	}

	return rowList, nil
}
