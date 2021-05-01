package redis

import "fmt"

func (r *Redis) exec(commands []string) {
	action := commands[0]

	if action == "get" {
		if len(commands) != 2 {
			fmt.Println("Error: Invalid usage for command 'GET'")
			fmt.Println("Usage: GET {key}")
			return
		}

		val, err := r.store.Get(commands[1])
		if err != nil {
			fmt.Println("Error:", err.Error())
			return
		}

		fmt.Println(val)
	} else if action == "set" {
		if len(commands) != 3 {
			fmt.Println("Error: Invalid usage for command 'SET'")
			fmt.Println("Usage: SET {key} {value}")
			return
		}

		ok := r.store.Set(commands[1], commands[2])
		if !ok {
			fmt.Println("Error setting value")
			return
		}

	} else if action == "del" {
		if len(commands) != 2 {
			fmt.Println("Error: Invalid usage for command 'DEL'")
			fmt.Println("Usage: DEL {key}")
			return
		}

		if err := r.store.Delete(commands[1]); err != nil {
			fmt.Println("Error:", err.Error())
			return
		}

		fmt.Printf("Deleted key '%s'\n", commands[1])
	}
}
