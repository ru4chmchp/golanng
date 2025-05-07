package main

import "fmt"

func isEven_isOdd(number int) string {
	if number == 0 {
		return "0"
	} else if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func countRunes(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}
	return counts
}

type Students struct {
	Name string
	Age  int
}

func CelsiusToFahrenheit(c float64) float64 {
	f := (c * 9 / 5) + 32
	return f
}

func FahrenheitToCelsius(f float64) float64 {
	c := (f - 32) / 1.8
	return c
}

const (
	ChuNhat = iota
	ThuHai
	ThuBa
	ThuTu
	ThuNam
	ThuSau
	ThuBay
)

func tenThu(thu int) string {
	switch thu {
	case ChuNhat:
		return "Chủ Nhật"
	case ThuHai:
		return "Thứ Hai"
	case ThuBa:
		return "Thứ Ba"
	case ThuTu:
		return "Thứ Tư"
	case ThuNam:
		return "Thứ Năm"
	case ThuSau:
		return "Thứ Sáu"
	case ThuBay:
		return "Thứ Bảy"
	default:
		return "Giá trị không hợp lệ (0–6)"
	}

}

func main() {
	// Ex1
	var age int
	var widthh float64
	var name string
	var isChicken bool

	fmt.Printf("age: %T\nwidth: %T\nname: %T\nisChicken: %T\n", age, widthh, name, isChicken)

	// Ex2
	var width float64 = 3.24
	var height float64 = 7.5

	fmt.Println(width * height)

	// Ex3
	var blue int = 5

	a := float64(blue)
	b := string(blue)
	fmt.Printf("a : %T\nb : %T\n", a, b)

	// Ex4
	fmt.Println(isEven_isOdd(2123123121))

	// Ex5
	scores := []float64{1, 2, 3, 4, 5}
	total := 0.0
	for _, v := range scores { //
		total += v
	}

	fmt.Println(total / float64(len(scores)))

	// Ex6
	var s string = "nguyenhuudung"
	result := countRunes(s)
	for r, c := range result {
		fmt.Printf("'%c' : %d\n", r, c)
	}

	// Ex7

	var stdu1 Students = Students{"huudung", 21}

	fmt.Println(stdu1.Name, stdu1.Age)

	// Ex8

	people := make(map[string]int)

	for {
		var name string
		var age int

		fmt.Print("Nhập tên (hoặc 'exit' để thoát): ")
		fmt.Scanln(&name)
		if name == "exit" {
			break
		}

		fmt.Print("Nhập tuổi: ")
		fmt.Scanln(&age)

		people[name] = age
		fmt.Println("Đã lưu", name, "-", age, "tuổi\n")
	}

	fmt.Println("\n Danh sách đã lưu:")
	for name, age := range people {
		fmt.Printf("- %s: %d tuổi\n", name, age)
	}

	// Ex9
	fmt.Println(CelsiusToFahrenheit(5))
	fmt.Println(FahrenheitToCelsius(43))

	// Ex10
	fmt.Println(tenThu(2))

}
