package main

import (
	"context"
	"fmt"

	"github.com/9ssi7/deckv"
	"github.com/9ssi7/deckv/deckvredis"
)

func main() {
	storage := deckvredis.New(deckvredis.Config{
		Host:     "localhost",
		Port:     "6379",
		Password: "",
		DB:       0,
	})
	client := deckv.New(deckv.WithConfFilePath("./blocklist.conf"), deckv.WithStorage(storage))
	err := client.Load(context.Background())
	if err != nil {
		panic(err)
	}
	ok, err := client.Check(context.Background(), deckv.FromEmail("test@0-mail.com"))
	if err != nil {
		panic(err)
	}
	fmt.Println(ok)
}
