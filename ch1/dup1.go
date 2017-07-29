package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
shihongyang: ch1 $ echo 'hello' > /tmp/lines
shihongyang: ch1 $ echo 'hello' >> /tmp/lines
shihongyang: ch1 $ cat /tmp/lines | go run dup1.go
*/
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
		fmt.Println(input.Text())
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
