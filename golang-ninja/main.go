package main

import (
	"fmt"
	"golang-ninja/basic/cache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	deletedUserId := cache.Get("userId") // Переименовано userId

	fmt.Println(deletedUserId)
}
