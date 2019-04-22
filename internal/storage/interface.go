package storage

import (
	"errors"
)

// ErrNotFound ошибка возвращаемая в случае если запрашиваемое значение не найдено в хранилище
var ErrNotFound = errors.New("key not found")

// ShardedStorage интерфейс хранилища имеющий методы для сохранения,
// получения и проверки существования значения с поддержкой шардирования
type ShardedStorage interface {
	GetByUrl(url string, shard int) (string, error)
	Get(id uint64, shard int) (string, error)
	Put(url string, shard int) (uint64, error)
}
