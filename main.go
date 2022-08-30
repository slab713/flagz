package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
	fmt.Println(imageMap)

	cache.New(cache.NoExpiration, cache.NoExpiration)

	ReadConfig()

	Start()

	<-make(chan struct{})
}
