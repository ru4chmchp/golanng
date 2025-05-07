package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Ex1
func subtract(a, b int) int {
	return a - b
}

func square(n int) int {
	return n * n
}

// Ex2
func minMax1(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

type MinMax struct {
	Min int
	Max int
}

func minMax2(a, b int) MinMax {
	if a < b {
		return MinMax{a, b}
	} else {
		return MinMax{b, a}
	}
}

func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("khong the chia 0")
	}
	return a / b, nil
}

// Ex4

func applyTwice(f func(int) int, x int) int {
	return f(f(x))
}
func evenNumber(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

func filterEven(nums []int, f func(int) bool) []int {
	var s []int
	for _, v := range nums {
		if f(v) {
			s = append(s, v)
		}
	}
	return s
}

//Ex5

func makeAdder(x int) func(int) int {
	return func(i int) int {
		return x + i
	}
}

func count() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

//Ex6

func demoDefer() {
	fmt.Println("Start")

	defer fmt.Println("Cleanup")

	fmt.Println("Running")
}

func openAndCloseFile() {
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	fmt.Println("File opened successfully")

}

//Ex7

func average(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	var sum float64
	for _, n := range nums {
		sum += n
	}
	return sum / float64(len(nums))
}

func concatStrings(strs ...string) string {
	return strings.Join(strs, "")
}

func main() {

	//Ex1
	fmt.Println(subtract(10, 2))
	fmt.Println(square(5))

	//Ex2
	min, max := minMax1(10, 50)
	fmt.Printf("So lon hon la %d, So nho hon la %d\n", max, min)

	so_a, so_b := safeDivide(90, 4)
	fmt.Println(so_a, so_b)

	//Ex3

	power := func(x, y int) int {
		return x * y
	}
	fmt.Println(power(5, 2))

	hellos := []func() string{
		func() string { return "xin chao" },
		func() string { return "hello" },
		func() string { return "bonjour" },
	}

	for _, hello := range hellos {
		fmt.Println(hello())
	}
	//Ex4

	result := applyTwice(square, 3)
	fmt.Println(result)

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(filterEven(s, evenNumber))

	//Ex5
	add := makeAdder(5)
	fmt.Println(add(100))

	count1 := count()

	fmt.Println(count1())
	fmt.Println(count1())
	fmt.Println(count1())

	count2 := count()
	fmt.Println(count2())
	fmt.Println(count2())

	//Ex6

	demoDefer()
	openAndCloseFile()

	//Ex7
	fmt.Println("Average:", average(1, 2, 3, 4, 5))
	fmt.Println(concatStrings("Hello", " ", "world", "!"))
}
