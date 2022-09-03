package main

import (
	"fmt"
	"os"
)

// func main() {
// 	// os.Exit(0)
// 	// fmt.Println(1)
// }

// func main() {
// 	_, err := os.Open("A.txt")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

func main() {
	fmt.Println(os.Args)
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])

	fmt.Printf("length=%d\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}
