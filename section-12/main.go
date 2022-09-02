package main

import "fmt"

/*
type Stringer interface {
	String() string
}
*/

type Point struct {
	A int
	B string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

func main() {
	p := &Point{A: 100, B: "ABC"}
	fmt.Println(p)
}

// type MyError struct {
// 	Message string
// 	ErrCode int
// }

// func (e *MyError) Error() string {
// 	return e.Message
// }

// func RaiseError() error {
// 	return &MyError{
// 		Message: "カスタムエラーが発生しました",
// 		ErrCode: 1234,
// 	}
// }

// func main() {
// 	err := RaiseError()
// 	fmt.Println(err.Error())

// 	e, ok := err.(*MyError)
// 	if ok {
// 		fmt.Println(e.ErrCode)
// 	}
// }

// type Stringify interface {
// 	ToString() string
// }

// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p *Person) ToString() string {
// 	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
// }

// type Car struct {
// 	Name  string
// 	Model string
// }

// func (c *Car) ToString() string {
// 	return fmt.Sprintf("Name=%v, Model=%v", c.Name, c.Model)
// }

// func main() {
// 	vs := []Stringify{
// 		&Person{Name: "Taro", Age: 21},
// 		&Car{Name: "123-456", Model: "AB-1234"},
// 	}

// 	for _, v := range vs {
// 		fmt.Println(v.ToString())
// 	}
// }
