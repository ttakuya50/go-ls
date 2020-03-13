package main

import (
	"fmt"
	"log"
	"os"
)

// a
type a struct {
	*bool
}

func (a a) isA() *bool {
	var t int
	t = 0
	log.Println(t)

	return a.bool
}

func (a a) run(fileInfos []os.FileInfo) {
	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name())
	}
}
