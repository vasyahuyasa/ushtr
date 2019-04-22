package simple

import (
	"errors"
	"fmt"
)

const (
	bitsInChar = 6
	maxChar    = 63 // 0b00111111
	maxLen     = 11

	defaultChars = "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz0123456789-_"
)

type SimpleShortener struct {
	chars *charmap
}

func New() (*SimpleShortener, error) {
	chars, err := charmapFromString(defaultChars)
	if err != nil {
		return nil, errors.New("can not create charmap: " + err.Error())
	}

	return &SimpleShortener{
		chars: chars,
	}, nil
}

// Encode кодирует id через алфавит состоящий из 64 символов
// id должна быть больше 0 иначе, на выходе будет пустая строка
func (s *SimpleShortener) Encode(id uint64) string {
	// uint 64 можно представить как 11 групп по 6 бит
	// в 6 бит можно закодировать 64 символа что должно совпадать с длиной алфавита
	// заполним слайс значениями кодированных символов
	values := make([]byte, maxLen)
	for i := maxLen - 1; i >= 0; i-- {
		b := byte(id & maxChar) // id & 0b00111111
		values[i] = b
		id >>= 6
	}

	// откинем старшие нули
	filtered := []byte{}
	for i, v := range values {
		if v != 0 {
			filtered = values[i:]
			break
		}
	}

	// преобразует закодированные значения в соответствующие символы
	encoded := ""
	for _, v := range filtered {
		ch, ok := s.chars.getChar(v)

		// этот код не должен никогда выполняться
		if !ok || v > maxChar {
			panic("alphabet or bit mask is corrupted")
		}

		encoded += string(ch)
	}

	return encoded
}

func (s *SimpleShortener) Decode(code string) (uint64, error) {
	runes := []rune(code)
	if len(runes) > maxLen {
		return 0, fmt.Errorf("code can not be longer than %d chars", maxLen)
	}

	var value uint64
	codeLen := len(runes)
	for i := codeLen - 1; i >= 0; i-- {
		// индекс текущей группы бит начиная с младшей
		bindex := codeLen - 1 - i
		b, ok := s.chars.get(runes[i])
		if !ok {
			return 0, fmt.Errorf("rune %q not in alphabet", runes[i])
		}
		value = value | (uint64(b) << uint(bindex*bitsInChar))
	}
	return value, nil
}
