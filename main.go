package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/1995parham-learning/eldis/internal/command"
	"github.com/1995parham-learning/eldis/internal/redis"
)

func main() {
	r := redis.New(3)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		cmd, err := command.Parse(input)
		if err != nil {
			log.Println(err)
		}

		switch t := cmd.(type) {
		case command.Get:
			fmt.Println(r.Get(t.Key))
		case command.Set:
			r.Set(t.Key, t.Value)
		}
	}
}
