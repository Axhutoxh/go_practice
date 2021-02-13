//GO LANGUAGE DATATYPES
//go mainly has 3 datatypes
//1.Atomic [string,int,int32,int64,uint,uint32,uint64]
//2.Unsafe [Pointer]
//3.Abstract Datatypes [ map[]<datatypes>,struct{},interface{}]

package main

import "fmt"

func main() {
	var (
		m1 = 2
		m2 = -1
	)
	fmt.Println(m1 + m2)

	var n1 int16
	var n2 int64
	fmt.Println(int64(n1) + n2)
}
