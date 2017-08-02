// 脩改echo程序，使其能夠打印os.Args[0]。
package main

import "fmt"

func main()  {
    fmt.Println(strings.Join(os.Args[0:], " "))
}
