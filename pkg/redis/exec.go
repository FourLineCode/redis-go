package redis

func (r *Redis) exec(cmd []string) {
	action := cmd[0]

	ok := checkUsage(action, cmd)
	if !ok {
		return
	}

	switch action {
	case "get":
		r.get(cmd[1])
	case "set":
		r.set(cmd[1], cmd[2])
	case "del":
		r.del(cmd[1])
	case "hget":
		r.hget(cmd[1], cmd[2])
	case "hset":
		r.hset(cmd[1], cmd[2], cmd[3])
	case "hgetall":
		r.hgetall(cmd[1])
	case "hdel":
		r.hdel(cmd[1])
	case "lpush":
		r.lpush(cmd[1], cmd[2])
	case "rpush":
		r.rpush(cmd[1], cmd[2])
	case "lpop":
		r.lpop(cmd[1])
	case "rpop":
		r.rpop(cmd[1])
	case "lindex":
		r.lindex(cmd[1], cmd[2])
	case "lrange":
		r.lrange(cmd[1], cmd[2], cmd[3])
	case "ldel":
		r.ldel(cmd[1])
	default:
		return
	}

}
