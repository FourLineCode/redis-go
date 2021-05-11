package redis

import "fmt"

type commandUsage struct {
	command string
	usage   string
	length  int
}

var usage map[string]commandUsage = map[string]commandUsage{
	"get": {
		command: "GET",
		usage:   "GET {key}",
		length:  2,
	},
	"set": {
		command: "SET",
		usage:   "SET {key} {value}",
		length:  3,
	},
	"del": {
		command: "DEL",
		usage:   "DEL {key}",
		length:  2,
	},
	"hget": {
		command: "HGET",
		usage:   "HGET {key} {field}",
		length:  3,
	},
	"hset": {
		command: "HSET",
		usage:   "HSET {key} {field} {value}",
		length:  4,
	},
	"hgetall": {
		command: "HGETALL",
		usage:   "HGETALL {key}",
		length:  2,
	},
	"hdel": {
		command: "HDEL",
		usage:   "HDEL {key}",
		length:  2,
	},
	"lpush": {
		command: "LPUSH",
		usage:   "LPUSH {key} {value}",
		length:  3,
	},
	"rpush": {
		command: "RPUSH",
		usage:   "RPUSH {key} {value}",
		length:  3,
	},
	"lpop": {
		command: "LPOP",
		usage:   "LPOP {key}",
		length:  2,
	},
	"rpop": {
		command: "RPOP",
		usage:   "RPOP {key}",
		length:  2,
	},
	"lindex": {
		command: "LINDEX",
		usage:   "LINDEX {key} {index}",
		length:  3,
	},
	"lrange": {
		command: "LRANGE",
		usage:   "LRANGE {key} {start} {end}",
		length:  4,
	},
	"ldel": {
		command: "LDEL",
		usage:   "LDEL {key}",
		length:  2,
	},
}

func checkUsage(commandType string, commands []string) bool {
	if len(commands) != usage[commandType].length {
		fmt.Printf("Error: Invalid usage for command '%s'\n", usage[commandType].command)
		fmt.Printf("Usage: %s\n", usage[commandType].usage)
		return false
	}

	return true
}

func getAvailableCommands() []string {
	keys := make([]string, 0, len(usage))

	for k := range usage {
		keys = append(keys, k)
	}

	return keys
}
