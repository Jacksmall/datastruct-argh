package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("欢迎切片！")

	slice := make([]int, 3)
	slice[0] = 1
	slice[1] = 3

	fmt.Println(len(slice), cap(slice))

	slice = append(slice, 1, 3)

	fmt.Println(len(slice), cap(slice))

	// for _, v := range slice {
	// 	fmt.Printf("%d ", v)
	// }
	// fmt.Printf("\n")
	// // 切片的截取 - 如果修改copySlice 会改变原来slice的值，这是因为截取操作是将新的切片Data通过指针指向数组
	// copySlice := slice[0:2]
	// copySlice[0] = 5

	// for _, v := range slice {
	// 	fmt.Printf("%d ", v)
	// }
	// fmt.Printf("\n")
	vmap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(vmap["two"]) // 2
	delete(vmap, "two")
	fmt.Println(vmap["two"]) // 0

	var r interface{}
	fmt.Println(reflect.TypeOf(r))
}

// interface 并不是动态校验的
type ReadCloser interface {
	Read(b []byte) (n int, err error)
	Close()
}

// r 可以是任意类型的值 只要实现Read 和 Close 方法的
func ReadAndWrite(r ReadCloser, buf []byte) (n int, err error) {
	for len(buf) > 0 && err != nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	r.Close()
	return
}
