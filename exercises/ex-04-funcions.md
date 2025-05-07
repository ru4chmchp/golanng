# 🧠 Bài Tập Thực Hành Với Các Dạng Hàm Trong Golang

---

## ✅ 1. Hàm Cơ Bản (Basic Function)

### 📝 Bài tập:
- Viết hàm `subtract(a, b int) int` trả về hiệu của 2 số
- Viết hàm `square(n int) int` trả về bình phương

---

## ✅ 2. Hàm Trả Về Nhiều Giá Trị

### 📝 Bài tập:
- Viết hàm `minMax(a, b int) (int, int)` trả về giá trị nhỏ và lớn hơn
- Viết hàm `safeDivide(a, b float64) (float64, error)` xử lý chia 0 như ví dụ đã nêu

---

## ✅ 3. Hàm Ẩn Danh (Anonymous Function)

### 📝 Bài tập:
- Viết một hàm ẩn danh gán vào biến `power` để tính lũy thừa `x^y` (x, y là int)
- Viết một slice các hàm ẩn danh trả về các chuỗi `"Xin chào"`, `"Hello"`, `"Bonjour"`

---

## ✅ 4. Hàm Làm Đối Số (Function as Argument)

### 📝 Bài tập:
- Viết hàm `applyTwice(f func(int) int, x int) int` → áp dụng f hai lần lên x
- Viết hàm `filterEven(nums []int, f func(int) bool) []int` → lọc các số chẵn dùng hàm truyền vào

---

## ✅ 5. Hàm Trả Về Hàm (Closure / Higher-order Function)

### 📝 Bài tập:
- Viết hàm `makeAdder(x int) func(int) int` → trả về hàm cộng thêm `x`
- Viết hàm `counter() func() int` → mỗi lần gọi tăng giá trị trả về 1 đơn vị

---

## ✅ 6. Hàm Với defer (Cleanup)

### 📝 Bài tập:
- Viết hàm `demoDefer()` in các bước:
  - "Start"
  - defer `"Cleanup"`
  - "Running"
- Viết hàm mở và đóng file, sử dụng `defer` để đảm bảo đóng file dù có lỗi

---

## ✅ 7. Hàm Biến Số (Variadic Function)

### 📝 Bài tập:
- Viết hàm `average(nums ...float64) float64` tính trung bình
- Viết hàm `concatStrings(strs ...string) string` nối các chuỗi lại thành một chuỗi duy nhất

---

📌 **Gợi ý nâng cao**: Hãy thử viết unit test đơn giản cho một vài hàm bạn vừa hoàn thành.

