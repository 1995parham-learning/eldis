package redis

import (
	"github.com/1995parham-learning/eldis/internal/file"
)

type Redis struct {
	Threshold  int
	Memory     map[string]string
	StoredKeys []string
	File       file.File
}

func New(th int) Redis {
	return Redis{
		Threshold:  th,
		Memory:     make(map[string]string),
		StoredKeys: make([]string, th),
		File:       file.TextFile{Threshold: th},
	}
}

func (r *Redis) Set(key string, value string) {
	r.StoredKeys[len(r.Memory)] = key
	r.Memory[key] = value

	if len(r.Memory) == r.Threshold {
		r.File.Flush(r.Memory, r.StoredKeys)
		r.Memory = make(map[string]string)
		r.StoredKeys = make([]string, r.Threshold)
	}
}

func (r *Redis) Get(key string) string {
	v, ok := r.Memory[key]
	if ok {
		return v
	}

	return r.File.Search(key)
}
