package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

/*
$ docker pull redis
$ docker run —name redis-testkey-instance -p 6379:6379 -d redis
*/

func main() {

	// var ct = context.Background()
	// _ := context.WithTimeout(ct, time.Second*5)
	var host, port string

	host = "127.0.0.1"

	port = "6379"

	server := host + ":" + port

	fmt.Println("testkey connect to redis on: " + server)

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

	err := rdb.Set("testkey", "testvalue", 0).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	val, err := rdb.Get("testkey").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Get key, testkey, value: ", val)

	val2, err := rdb.Get("noexistantkey").Result()
	if err == redis.Nil {
		fmt.Println("noexistantkey does not exist")
	} else if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("noexistantkey", val2)
	}
	// Output: key value
	// noexistantkey does not exist

	payload := map[string]interface{}{
		"name": "gena",
		"age":  "300",
	}

	slb, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = rdb.HSet("htestkey", "me", slb).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	payload2 := map[string]interface{}{
		"name": "sasha",
		"age":  "100",
	}

	slb2, err := json.Marshal(payload2)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = rdb.HSet("htestkey", "me2", slb2).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	val3, err := rdb.HGet("htestkey", "me").Result()
	if err == redis.Nil {
		fmt.Println("htestkey age does not exist")
	} else if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("htestkey: ", val3)
	}

	v, err := rdb.HGetAll("htestkey").Result()
	if err == redis.Nil {
		fmt.Println("htestkey age does not exist")
	} else if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("htestkey all", v)
	}
	err = rdb.Del("testkey", "htestkey").Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}
