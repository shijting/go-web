package main

import (
	"fmt"
	"github.com/shijting/go-web/libs/jwt"
	"sort"
	"time"
)

func sortKeys(m map[string]interface{}) []string {
	keys := make([]string, 0)
	for key, _ := range m {
		keys = append(keys, key)
	}
	//  排好序后返回(asc)
	//sort.Sort(sort.StringSlice(keys))

	// desc
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	return keys
}

func main() {
	//m := map[string]interface{}{"id": 1, "name": "张三", "age": 19, "addr": "gx"}
	//fmt.Println(m)
	//sortedKeys := sortKeys(m)
	//for _, k := range sortedKeys {
	//	fmt.Printf("%s->%v\n", k, m[k])
	//}
	aToken, rToken, err := jwt.GenToken(int64(10))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(aToken)
	fmt.Println(rToken)
	m, err := jwt.ParseToken(aToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("frist %#v\n", m)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	// 10 后请求refresh token (已过期)
	time.Sleep(time.Second * 2)
	newToken, err := jwt.ParseRefreshToken(aToken, rToken)
	if err != nil {
		fmt.Println("refreshed", err)
		return
	}

	fmt.Printf("refreshed %#v\n", newToken)

	m, err = jwt.ParseToken(newToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("再次 %#v\n", m)
}
