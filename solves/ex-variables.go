package main

import "fmt"

func swap1(a, b int) (int, int) {
	var x int = 0
	x = a
	a = b
	b = x

	return a, b
}

func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	// Ex1
	var name string = "Huu Dung"
	var age int = 21
	var isMember bool = true

	fmt.Println("Name : ", name)
	fmt.Println("Age : ", age)
	fmt.Println("Ismember : ", isMember)

	// Ex2
	var width float64 = 3.24
	var height float64 = 7.5

	fmt.Println(width * height)

	// Ex3
	var x int = 5
	x += 10
	x -= 3

	fmt.Println(x)

	// Ex4
	var a, b, c int = 1, 2, 3

	fmt.Println(a + b + c)

	// Ex5
	x, y := swap1(5, 4)

	fmt.Println(x, y)

	// Ex6
	const AppVersion = "1.0.0"

	fmt.Println("Running App version", AppVersion)

	// Ex7
	scores := []int{90, 80, 70}

	scores = append(scores, 100)

	var total int
	for i, v := range scores {
		fmt.Printf("scores[%d] = %d\n", i, v)
		total += v
	}
	fmt.Println("Total: ", total)

	// Ex8

	points := map[string]int{
		"Alice":    100,
		"Huu Dung": 10000,
	}

	fmt.Println(points["Alice"])

}
