package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("$ tools.exe -d C:/test -p old -r my")

	var dir = ""
	flag.StringVar(&dir, "d", "", "directory path")

	var pattern = ""
	flag.StringVar(&pattern, "p", "", "pattern name")

	var replace = ""
	flag.StringVar(&replace, "r", "", "replace name")

	flag.Parse()
	if dir == "" || pattern == "" || replace == "" {
		flag.Usage()
		return
	}

	change(dir, pattern, replace, false)
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter y to confirm > ")
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		if strings.Contains(strings.ToLower(string(sentence)), "y") {
			change(dir, pattern, replace, true)
		} else {
			os.Exit(0)
		}
	}

}

func change(dir, pattern, replace string, confirm bool) {
	// 遍历文件夹，获取文件路径
	paths := make([]string, 0)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	// 遍历文件路径，修改文件名
	for _, path := range paths {
		var st = filepath.Base(path)
		st = strings.Replace(st, pattern, replace, 1)
		newPath := filepath.Join(filepath.Dir(path), st)
		if confirm {
			os.Rename(path, newPath)
		} else {
			fmt.Println(path + " => " + newPath)
		}
	}

	if confirm {
		fmt.Println("ok")
	}
}
