package storage

import (
	"context"
	"errors"
	restservice "task/rest_service"
)

var ErrorTheSameId = errors.New("cannot use the same uuid for 'id' and 'parent_id' fields")
var ErrorProjectId = errors.New("not valid 'project_id'")
var ErrorNotFound = errors.New("not found")

type StorageI interface {
	CloseDB()
	Phone() PhoneRepoI
}

type PhoneRepoI interface {
	Create(ctx context.Context, entity restservice.Phone) (id string, err error)
	GetByPK(ctx context.Context, id string) (phone restservice.Phone, err error)
	GetList(ctx context.Context, limit, offset int) (res []restservice.Phone, count int32, err error)
	Update(ctx context.Context, entity restservice.Phone) (rowsAffected int64, err error)
	Delete(ctx context.Context, id string) (rowsAffected int64, err error)
}
