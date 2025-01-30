package store

import (
	"github.com/google/uuid"
)

type Resource struct {
	Id   uuid.UUID `json:"id"`
	Type string    `json:"type"`
	Data []byte    `json:"data"`
}

func CreateResource(resource Resource, store Store) (Resource, error) {

	err := store.db.QueryRow(
		"insert into resources (id, type, data) values ($1, $2, $3) returning id, type, data",
		resource.Id,
		resource.Type,
		resource.Data,
	).Scan(&resource.Id, &resource.Type, &resource.Data)

	return resource, err
}

func QueryResource(id string, store Store) (Resource, error) {
	resource := Resource{}

	err := store.db.QueryRow(
		"select id,type,data from resources where id = $1;",
		id,
	).Scan(&resource.Id, &resource.Type, &resource.Data)

	if err != nil {
		return resource, err
	}

	return resource, nil
}
