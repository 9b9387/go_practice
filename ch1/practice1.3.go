// 上手實踐前面提到的strings.Join和直接Println，併觀察輸出結果的區别。
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
}
