package storage

import (
	"errors"
)

// ErrNotFound ошибка возвращаемая в случае если запрашиваемое значение не найдено в хранилище
var ErrNotFound = errors.New("key not found")

// URL интерфейс хранилища имеющий методы для сохранения,
// получения и проверки существования значения с поддержкой шардирования
type URL interface {
	Get(id uint64) (string, error)
	Put(url string) (uint64, error)
}
