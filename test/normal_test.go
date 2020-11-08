package test

import (
	"encoding/base64"
	"log"
	"testing"
)

func TestJWTDecode(t *testing.T) {
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDQ4MjQ5NjIsImlzcyI6ImJsb2ctc2VydmljZSIsImFwcF9rZXkiOiIzNDQ3YWRmZDc0MmNkZmI5MDQ4YTNiMjliYWYxYWU3ZCIsImFwcF9zZWNyZXQiOiJlMjA4YjhiZDYwYjM5NzViZjUzMjM0Mzc4YTNkZDc3ZSJ9.5VMqFl-1MHwajggVy4lDyfsU4FM-yMZQWGXOof1uunI
	payload := "eyJleHAiOjE2MDQ4MjQ5NjIsImlzcyI6ImJsb2ctc2VydmljZSIsImFwcF9rZXkiOiIzNDQ3YWRmZDc0MmNkZmI5MDQ4YTNiMjliYWYxYWU3ZCIsImFwcF9zZWNyZXQiOiJlMjA4YjhiZDYwYjM5NzViZjUzMjM0Mzc4YTNkZDc3ZSJ9"
	DecodeString, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		panic(err)
	}
	log.Println(string(DecodeString))
}
