package randomhash

import (
	"fmt"
	"testing"
)

func TestCharLength(t *testing.T) {

	charLength := charLength("abcdefgh")

	if charLength == 3 {
		t.Log("pass")
	} else {
		t.Error("failed")
	}
}

func BenchmarkGenHash(b *testing.B) {
	randomHash := New("")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randomHash.GenerateHash(10)
	}

}
func TestGenHash(t *testing.T) {

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"

	randomHash := New(charset)
	result, err := randomHash.GenerateHash(10)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("random hash %s", result)
	// t.Logf(result)
}

func TestGenHashEmoji(t *testing.T) {
	charset := "😁😎😍😇🤓🤔😴😝🤑🤒😭😈👻👽🤖💩🎅💪👈👉👆👇✌✋👌👍👎👐👂👃👣👁"

	randomHash := New(charset)
	result, _ := randomHash.GenerateHash(16)
	fmt.Printf("random hash %s", result)
}

func TestEncodingEmoji(t *testing.T) {
	charset := "😁😎😍😇🤓🤔😴😝🤑🤒😭😈👻👽🤖💩🎅💪👈👉👆👇✌✋👌👍👎👐👂👃👣👁"

	randomHash := New(charset)
	data := []byte{byte(27), byte(46)}
	result := randomHash.encoding(data)

	if result == "😇👻✋" {
		t.Log("pass")
	} else {
		t.Error("failed")
	}
}

func TestEncoding(t *testing.T) {
	randomHash := New("")
	data := []byte{byte(99), byte(84), byte(1)}
	result := randomHash.encoding(data)

	if result == "y1qb" {
		t.Log("pass")
	} else {
		t.Error("failed")
	}
}
