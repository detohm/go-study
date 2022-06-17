package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	// store basic value
	err = client.Set("name", "john", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	// get the value from key string
	val, err := client.Get("name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

	// store composite value
	jsonValue, err := json.Marshal(Person{"Devid P.", 12})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("user111", jsonValue, 0).Err()
	if err != nil {
		fmt.Println("store err:", err)
	}

	// get the composite value
	comVal, err := client.Get("user111").Bytes()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", comVal)

	var p Person
	err = json.Unmarshal(comVal, &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", p)
}
