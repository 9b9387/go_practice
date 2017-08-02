// 脩改echo程序，使其打印value和index，每個value和index顯示一行。
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}
}
