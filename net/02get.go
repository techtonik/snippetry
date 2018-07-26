// transpilery:primary
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://g.cn/robots.txt")
	line, _ := bufio.NewReader(resp.Body).ReadString('\n')
	fmt.Printf("%v\n%v", resp.StatusCode, line)
}

// stdout:
// 200\n
// User-agent: *\n
