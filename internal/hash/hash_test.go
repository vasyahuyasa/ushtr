package hash

import (
	"fmt"
	"math/rand"
	"testing"
)

func randomStr(length int) string {
	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = 'a' + byte(rand.Intn(64))
	}
	return string(buf)
}

func TestNewWithAlphabet(t *testing.T) {
	type args struct {
		alphabet string
		hashlen  int
	}
	tests := []struct {
		name     string
		hashFunc HashFunc
		wantLen  int
	}{
		{
			"one len",
			NewWithAlphabet("zxcvbnmasdfghjklqwertyuiop1234567890_", 1),
			1,
		},
		{
			"ten len",
			NewWithAlphabet("zxcvbnmasdfghjklqwertyuiop1234567890_", 10),
			10,
		},
		{
			"100 len",
			NewWithAlphabet("zxcvbnmasdfghjklqwertyuiop1234567890_", 100),
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: generate random string
			for i := 1; i < 100; i++ {
				str := []byte(randomStr(i))

				fmt.Println(tt.hashFunc(str))

				if got := len(tt.hashFunc(str)); got != tt.wantLen {
					t.Errorf("len(hashFunc(str)) = %v, want %v", got, tt.wantLen)
				}
			}

		})
	}
}
