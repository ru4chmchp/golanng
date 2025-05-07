# 🧪 Bài Tập Về Kiểu Dữ Liệu Trong Go

Tổng hợp bài tập thực hành để nắm vững các kiểu dữ liệu cơ bản và nâng cao trong ngôn ngữ lập trình Go.

---

## 📘 Level 1: Cơ Bản

### 1️⃣ Khai báo và in ra các kiểu dữ liệu cơ bản
- Khai báo các biến kiểu:
  - `int`
  - `float64`
  - `string`
  - `bool`
- In ra giá trị và kiểu dữ liệu của chúng bằng `fmt.Printf`

---

### 2️⃣ Tính diện tích hình chữ nhật
- Nhập chiều dài và chiều rộng kiểu `float64`
- Tính và in ra diện tích

---

### 3️⃣ Chuyển đổi kiểu dữ liệu
- Nhập một biến `int`
- Chuyển nó thành:
  - `float64`
  - `string`
- In kết quả ra màn hình

---

### 4️⃣ Kiểm tra số chẵn / lẻ
- Nhập một số nguyên
- Kiểm tra số đó có phải là số chẵn hay không
- In ra kết quả

---

## 🧠 Level 2: Slice, Map và Struct

### 5️⃣ Quản lý danh sách điểm
- Tạo slice `[]float64` lưu điểm của 5 học sinh
- Tính điểm trung bình

---

### 6️⃣ Đếm tần suất ký tự
- Nhập một chuỗi từ người dùng
- Dùng `map[rune]int` để đếm số lần xuất hiện của mỗi ký tự

---

### 7️⃣ Tạo struct `Student`
- Struct có các trường: `Name`, `Age`, `GPA`
- Viết hàm in thông tin học sinh

---

### 8️⃣ Nhập liệu và lưu vào map
- Cho người dùng nhập tên và tuổi
- Lưu vào map `map[string]int`
- In toàn bộ nội dung map

---

## 🧨 Level 3: Nâng Cao

### 9️⃣ Chuyển đổi nhiệt độ
- Viết hai hàm:
```go
func CelsiusToFahrenheit(c float64) float64
func FahrenheitToCelsius(f float64) float64
```
### 🔟 Mô phỏng ngày trong tuần với enum

+ Dùng const iota để tạo enum từ 0–6 đại diện cho các ngày trong tuần

+ Viết hàm in tên thứ tương ứng với số được nhập vào 