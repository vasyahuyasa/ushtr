package storage

import (
	"errors"
)

// ErrNotFound ошибка возвращаемая в случае если запрашиваемое значение не найдено в хранилище
var ErrNotFound = errors.New("url not found")

// Saver позволяет сохранять ссылку в хранилище и в замен получаьть уникальный идентификатор
type Saver interface {
	Save(url string) (uint64, error)
}

// Getter позволяет по уникальному идентификатору получить сохранненые данные.
// Если запиь не найдна, то будет возыращена ошибка ErrNotFound
type Getter interface {
	Get(id uint64) (string, error)
}

type GetterSaver interface {
	Getter
	Saver
}
