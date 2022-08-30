package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

func walkDir(root string) ([]string, error) {
	var contents []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			contents = append(contents, path)
		}
		return nil
	})
	return contents, err
}

func main() {

	search, err := walkDir("./images")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var indices []int
	var files []string
	for i, el := range search {
		indices = append(indices, i+1)
		strings.Fields(el)
		files = append(files, el)
	}

	imageMap := make(map[int]string)

	for i := 0; i < len(files); i++ {
		imageMap[indices[i]] = files[i]
	}

	c := cache.New(-1, -1)
	c.Add("mapInCache", imageMap, -1)
	fmt.Println(c.Get("mapInCache"))

	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 197
	randomNr := (rand.Intn(max-min+1) + min)
	fmt.Println(randomNr)

	ReadConfig()

	Start()

	<-make(chan struct{})
}
