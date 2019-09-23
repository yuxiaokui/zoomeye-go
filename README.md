# zoomeye-go #

zoomeye api code by golang

## Usage ##

```go
package main

import (
	"fmt"

	"github.com/yuxiaokui/zoomeye-go/zoomeye"
)

func main() {
	for i := 1; i <10; i++ {
		for _,target := range zoomeye.Search("jboss",i,"xxx@xxx.com","xxxxxx") {
			fmt.Println(target)
		}
	}
}

```


```bash
➜  example go run demo.go
Total: 445724
186.4.129.43:631

```
