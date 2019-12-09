package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// t
type t struct {
	*bool
}

// isT
func (t t) isT() *bool {
	return t.bool
}

// run
func (t t) run(fileInfos []os.FileInfo) {
	sort.Slice(fileInfos, func(i, j int) bool { return fileInfos[i].ModTime().After(fileInfos[j].ModTime()) })

	for _, fileInfo := range fileInfos {
		// 隠しファイルは非表示
		if strings.HasPrefix(fileInfo.Name(), ".") {
			continue
		}

		fmt.Println(fileInfo.Name())
	}
}
