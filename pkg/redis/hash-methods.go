package redis

import "fmt"

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
		fmt.Println("Error: couldn't set value")
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
