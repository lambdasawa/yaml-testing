package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	path := os.Args[1]

	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var obj any
	if err = yaml.Unmarshal([]byte(content), &obj); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", obj)
}
