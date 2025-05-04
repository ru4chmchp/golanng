# Data types

## ✅ 1. Basic Data types

| **Kiểu**                                 | **Mô tả**                                             | **Ví dụ**                        |
|------------------------------------------|------------------------------------------------------|----------------------------------|
| `bool`                                   | Giá trị logic                                         | `true`, `false`                 |
| `string`                                 | Chuỗi ký tự                                           | `"hello"`                       |
| `int`, `int8`, `int16`, `int32`, `int64` | Số nguyên có dấu                                      | `var a int = 10`                |
| `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr` | Số nguyên không dấu                        | `var b uint8 = 255`             |
| `byte`                                   | Alias của `uint8`, thường dùng cho chuỗi/byte slice  | `[]byte("abc")`                |
| `rune`                                   | Alias của `int32`, đại diện ký tự Unicode            | `'あ'`, `'A'`                   |
| `float32`, `float64`                     | Số thực dấu phẩy động                                 | `3.14`                          |
| `complex64`, `complex128`                | Số phức                                                | `complex(1, 2)`

## ✅ 2. Kiểu dữ liệu cấu trúc (Composite Types)

| **Kiểu**   | **Mô tả**                         | **Ví dụ**                                                  |
|------------|-----------------------------------|-------------------------------------------------------------|
| `array`    | Mảng có độ dài cố định            | `[3]int{1, 2, 3}`                                           |
| `slice`    | Mảng động (linh hoạt độ dài)      | `[]int{1, 2, 3}`                                            |
| `map`      | Bảng ánh xạ `key -> value`        | `map[string]int{"a": 1}`                                   |
| `struct`   | Nhóm nhiều trường (fields)        | `type Person struct { Name string; Age int }`              |

## ✅ 3. Kiểu tham chiếu (Reference Types)

+ `pointer`: trỏ tới địa chỉ bộ nhớ (`var p *int`)

+ `channel`: giao tiếp giữa các goroutine (`chan int`)

+ `interface`: định nghĩa hành vi trừu tượng (`interface{}` hoặc `interface` cụ thể)

+ `function`: kiểu hàm (`func(int) string`)

## ✅ 4. Kiểu đặc biệt

| **Kiểu**       | **Mô tả**                                                                 |
|----------------|---------------------------------------------------------------------------|
| `nil`          | Giá trị rỗng cho các reference types như `pointer`, `slice`, `map`, `func`, `interface`, `channel` |
| `error`        | Kiểu tiêu chuẩn để biểu diễn lỗi. Được tạo bằng `errors.New("message")` hoặc `fmt.Errorf(...)` |
| `interface{}`  | Kiểu trống – đại diện cho mọi kiểu dữ liệu. Dùng để viết hàm tổng quát (generic-like) |

---

Trong thực tế phát triển ứng dụng với Go, một số kiểu dữ liệu được sử dụng thường xuyên hơn nhiều so với phần còn lại. Dưới đây là tổng hợp các nhóm kiểu dữ liệu phổ biến nhất và tình huống sử dụng của chúng.

---

## ✅ 1. Kiểu cơ bản — siêu phổ biến

| **Kiểu**  | **Tình huống sử dụng**                                    |
|-----------|------------------------------------------------------------|
| `string`  | Hầu như ở mọi nơi: log, input/output, xử lý text          |
| `int`, `int64` | Dùng cho ID, đếm vòng lặp, xử lý số nguyên                 |
| `bool`    | Điều kiện, kiểm tra logic                                 |
| `float64` | Xử lý số thực: tính toán, thống kê                        |
| `error`   | Cực kỳ quan trọng trong xử lý lỗi (chuẩn Go)             |

---

## ✅ 2. Kiểu cấu trúc — cốt lõi cho modeling

| **Kiểu**             | **Tình huống sử dụng**                                       |
|----------------------|--------------------------------------------------------------|
| `struct`             | Mô hình hóa dữ liệu (user, post, product, ...)               |
| `[]T` (slice)        | Lưu danh sách (users, records, ...)                          |
| `map[string]interface{}` | Xử lý JSON động hoặc dữ liệu không có schema cứng                 |
| `map[string]string` / `map[string]int` | Lưu cấu hình, ánh xạ lookup nhanh                        |

---

## ✅ 3. Kiểu tham chiếu — dùng trong concurrency và interface

| **Kiểu**         | **Tình huống sử dụng**                                            |
|------------------|-------------------------------------------------------------------|
| `interface{}` / `any` | Tổng quát hóa dữ liệu, đặc biệt khi parse JSON hoặc dùng `context` |
| `chan`           | Truyền dữ liệu giữa các goroutine (concurrent patterns)          |
| `func`           | Hàm dưới dạng biến – dùng trong callback, middleware,...         |

---

## 📌 Thống kê từ thực tế (trích từ các repo Go phổ biến trên GitHub)

| **Kiểu**                    | **Tần suất sử dụng**                                      |
|-----------------------------|-----------------------------------------------------------|
| `string`, `int`, `bool`     | ✔️ Gần như 100% dự án                                     |
| `struct`, `[]T`, `map[...]` | ✔️ Rất phổ biến trong model, xử lý dữ liệu                |
| `error`                     | ✔️ Có mặt ở hầu hết các hàm như một phần chuẩn lỗi         |
| `interface{}`               | 🤔 Dùng vừa phải – cần cẩn thận khi dùng vì mất kiểm tra kiểu |
| `chan`, `func`              | 💡 Dùng nhiều trong app lớn: microservice, worker,...     |

---

> 📘 Bạn nên nắm vững nhóm **kiểu cơ bản** và **kiểu cấu trúc**, vì chúng chiếm tới ~90% code thực tế bạn sẽ viết với Go.

---