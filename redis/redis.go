package redis

import (
	"cafeBazarInterview/file"
	"errors"
)

type Redis struct {
	Threshold int
	Memory    map[string]string
	File      file.File
}

func New(th int) Redis {
	return Redis{
		Threshold: th,
		Memory:    make(map[string]string),
		File:      file.TextFile{},
	}
}

func (r *Redis) Set(key string, value string) {
	r.Memory[key] = value

	if len(r.Memory) == r.Threshold {
		r.File.Flush(r.Memory)
		r.Memory = make(map[string]string)
	}
}

func (r *Redis) Get(key string) string {
	v, ok := r.Memory[key]
	if ok {
		return v
	}

	return r.File.Search(key)
}

func SetValidation(commands []string) error {
	if len(commands) != 3 {
		return errors.New("SET command takes 2 arguments")
	}

	if !(len(commands[1]) == len(commands[2]) && len(commands[1]) == 8) {
		return errors.New("key and value should be 8 characters")
	}

	return nil
}

func GetValidation(commands []string) error {
	if len(commands) != 2 {
		return errors.New("SET command takes 1 argument")
	}

	return nil
}
