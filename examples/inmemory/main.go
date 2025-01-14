package main

import (
	"context"
	"fmt"

	"github.com/9ssi7/deckv"
)

func main() {
	client := deckv.New(deckv.WithConfFilePath("./blocklist.conf"))
	err := client.Load(context.Background())
	if err != nil {
		panic(err)
	}
	ok, err := client.Check(context.Background(), "0-mail.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(ok)
}
