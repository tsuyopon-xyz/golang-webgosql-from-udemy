package main

import (
	"fmt"
)

func Sum(numbers ...int) int {
	fmt.Println(numbers)
	fmt.Printf("%T\n", numbers)

	sum := 0
	for _, v := range numbers {
		sum += v
	}

	return sum
}

func main() {
	// result1 := Sum(1, 2, 3, 4)
	// fmt.Println("result1 : ", result1)

	// sl := []int{1, 2, 3}
	// result2 := Sum(sl...)
	// fmt.Println("result2 : ", result2)

	// sl := []int{1, 2, 3, 4, 5, 6}
	// sl2 := make([]int, 6)
	// copy(sl2, sl)

	// sl2[0] = 100
	// sl2 = append(sl2, 91, 92, 93, 94, 95, 96)

	// fmt.Println(sl)
	// fmt.Println(sl2)
	// fmt.Println(cap(sl2))

	// チャネル

	// ch := make(chan int, 5)

	// ch <- 1
	// ch <- 2
	// ch <- 3

	// defer func() {
	// 	if n := recover(); n != nil {
	// 		fmt.Println("error : ", n)
	// 	}
	// }()

	// for k := range ch {
	// 	fmt.Println(k)
	// }

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// go receiver("ch1", ch1)
	// go receiver("ch2", ch2)

	// i := 0
	// for i < 100 {
	// 	ch1 <- i
	// 	ch2 <- i
	// 	time.Sleep(50 * time.Millisecond)

	// 	i++
	// }

	// ch1 := make(chan int)

	// go receiver("1.goroutin", ch1)
	// go receiver("2.goroutin", ch1)
	// go receiver("3.goroutin", ch1)

	// i := 0
	// for i < 100 {
	// 	ch1 <- i
	// 	time.Sleep(50 * time.Millisecond)

	// 	i++
	// }
	// close(ch1)
	// time.Sleep(3 * time.Second)

	// ch := make(chan int, 3)
	// ch <- 1
	// ch <- 2
	// ch <- 3
	// close(ch)
	// for i := range ch {
	// 	fmt.Println(i)
	// }

	// ch1 := make(chan int, 2)
	// ch2 := make(chan string, 2)

	// ch2 <- "A"
	// ch1 <- 1

	// select {
	// case v1 := <-ch1:
	// 	fmt.Println(v1 + 1000)
	// case v2 := <-ch2:
	// 	fmt.Println(v2 + "!?")
	// default:
	// 	fmt.Println("どちらでもない")
	// }

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for {
			i := <-ch1
			ch2 <- i * 2
		}
	}()

	go func() {
		for {
			i := <-ch2
			ch3 <- i - 1
		}
	}()

	n := 0
	for {
		select {
		case ch1 <- n:
			n++
		case i := <-ch3:
			fmt.Println(i)
		}

		if n > 100 {
			break
		}
	}

}

func receiver(name string, c chan int) {
	for {
		i, ok := <-c
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + " END")
}
