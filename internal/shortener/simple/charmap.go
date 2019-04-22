package simple

import (
	"errors"
)

type charmap struct {
	next byte
	m    map[rune]byte
	rev  map[byte]rune
}

func charmapFromString(alphabet string) (*charmap, error) {
	chars := &charmap{
		next: 0,
		m:    map[rune]byte{},
		rev:  map[byte]rune{},
	}
	for _, r := range alphabet {
		if chars.has(r) {
			return nil, errors.New("char " + string(r) + " is already present in alphabet")
		}
		chars.add(r)
	}
	return chars, nil
}

func (m *charmap) has(char rune) bool {
	_, ok := m.m[char]
	return ok
}

func (m *charmap) add(char rune) {
	m.m[char] = m.next
	m.rev[m.next] = char
	m.next++
}

func (m *charmap) get(char rune) (byte, bool) {
	val, ok := m.m[char]
	return val, ok
}

func (m *charmap) getChar(value byte) (rune, bool) {
	ch, ok := m.rev[value]
	return ch, ok
}
