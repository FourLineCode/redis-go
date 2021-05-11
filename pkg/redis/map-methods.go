package redis

import "fmt"

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
		fmt.Println("Error: couldn't set value")
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
