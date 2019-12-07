package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	aFlg := flag.Bool("a", false, "Include directory entries whose names begin with a dot (.).")
	tFlg := flag.Bool("t", false, "Sort by time modified (most recently modified first) before sorting the operands by lexicographical order.")
	flag.Parse()

	// カレントディレクトリ取得
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// 全て表示
	if *aFlg {
		for _, fileInfo := range fileInfos {
			fmt.Println(fileInfo.Name())
		}
		return
	}

	// 更新時間順にソート
	if *tFlg {
		sort.Slice(fileInfos, func(i, j int) bool { return fileInfos[i].ModTime().After(fileInfos[j].ModTime()) })

		for _, fileInfo := range fileInfos {
			// 隠しファイルは非表示
			if strings.HasPrefix(fileInfo.Name(), ".") {
				continue
			}

			fmt.Println(fileInfo.Name())
		}
		return
	}

	// オプションなし
	for _, fileInfo := range fileInfos {
		// 隠しファイルは非表示
		if strings.HasPrefix(fileInfo.Name(), ".") {
			continue
		}

		fmt.Println(fileInfo.Name())
	}
}
