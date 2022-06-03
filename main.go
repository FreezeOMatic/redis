package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

/*
$ docker pull redis
$ docker run —name redis-test-instance -p 6379:6379 -d redis
*/

func main() {

	var ct = context.Background()
	ctx, _ := context.WithTimeout(ct, time.Second*5)
	var host, port string

	host = "127.0.0.1"

	port = "6379"

	server := host + ":" + port

	fmt.Println("Test connect to redis on: " + server)

	// if wanna use cluster
	/*
	   redis.NewClusterClient(&redis.ClusterOptions{
	   Addrs: []string{"foo", "bar"},
	   })
	*/
	rdb := redis.NewClient(&redis.Options{
		Addr:     server,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "test", "testtest", 0).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	val, err := rdb.Get(ctx, "test").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist

	payload := map[string]interface{}{
		"name": "serega",
		"age":  "300",
	}

	slb, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = rdb.HSet(ctx, "htest", "me", slb).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	payload2 := map[string]interface{}{
		"name": "nata",
		"age":  "100",
	}

	slb2, err := json.Marshal(payload2)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = rdb.HSet(ctx, "htest", "me2", slb2).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	val3, err := rdb.HGet(ctx, "htest", "me2").Result()
	if err == redis.Nil {
		fmt.Println("htest age does not exist")
	} else if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("htest age", val3)
	}

	v, err := rdb.HGetAll(ctx, "htest").Result()
	if err == redis.Nil {
		fmt.Println("htest age does not exist")
	} else if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("htest all", v)
	}
}
