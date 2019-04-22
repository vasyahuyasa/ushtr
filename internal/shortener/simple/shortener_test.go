package simple

import (
	"testing"
)

func TestSimpleShortener_Encode(t *testing.T) {
	s, err := NewWithAlphabet(defaultChars)

	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name string
		id   uint64
		want string
	}{
		{
			"Zero",
			0,
			"",
		},
		{
			"1",
			1,
			"a",
		},
		{
			"10",
			64,
			"aA",
		},
		{
			"max",
			18446744073709551615,
			"h__________",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.Encode(tt.id); got != tt.want {
				t.Errorf("SimpleShortener.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleShortener_Decode(t *testing.T) {
	s, err := NewWithAlphabet(defaultChars)

	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		code    string
		want    uint64
		wantErr bool
	}{
		{
			"1",
			"a",
			1,
			false,
		},
		{
			"10",
			"aA",
			64,
			false,
		},
		{
			"max",
			"h__________",
			18446744073709551615,
			false,
		},
		{
			"empty",
			"",
			0,
			false,
		},
		{
			"long",
			"aaaaaaaaaaaa",
			0,
			true,
		},
		{
			"rus",
			"я",
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Decode(tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("SimpleShortener.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SimpleShortener.Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
