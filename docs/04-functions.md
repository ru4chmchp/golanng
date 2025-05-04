# Functions

Go hỗ trợ nhiều dạng hàm giúp viết code rõ ràng, linh hoạt, và mạnh mẽ. Dưới đây là tổng hợp các dạng phổ biến và cách sử dụng.

---

## ✅ 1. Hàm cơ bản (Basic Function)

```go
func add(a int, b int) int {
    return a + b
}
```

🔍 Giải thích:

+ Hàm `add` nhận hai tham số `a` và `b` kiểu `int`
+ Trả về tổng của hai số `(int)`

🎯 Khi dùng: Đây là dạng hàm cơ bản nhất, dùng để tổ chức logic rõ ràng hơn.

## ✅ 2. Hàm trả về nhiều giá trị (Multiple Return Values)

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
```

🔍 Giải thích:

+ Trả về kết quả phép chia và lỗi nếu có
+ Go không có `try-catch`, nên pattern `(value, error)` cực kỳ phổ biến

🎯 Khi dùng: Hầu như mọi hàm có khả năng lỗi đều dùng kiểu này — chuẩn Go idiom.

## ✅ 3. Hàm ẩn danh (Anonymous Function)

```go
multiply := func(a, b int) int {
    return a * b
}
```

🔍 Giải thích:

+ Hàm không tên, gán vào biến `multiply`
+ Có thể truyền đi như biến hoặc dùng tại chỗ

🎯 Khi dùng: Callback logic, closures, hoặc logic tạm thời.

## ✅ 4. Hàm như tham số (Function as Argument)

```go
func operate(a, b int, f func(int, int) int) int {
    return f(a, b)
}
```

🔍 Giải thích:

+ Hàm `operate` nhận một hàm khác làm đối số
+ Hàm `f` có chữ ký `(int, int) int`

🎯 Khi dùng: Middleware, decorator, logic linh hoạt theo hàm truyền vào.

## ✅ 5. Hàm trả về hàm (Closure / Higher-order Function)

```go
func makeMultiplier(x int) func(int) int {
    return func(y int) int {
        return x * y
    }
}
```

🔍 Giải thích:

+ Biến `x` được "ghi nhớ" trong hàm trả về
+ Dùng để tạo ra các hàm mới như `double := makeMultiplier(2)`

🎯 Khi dùng: Generator, config logic, mô hình hóa DSL đơn giản.

## ✅ 6. Hàm với defer (Cleanup)

```go
func readFile() {
    fmt.Println("Opening file")
    defer fmt.Println("Closing file")
    fmt.Println("Reading file")
}
```

🔍 Giải thích:

+ defer hoãn thực thi cho đến khi hàm kết thúc
+ Dùng để cleanup: đóng file, unlock, thu hồi tài nguyên

🎯 Khi dùng: Xử lý tệp, panic, cleanup tài nguyên.

## ✅ 7. Hàm biến số (Variadic Function)

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

🔍 Giải thích:

+ `nums ...int` cho phép truyền nhiều int vào

+ Biến `nums` là slice (`[]int`)

🎯 Khi dùng: Logging, thống kê, truyền danh sách tham số tiện lợi.