package models

import (
	"sync"
)

const shardCount = 32

type Store struct {
	shards [shardCount]*shard
}

type shard struct {
	sync.RWMutex
	data map[string]string
}

func NewStore() *Store {
	s := &Store{}
	for i := 0; i < shardCount; i++ {
		s.shards[i] = &shard{data: make(map[string]string)}
	}
	return s
}

func (s *Store) getShard(key string) *shard {
	return s.shards[uint(fnv32(key)%shardCount)]
}

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

func (s *Store) Set(key, value string) {
	shard := s.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	shard.data[key] = value
}

func (s *Store) Get(key string) (string, bool) {
	shard := s.getShard(key)
	shard.RLock()
	defer shard.RUnlock()
	val, ok := shard.data[key]
	return val, ok
}
