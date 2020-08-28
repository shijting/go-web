package main

import "fmt"

func main() {
	//a := "egg"
	//v := reflect.ValueOf(&a)
	//v = v.Elem()
	//fmt.Println(v.String())
	//
	//v.SetString("abc")
	//fmt.Println(a)
	//var b string
	//fmt.Println("var a *int IsNil:", reflect.ValueOf(b).IsValid())
	a := [4]int{0, 1, 2, 3}
	b := a[1:2]
	fmt.Println(a, b)
	b[0] = 9
	fmt.Println(a, b)
	// append 后a和b的内存有什么变化？
	b = append(b, 5)
	b[0] = 10
	fmt.Println(a, b)
	b = append(b, 6)
	b[0] = 11
	fmt.Println(a, b)
	b = append(b, 6, 7)
	b[0] = 12
	fmt.Println("往b中追加数据 超出a的最后一元素时会给b重新申请一块新内容，此时 ab再无关联(b不只是a的引用)", a, b)
}
