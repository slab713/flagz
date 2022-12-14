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

// Create a slice containing a list of all files starting from the specified location (exclude directories, go into subfolders)
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

const min int = 1
const max int = 197

// Function for picking a random number between 1 and 197
func pickRandomNr() int {
	rand.Seed(time.Now().UnixNano())
	randomNr := (rand.Intn(max-min+1) + min)
	return randomNr
}

func main() {

	// Call walkDir function
	search, err := walkDir("./images")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Loop over the slice returned from the walkDir function and store indices + files in 2 new slices
	var indices []int
	var files []string
	for i, el := range search {
		indices = append(indices, i+1)
		strings.Fields(el)
		files = append(files, el)
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Create a map to match index + file
	imageMap := make(map[int]string)
	for i := 0; i < len(files); i++ {
		imageMap[indices[i]] = files[i]
	}

	// Create a cache (no expiration) and store the map in it
	c := cache.New(-1, -1)
	c.Add("mapInCache", imageMap, -1)
	fmt.Println(c.Get("mapInCache"))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Login the bot
	ReadConfig()

	// Start the bot
	Start()

	// Keeps the connection going
	<-make(chan struct{})
}
