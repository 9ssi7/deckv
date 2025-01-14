// Package deckv provides a simple and efficient blocklist implementation
// with support for multiple storage backends.
//
// Deckv allows you to manage and check against blocklists (e.g., email domains,
// IP addresses, usernames) using either in-memory storage or Redis as a backend.
//
// Basic usage with in-memory storage:
//
//	client := deckv.New(deckv.WithConfFilePath("./blocklist.conf"))
//	err := client.Load(context.Background())
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	blocked, err := client.Check(context.Background(), "blocked-domain.com")
//
// Using with Redis storage:
//
//	storage := deckvredis.New(context.Background(), deckvredis.Config{
//	    Host: "localhost",
//	    Port: "6379",
//	})
//	client := deckv.New(
//	    deckv.WithConfFilePath("./blocklist.conf"),
//	    deckv.WithStorage(storage),
//	)
package deckv
