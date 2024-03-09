package main

import (
	"fmt"
	"os"

	"sigs.k8s.io/yaml"
)

func main() {
	path := os.Args[1]

	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var obj interface{}
	if err = yaml.Unmarshal(content, &obj); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", obj)
}
