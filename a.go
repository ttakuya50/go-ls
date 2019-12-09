package main

import (
	"fmt"
	"os"
)

// a
type a struct {
	*bool
}

// isA
func (a a) isA() *bool {
	return a.bool
}

// run
func (a a) run(fileInfos []os.FileInfo) {
	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name())
	}
}
