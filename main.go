package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	opt := newOpt()
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

	switch {
	case *opt.a.isA(): // 全て表示
		opt.a.run(fileInfos)
	case *opt.t.isT(): // 更新時間順にソート
		opt.t.run(fileInfos)
	default:
		// オプションなし
		for _, fileInfo := range fileInfos {
			// 隠しファイルは非表示
			if strings.HasPrefix(fileInfo.Name(), ".") {
				continue
			}

			fmt.Println(fileInfo.Name())
		}
	}
}
