// Package deckv provides blocklist functionality with multiple storage backend support.
package deckv

import (
	"bufio"
	"context"
	"os"
	"strings"
)

// Deckv represents the main blocklist client.
type Deckv struct {
	cfg *Config
}

// New creates a new Deckv instance with the provided options.
func New(opts ...Option) *Deckv {
	cfg := &Config{}
	for _, opt := range opts {
		opt(cfg)
	}
	withDefaults(cfg)
	return &Deckv{cfg: cfg}
}

// Load reads the blocklist from the configured file and stores it in the configured storage.
// If no configuration file path or storage is set, this operation is a no-op.
func (d *Deckv) Load(ctx context.Context) error {
	if d.cfg.ConfFilePath == "" {
		return nil
	}
	if d.cfg.Storage == nil {
		return nil
	}
	list := map[string]uint8{}
	f, err := os.Open(d.cfg.ConfFilePath)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		list[line] = 1
	}
	return d.cfg.Storage.Save(ctx, list)
}

// Check verifies if the provided key exists in the blocklist.
// Returns true if the key is blocked, false otherwise.
func (d *Deckv) Check(ctx context.Context, key string) (bool, error) {
	if d.cfg.Storage == nil {
		return false, nil
	}
	return d.cfg.Storage.Check(ctx, key)
}
