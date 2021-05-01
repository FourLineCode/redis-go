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
	s *store.Store
}

var validCommands = []string{
	"get",
	"set",
	"del",
}

func NewClient() *Redis {
	return &Redis{
		s: store.New(),
	}
}

func (r *Redis) Run() {

	for {
		commands := getInput()

		if strings.ToLower(commands[0]) == "exit" {
			fmt.Println("Redis client exited successfully!")
			return
		}

		for _, val := range validCommands {
			if val == commands[0] {
				exec(commands)
				break
			}
		}

		fmt.Println("Invalid command!")

	}
}

func getInput() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("#redis> ")

	res, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Stdin: ", err)
	}

	cmd := strings.TrimSpace(strings.Replace(res, "\r\n", "", -1))

	actions := strings.Split(cmd, " ")

	return actions
}
