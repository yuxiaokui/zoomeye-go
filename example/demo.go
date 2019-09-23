// ➜  example go run demo.go
// Total: 445724
// 186.4.129.43:631
// 192.82.66.60:5984
// 181.211.115.26:8010
// 112.91.59.57:82
// 5.2.144.184:888
// 85.175.7.21:12345
// 218.58.70.201:5984
// 218.5.173.5:5984
// 61.183.114.243:5984
// 178.23.73.40:9200
// 139.129.5.143:8089
// 13.208.105.156:9200

package main

import (
    "fmt"

    "github.com/yuxiaokui/zoomeye-go/zoomeye"
)

func main() {
    for i := 1; i <10; i++ {
        for _,target := range zoomeye.Search("jboss",i,"xxx@xxx.com","xxxxxxxx") {
          fmt.Println(target)
        }
    }
}
