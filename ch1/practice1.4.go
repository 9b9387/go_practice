//練習 1.4： 脩改dup2，使其可以打印重複的行分别出現在哪些文件。
package main

import (
	"bufio"
	"fmt"
	"os"
)

type lineinfo struct {
	count     int
	filenames map[string]bool
}

func main() {
	counts := make(map[string]lineinfo)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}
	for line, info := range counts {
		if info.count > 1 {
			// This actually shows the map instead of listing its keys (laziness).
			fmt.Printf("%d\t%s\t%v\n", info.count, line, info.filenames)
		}
	}
}

// counts is a map, which is passed by reference - so modifications within
// countLines will be visible outside.
func countLines(f *os.File, filename string, counts map[string]lineinfo) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		info := counts[input.Text()]
		// The zero value for maps is 'nil'. So this has to be explicitly created
		// when encountered for the first time.
		if info.filenames == nil {
			info.filenames = make(map[string]bool)
		}
		info.count++
		info.filenames[filename] = true
		counts[input.Text()] = info
	}
}
