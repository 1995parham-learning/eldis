package command

import (
	"errors"
	"strings"
)

var (
	ErrSetCommandArgs  = errors.New("set command takes exactly 2 arguments")
	ErrGetCommandArgs  = errors.New("get command takes exactly 1 arguments")
	ErrCommandNotFound = errors.New("command is not available")
	ErrKeyValueLen     = errors.New("key and value should be 8 characters")
)

type Command interface{}

const (
	Len            = 8
	GetCommandArgs = 2
	SetCommandArgs = 3
	SET            = "SET"
	GET            = "GET"
)

type Get struct {
	Key string
}

type Set struct {
	Key   string
	Value string
}

func Parse(input string) (Command, error) {
	input = strings.TrimSpace(input)

	args := strings.Fields(input)

	switch args[0] {
	case SET:
		cmd, err := NewSet(args)
		if err != nil {
			return nil, err
		}

		return cmd, nil
	case GET:
		cmd, err := NewGet(args)
		if err != nil {
			return nil, err
		}

		return cmd, nil
	}

	return nil, ErrCommandNotFound
}

func NewSet(args []string) (Set, error) {
	if len(args) != SetCommandArgs {
		return Set{}, ErrSetCommandArgs
	}

	if len(args[1]) != len(args[2]) || len(args[1]) != Len {
		return Set{}, ErrKeyValueLen
	}

	return Set{Key: args[1], Value: args[2]}, nil
}

func NewGet(args []string) (Get, error) {
	if len(args) != GetCommandArgs {
		return Get{}, ErrGetCommandArgs
	}

	if len(args[1]) != Len {
		return Get{}, ErrKeyValueLen
	}

	return Get{Key: args[1]}, nil
}
