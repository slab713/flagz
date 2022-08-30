package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {
	files, err := os.ReadDir("./images")

	var num int
	for i := range files {
		num += i
	}

	fmt.Println(num)

	// numc := string(num)

	imageMap := map[string]cache.Item{}

	cache.NewFrom(cache.NoExpiration, cache.NoExpiration, imageMap)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rand.Seed(time.Now().UnixNano())

	Start()

	<-make(chan struct{})
}
