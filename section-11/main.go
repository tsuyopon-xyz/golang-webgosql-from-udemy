package main

import "fmt"

type T struct {
	// User User
	User
}

type User struct {
	Name string
	Age  uint
	// X, Y int
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u *User) SetName(name string) {
	u.Name = name
}

func NewUser(name string, age uint) *User {
	return &User{Name: name, Age: age}
}

func NewUser2(u User) *User {
	return &User{Name: u.Name, Age: u.Age}
}

type Users []*User

/*
type Users struct {
	Users []*Users
}
*/

type MyInt int

func (mi MyInt) Print() {
	fmt.Println(mi, "@@@@@@@@@")
}

func main() {
	var mi MyInt
	fmt.Printf("%T\n", mi)

	mi = 100

	mi.Print()
}

// func main() {
// 	m := map[int]*User{
// 		1: {Name: "user1", Age: 10},
// 		2: {Name: "user2", Age: 20},
// 		3: {Name: "user3", Age: 30},
// 	}
// 	fmt.Println(m)

// 	for k, v := range m {
// 		v.SetName(fmt.Sprintf("%s%d", "Updated Name", k))

// 		fmt.Println("k : ", k)
// 		fmt.Println("v : ", v)
// 	}
// }

// func main() {
// 	users := Users{}
// 	for i := 0; i <= 3; i++ {
// 		name := fmt.Sprintf("user%d", i)
// 		// user := NewUser(name, uint(10*i))
// 		user := NewUser2(User{Name: name, Age: uint(10 * 1)})
// 		fmt.Println(user)
// 		users = append(users, user)
// 	}

// 	fmt.Println(users)
// 	for _, u := range users {
// 		fmt.Println(*u)
// 	}
// }

// func main() {
// 	t := T{User{Name: "Taro", Age: 12}}
// 	fmt.Println(t.User)
// 	fmt.Println(t.User.Name)
// 	t.User.SetName("Hanako")
// 	t.User.SayName()
// 	t.SayName()
// 	// t.SetName("Hanako")
// 	// t.SayName()
// }

// func main() {
// 	var user1 User
// 	user1.Name = "Taro"
// 	user1.Age = 11

// 	fmt.Println(user1)

// 	user2 := User{
// 		Name: "Hanako",
// 		Age:  22,
// 	}
// 	fmt.Println(user2)

// 	user3 := new(User)
// 	fmt.Println(user3)

// 	user4 := &User{}
// 	fmt.Println(user4)
// }
