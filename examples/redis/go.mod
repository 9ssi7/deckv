module github.com/9ssi7/deckv/examples/redis

go 1.23.0

replace (
	github.com/9ssi7/deckv => ../../
	github.com/9ssi7/deckv/deckvredis => ../../deckvredis
)

require (
	github.com/9ssi7/deckv v0.0.0-00010101000000-000000000000
	github.com/9ssi7/deckv/deckvredis v0.0.0-00010101000000-000000000000
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/redis/go-redis/v9 v9.7.0 // indirect
)
