package main

import (
	"bufio"
	"cafeBazarInterview/command"
	"cafeBazarInterview/redis"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	r := redis.New(3)

	reader := bufio.NewReader(os.Stdin)

	for {
		c, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		c = strings.TrimSuffix(c, "\n")

		commands := strings.Split(c, " ")

		if commands[0] == command.SET {
			if err := redis.SetValidation(commands); err != nil {
				fmt.Println(err.Error())
				continue
			}

			r.Set(commands[1], commands[2])
		} else if commands[0] == command.GET {
			if err := redis.GetValidation(commands); err != nil {
				fmt.Println(err.Error())
				continue
			}

			fmt.Println(r.Get(commands[1]))
		} else {
			fmt.Println("Just SET and GET commands are valid")
		}
	}
}
