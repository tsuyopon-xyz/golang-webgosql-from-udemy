package main

import "fmt"

func main() {
	// 明治的
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2     int    = 200
		s2     string = "Golang"
		t2, f2 bool   = true, false
	)
	fmt.Println(i2, s2, t2, f2)

	var i3 int
	var s3 string
	var b3 bool
	println(i3, s3, b3)

	// 暗黙的な定義
	// 変数名 := 値
	i4 := 400
	println(i4)
}
