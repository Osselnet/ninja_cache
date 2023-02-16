package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	var err error

	cache := NewCache()

	cache.Set("userId", 42, time.Second * 5)
	userId, _ := cache.Get("userId")	

	fmt.Println(userId)

	time.Sleep(time.Second * 6)

	userId, err = cache.Get("userId")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userId)
}
