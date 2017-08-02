// 练习1.8：修改fetch这个范例，如果输入的url参数没有
// http://前缀的话，为这个url加上该前缀。你可能会用
// 到strings.HasPrefix这个函数。
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

//shihongyang: ch1 $ ./fetch www.baidu.com
func main() {
	for _, url := range os.Args[1:] {

		if strings.HasPrefix(url, "http://") == false {
			url = "http://" + url
		}

		resp, err1 := http.Get(url)
		if err1 != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err1)
			os.Exit(1)
		}
		_, err2 := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err2 != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err2)
			os.Exit(1)
		}
	}
}
