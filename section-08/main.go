package main

import (
	"fmt"
	"time"
)

func UseDefer() int {
	defer fmt.Println("END")
	defer fmt.Println("START")

	return 12345
}

func sub(id int) {
	for {
		fmt.Println("Sub loop", id)
		time.Sleep((100 * time.Millisecond))
	}
}

func init() {
	fmt.Println("Init1")
}

func init() {
	fmt.Println("Init2")
}

func main() {
	fmt.Println("Main")
	// go sub(1)
	// go sub(2)

	// for {
	// 	fmt.Println("Main loop")
	// 	time.Sleep((200 * time.Millisecond))
	// }

	// i := UseDefer()
	// fmt.Println(i, "@@@@@")

	// arr := []string{"A", "B", "C"}

	// for i, v := range arr {
	// 	fmt.Println(i, v)
	// }

	// sl := []string{"Python", "PHP", "Ruby"}
	// fmt.Println(sl)
	// sl = append(sl, "JavaScript")
	// fmt.Println(sl)
	// for i, v := range sl {
	// 	fmt.Println(i, v)
	// }

	// sl = append(sl, arr...)
	// for i, v := range sl {
	// 	fmt.Println(i, v)
	// }

	// m := map[string]int{"apple": 100, "banana": 200}
	// for key, value := range m {
	// 	fmt.Println(key, value)
	// }

	// s := "hello"
	// i := 123

	// si := fmt.Sprintf("%s %d", s, i)
	// fmt.Println(si)

	// var x interface{} = []uint{1, 2, 3}
	// var x interface{} = 1
	// i, isInt := x.(int)
	// fmt.Println(i, isInt)

	// f, isFloat64 := x.(float64)
	// fmt.Println("@@@", f, isFloat64)
	// switch v := x.(type) {
	// case int:
	// 	fmt.Println(v, "x is Int")
	// case string:
	// 	fmt.Println(v, "x is String")
	// default:
	// 	fmt.Printf("%T @@@ %s\n", x, v)
	// }
}
