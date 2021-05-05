package redis

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/FourLineCode/redis-go/pkg/store"
)

type Redis struct {
	store *store.Store
}

var validCommands = []string{
	"get",
	"set",
	"del",
	"hget",
	"hset",
	"hgetall",
	"hdel",
}

func NewClient() *Redis {
	return &Redis{
		store: store.New(),
	}
}

func (r *Redis) Run() {

	for {
		commands := getCommands()

		if commands[0] == "" {
			continue
		}

		if commands[0] == "exit" {
			fmt.Println("Redis client exited successfully!")
			return
		}

		valid := false
		for _, val := range validCommands {
			if val == commands[0] {
				r.exec(commands)
				valid = true
				break
			}
		}

		if !valid {
			fmt.Println("Invalid command!")
		}

	}
}

func getCommands() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("#redis> ")

	res, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Stdin: ", err)
	}

	cmd := strings.TrimSpace(strings.Replace(res, "\r\n", "", -1))

	actions := strings.Split(strings.ToLower(cmd), " ")

	return actions
}
