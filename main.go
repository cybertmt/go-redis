package main

import (
	"fmt"

	"encoding/json"
	c "goredis/constants"

	"github.com/go-redis/redis"
)

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     c.RedisHost + c.RedisPort,
		Password: "",
		DB:       0,
	})

	json, err := json.Marshal(Author{Name: "Elliot", Age: 35})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("id1234", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.Get("id1234").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}
