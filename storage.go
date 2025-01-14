package deckv

import "context"

// DataStorage defines the interface for blocklist storage backends.
type DataStorage interface {
	// Save stores the provided blocklist data in the storage backend.
	Save(ctx context.Context, data map[string]uint8) error
	// Check verifies if a key exists in the storage backend.
	Check(ctx context.Context, key string) (bool, error)
}

// InMemStorage implements DataStorage interface using in-memory storage.
type InMemStorage struct {
	data map[string]uint8
}

// Save stores the blocklist data in memory.
func (s *InMemStorage) Save(ctx context.Context, data map[string]uint8) error {
	s.data = data
	return nil
}

// Check verifies if a key exists in the in-memory storage.
func (s *InMemStorage) Check(ctx context.Context, key string) (bool, error) {
	return s.data[key] == 1, nil
}
