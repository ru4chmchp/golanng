# 🧪 Bài Tập Go - Biến (Variables)

Luyện tập từ cơ bản đến nâng cao để thành thạo khai báo và sử dụng biến trong Go.

---

## ✅ Exercise 1: Khai báo và in biến

**Yêu cầu:**  
Khai báo các biến sau và in ra màn hình:
- `name` (`string`): tên người dùng
- `age` (`int`): tuổi
- `isMember` (`bool`): người dùng có phải thành viên không

📌 *Gợi ý:* dùng `fmt.Println()` để in.

---

## ✅ Exercise 2: Tính diện tích hình chữ nhật

**Yêu cầu:**  
Khai báo 2 biến `width` và `height`, sau đó tính diện tích và in ra kết quả.

---

## ✅ Exercise 3: Biến đổi giá trị

**Yêu cầu:**  
Gán một biến `x = 5`, sau đó:
- Tăng `x` lên 10 đơn vị
- Giảm `x` đi 3 đơn vị
- In kết quả cuối cùng

---

## ✅ Exercise 4: Khai báo nhiều biến

**Yêu cầu:**  
Khai báo 3 biến `a = 1`, `b = 2`, `c = 3` trong cùng một dòng.  
Tính và in tổng của 3 biến.

---

## ✅ Exercise 5: Trả về nhiều giá trị

**Yêu cầu:**  
Viết một hàm `swap(a, b int) (int, int)` trả về hoán đổi hai số.  

```go
x, y := swap(3, 4)
// Kết quả mong đợi: x = 4, y = 3
```

## ✅ Exercise 6: Hằng số

**Yêu cầu:**
Khai báo một hằng số`AppVersion = "1.0.0"`
In ra thông điệp:

```bash
Running App version: 1.0.0
```

## ✅ Exercise 7: Biến slice động

**Yêu cầu:**
Tạo một `slice scores := []int{90, 80, 70}`

Thêm phần tử 100 vào slice
Tính và in trung bình cộng của tất cả điểm

## ✅ Exercise 8: Dùng biến với map

**Yêu cầu:**
Tạo một map kiểu `map[string]int` với các key là tên học sinh, và value là điểm số.
In ra điểm của học sinh tên `"Alice"`.

