# zoomeye-go #

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
