package redis

import (
	"fmt"
	"strconv"
)

func (r *Redis) lpush(key, value string) {
	ok := r.store.LPush(key, value)

	if !ok {
		fmt.Println("Error: couldn't pushing value")
		return
	}
}

func (r *Redis) rpush(key, value string) {
	ok := r.store.RPush(key, value)

	if !ok {
		fmt.Println("Error: couldn't pushing value")
		return
	}
}

func (r *Redis) lpop(key string) {
	val, err := r.store.LPop(key)

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println(val)
}

func (r *Redis) rpop(key string) {
	val, err := r.store.RPop(key)

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println(val)
}

func (r *Redis) lindex(key, index string) {
	in, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println("Error: Provide number as index")
		return
	}

	val, err := r.store.LIndex(key, in)

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println(val)
}

func (r *Redis) lrange(key, start, end string) {
	st, err := strconv.Atoi(start)
	if err != nil {
		fmt.Println("Error: Provide number as start")
		return
	}

	en, err := strconv.Atoi(end)
	if err != nil {
		fmt.Println("Error: Provide number as end")
		return
	}

	val, err := r.store.LRange(key, st, en+1)

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	for in, el := range val {
		fmt.Printf("%d. %s\n", in+1, el)
	}
}

func (r *Redis) ldel(key string) {
	if err := r.store.LDelete(key); err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Printf("Deleted key '%s'\n", key)
}
