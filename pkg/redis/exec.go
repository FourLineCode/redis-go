package redis

import "fmt"

func (r *Redis) exec(commands []string) {
	action := commands[0]

	if action == "get" {
		if len(commands) != 2 {
			usageError("GET", "GET {key}")
			return
		}

		r.get(commands[1])
	} else if action == "set" {
		if len(commands) != 3 {
			usageError("SET", "SET {key} {value}")
			return
		}

		r.set(commands[1], commands[2])
	} else if action == "del" {
		if len(commands) != 2 {
			usageError("DEL", "DEL {key}")
			return
		}

		r.del(commands[1])
	} else if action == "hget" {
		if len(commands) != 3 {
			usageError("HGET", "HGET {key} {field}")
			return
		}

		r.hget(commands[1], commands[2])
	} else if action == "hset" {
		if len(commands) != 4 {
			usageError("HSET", "HSET {key} {field} {value}")
			return
		}

		r.hset(commands[1], commands[2], commands[3])
	} else if action == "hgetall" {
		if len(commands) != 2 {
			usageError("HGETALL", "HGETALL {key}")
			return
		}

		r.hgetall(commands[1])
	} else if action == "hdel" {
		if len(commands) != 2 {
			usageError("HDEL", "HDEL {key}")
			return
		}

		r.hdel(commands[1])
	}
}

func usageError(command, usage string) {
	fmt.Printf("Error: Invalid usage for command '%s'\n", command)
	fmt.Printf("Usage: %s\n", usage)
}

func (r *Redis) get(key string) {
	val, err := r.store.Get(key)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println(val)
}

func (r *Redis) set(key, value string) {
	ok := r.store.Set(key, value)
	if !ok {
		fmt.Println("Error setting value")
		return
	}
}

func (r *Redis) del(key string) {
	if err := r.store.Delete(key); err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Printf("Deleted key '%s'\n", key)
}

func (r *Redis) hget(key, field string) {
	val, err := r.store.HGet(key, field)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println(val)
}

func (r *Redis) hgetall(key string) {
	val, err := r.store.HGetAll(key)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Printf("[ hash: %s ]\n", key)
	for field, element := range val {
		fmt.Printf("%s: %s\n", field, element)
	}
}

func (r *Redis) hset(key, field, value string) {
	ok := r.store.HSet(key, field, value)
	if !ok {
		fmt.Println("Error setting value")
		return
	}
}

func (r *Redis) hdel(key string) {
	if err := r.store.HDelete(key); err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Printf("Deleted key '%s'\n", key)
}
