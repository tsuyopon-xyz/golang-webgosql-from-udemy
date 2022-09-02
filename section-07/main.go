package main

import "fmt"

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y

	return q, r
}

func Double(price int) (result int) {
	result = price * 2

	return
}

func main() {
	// 関数
	fmt.Println(Plus(1, 2))
	var q, r int = Div(9, 3)
	fmt.Println(q, r)
	fmt.Println(Double(100))

	// 無名関数
	f := func(x, y int) int {
		return x + y
	}

	fmt.Println(f(1, 2))

	i300 := func(x, y int) int {
		return x + y
	}(100, 200)

	fmt.Println(i300)

	// 関数を返す関数
	addF := AddToX(100)
	resultOfAddF := addF(5, 10)
	fmt.Println(resultOfAddF)

	// 関数を引数に取る関数
	CallFunc(Plus)

	// クロージャー

	// ジェネレータ（クロージャの応用でできる（Goの機能レベルではない））
}

func AddToX(x int) func(int, int) int {
	return func(y, z int) int {
		return x + y + z
	}
}

func CallFunc(cb func(x, y int) int) {
	result := cb(10, 20)
	fmt.Println("result : ", result)
}
